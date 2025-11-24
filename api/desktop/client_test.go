package desktop

import (
	"os"
	"testing"

	"github.com/fumiama/tienyik"
	"github.com/fumiama/tienyik/api/auth"
	"github.com/fumiama/tienyik/api/cdserv"
	"github.com/fumiama/tienyik/hcli"
)

func TestDesktop(t *testing.T) {
	cli := hcli.NewClient()
	sd, err := cdserv.GetServData()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("get serv data:", sd)
	x, err := auth.GenChallengeData(nil, cli)
	if err != nil {
		t.Fatal(err)
	}
	sd.SetClient(cli)
	rsp, err := auth.Login(nil, cli, &auth.RequestLogin{
		UserAccount:    os.Getenv("TYUSR"),
		Password:       tienyik.ChallengePassword(os.Getenv("TYPWD"), x.ChallengeCode),
		SHA256Password: tienyik.ChallengeSHA256Password(os.Getenv("TYPWD"), x.ChallengeCode),
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
		t.Fatal(err)
	}
	t.Log(rsp)
	rsp.SetClient(cli)
	pd, err := PageDesktop(nil, cli, &RequestPageDesktop{
		GetCnt:       1,
		DesktopTypes: []string{"1", tienyik.ArchX86, tienyik.ArchARM, tienyik.ArchHW},
		SortType:     DefaultRequestPageDesktopSortType,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(pd)
	for _, x := range pd.DesktopList {
		feat, err := Feature(nil, cli, x.DesktopID, x.ObjType, x.ObjID)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("feat:", feat)
		ext, err := GetDesktopExtraInfo(nil, cli, x.ObjID, x.ObjType)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("ext:", ext)
		s, err := State(nil, cli, []RequestState{{
			ObjID:   x.ObjID,
			ObjType: x.ObjType,
		}})
		if err != nil {
			t.Fatal(err)
		}
		t.Log("s1:", s)
		con, err := Connect(nil, cli, &RequestConnect{
			ObjID:                 x.ObjID,
			ObjType:               x.ObjType,
			OsType:                x.OsType,
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
			t.Fatal(err)
		}
		t.Log("con:", con)
		s, err = State(nil, cli, []RequestState{{
			ObjID:   x.ObjID,
			ObjType: x.ObjType,
		}})
		if err != nil {
			t.Fatal(err)
		}
		t.Log("s2:", s)
	}
	err = auth.Logout(nil, cli)
	if err != nil {
		t.Fatal(err)
	}
}
