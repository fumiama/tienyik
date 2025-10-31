package cdserv

import "testing"

func TestGetServData(t *testing.T) {
	r, err := GetServData()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	if len(r.GlobalSwitches.BodyMsgEType) == 0 {
		t.Fail()
	}
}
