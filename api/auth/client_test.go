package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"os"
	"testing"

	"github.com/fumiama/tienyik"
	"github.com/fumiama/tienyik/api/cdserv"
	"github.com/fumiama/tienyik/hcli"
	"github.com/sirupsen/logrus"
)

func TestNegotiationEncKey(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)

	k, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	tyr := (*tienyik.RSA)(k)
	tyr.E = 0x010001

	r, err := NegotiationEncKey(&RequestNegotiationEncKey{
		CertData: tyr.PublicKeyToSPKI(),
		CertType: tienyik.ETPYE_AES_CBC,
		Etype:    tienyik.ETPYE_AES_CBC,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("EncData: %s", r.EncData)
	t.Logf("EncKey: %s", r.EncKey)

	_, err = r.Unwrap(tyr)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLogin(t *testing.T) {
	cli := hcli.NewClient()
	sd, err := cdserv.GetServData()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("get serv data:", sd)
	x, err := GenChallengeData(nil, cli)
	if err != nil {
		t.Fatal(err)
	}
	sd.SetClient(cli)
	rsp, err := Login(nil, cli, &RequestLogin{
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
	err = Logout(nil, cli)
	if err != nil {
		t.Fatal(err)
	}
}
