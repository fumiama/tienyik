package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"

	"github.com/fumiama/tienyik"
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
