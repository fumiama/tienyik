package client

import (
	"testing"

	"github.com/fumiama/tienyik/hcli"
)

func TestGenData(t *testing.T) {
	r, err := GenData(nil, hcli.NewClient())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	if r.QrCodeId == "" {
		t.Fail()
	}
}
