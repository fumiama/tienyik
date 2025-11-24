package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"gopkg.in/yaml.v3"

	base14 "github.com/fumiama/go-base16384"
	"github.com/fumiama/tienyik"
	"github.com/fumiama/tienyik/api/auth"
	"github.com/fumiama/tienyik/api/cdserv"
	"github.com/fumiama/tienyik/api/desktop"
	"github.com/fumiama/tienyik/hcli"
	"github.com/fumiama/tienyik/internal/log"
	"github.com/fumiama/tienyik/internal/textio"
)

type config struct {
	UserAccount      string `yaml:"UserAccount"`
	PasswordB14      string `yaml:"PasswordB14"`      // PasswordB14 is a little bit more secure
	CheckIntervalSec int    `yaml:"CheckIntervalSec"` // CheckIntervalSec default 60
	GetDeviceCount   int    `yaml:"GetDeviceCount"`   // GetDeviceCount default 20
}

func main() {
	c := flag.String("c", "config.yaml", "load config file")
	s := flag.String("s", "", "save config file template")
	flag.Parse()

	cfg := config{}

	if *s != "" {
		fmt.Print("UserAccount: ")
		fmt.Scanln(&cfg.UserAccount)
		pwd := ""
		fmt.Print("Password: ")
		textio.NoEchoScanln(&pwd)
		cfg.PasswordB14 = base14.EncodeString(pwd)
		cfg.CheckIntervalSec = 60
		cfg.GetDeviceCount = 20
		data, err := yaml.Marshal(&cfg)
		if err != nil {
			log.Fatalln(err)
		}
		err = os.WriteFile(*s, data, 0644)
		if err != nil {
			log.Fatalln(err)
		}
		return
	}

	f, err := os.Open(*c)
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.NewDecoder(f).Decode(&cfg)
	_ = f.Close()
	if err != nil {
		log.Fatalln(err)
	}
	if cfg.UserAccount == "" {
		log.Fatalln("user account must be set")
	}
	if cfg.PasswordB14 == "" {
		log.Fatalln("password must be set (in b14 format)")
	}
	if cfg.CheckIntervalSec <= 0 {
		cfg.CheckIntervalSec = 60
	}
	if cfg.GetDeviceCount <= 0 {
		cfg.GetDeviceCount = 20
	}

RECONN:
	cli := hcli.NewClient()
	sd, err := cdserv.GetServData()
	if err != nil {
		log.Fatalln(err)
	}
	x, err := auth.GenChallengeData(nil, cli)
	if err != nil {
		log.Fatalln(err)
	}
	sd.SetClient(cli)
	pwd := base14.DecodeString(cfg.PasswordB14)
	rsp, err := auth.Login(nil, cli, &auth.RequestLogin{
		UserAccount:    cfg.UserAccount,
		Password:       tienyik.ChallengePassword(pwd, x.ChallengeCode),
		SHA256Password: tienyik.ChallengeSHA256Password(pwd, x.ChallengeCode),
		ChallengeID:    x.ChallengeID,
		DeviceCode:     cli.Devicecode,
		DeviceName:     tienyik.DeviceNameEdge,
		DeviceType:     cli.Devicetype,
		DeviceModel:    tienyik.DeviceModelMacOS,
		AppVersion:     tienyik.AppVersion,
		SysVersion:     tienyik.DeviceModelMacOS,
		ClientVersion:  cli.Version,
	})
	if err != nil {
		log.Fatalln(err)
	}
	rsp.SetClient(cli)
	defer auth.Logout(nil, cli)

	pd, err := desktop.PageDesktop(nil, cli, &desktop.RequestPageDesktop{
		GetCnt:       20,
		DesktopTypes: []string{"1", tienyik.ArchX86},
		SortType:     desktop.DefaultRequestPageDesktopSortType,
	})
	if err != nil {
		log.Fatalln(err)
	}

	mp := make(map[string][2]string, len(pd.DesktopList)*4)
	sb := strings.Builder{}
	sb.WriteString("available desktops:")
	for _, x := range pd.DesktopList {
		if x.UseStatusText == "运行中" {
			sb.WriteString(" |●")
		} else {
			sb.WriteString(" |○")
		}
		sb.WriteString("[")
		sb.WriteString(x.ObjID)
		sb.WriteString("]")
		sb.WriteString(x.ObjName)
		sb.WriteString("(")
		sb.WriteString(x.OsName)
		sb.WriteString(")|")
		mp[x.ObjID] = [2]string{x.ObjName, x.OsType}
	}
	log.Infoln(&sb)

	t := time.NewTicker(time.Second * time.Duration(cfg.CheckIntervalSec))
	defer t.Stop()

	mainStopCh := make(chan struct{})
	mc := make(chan os.Signal, 4)
	signal.Notify(mc, os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM)
	go func() {
		for {
			<-mc
			close(mainStopCh)
		}
	}()

	for {
		select {
		case <-t.C:
			log.Infoln("start refreshing...")

			reqs := make([]desktop.RequestState, len(pd.DesktopList))
			for i, x := range pd.DesktopList {
				reqs[i].ObjID = x.ObjID
				reqs[i].ObjType = x.ObjType
			}
			s, err := desktop.State(nil, cli, reqs)
			if err != nil {
				log.Warnln("get state err:", err)
				goto RECONN
			}
			for _, x := range s {
				log.Infof("%s [%s]%s status is %s", x.ObjID, mp[x.ObjID][0], x.DesktopState)
				if x.DesktopState == "ACTIVE" {
					continue
				}
				log.Infof("%s [%s]%s do refresh", x.ObjID, mp[x.ObjID][0])
				_, err = desktop.Connect(nil, cli, &desktop.RequestConnect{
					ObjID:                 x.ObjID,
					ObjType:               x.ObjType,
					OsType:                mp[x.ObjID][1],
					DeviceID:              int(cli.Devicetype),
					DeviceCode:            cli.Devicecode,
					DeviceName:            tienyik.DeviceNameEdge,
					SysVersion:            tienyik.DeviceModelMacOS,
					AppVersion:            tienyik.AppVersion,
					HostName:              tienyik.DeviceNameEdge,
					HardwareFeatureCode:   cli.Devicecode,
					SpecifiedCertCategory: 1,
				})
				if err != nil {
					log.Warnln("connect err:", err)
					goto RECONN
				}
			}
		case <-mainStopCh:
			log.Warnln("quit loop")
			return
		}
	}

}
