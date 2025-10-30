package tienyik

import (
	"encoding/base64"
	"encoding/binary"
	"testing"
)

func TestDecrypt(t *testing.T) {
	const (
		aescrypted = "vH40y0ly0DCgEGje5aQSAOQX880RtYMWs08Y46lMI8c="
		aesplain   = "moduleCode=DESKTOP_MSGCENTER"
	)
	var (
		rawkey = []uint32{
			2004378729, 1936745065, 1933079672, 1970627951,
			842425958, 1932686949, 1903374648, 1936290669,
		}
		key [32]byte
	)
	for i, k := range rawkey {
		binary.BigEndian.PutUint32(key[i*4:(i+1)*4], k)
	}
	t.Log(string(key[:])) // wxdispbis8txuueo26ffs2veqs18sism
	d, err := base64.StdEncoding.DecodeString(aescrypted)
	if err != nil {
		t.Fatal(err)
	}

	dat, err := NewAES(key[:]).Decrypt(d)
	if err != nil {
		t.Fatal(err)
	}
	s := string(dat)
	t.Log(s)
	if s != aesplain {
		t.Fail()
	}
}
