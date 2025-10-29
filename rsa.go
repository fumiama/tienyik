package tienyik

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
)

type TYRSA rsa.PrivateKey

func NewTYRSA() *TYRSA {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	privateKey.E = 0x010001

	return (*TYRSA)(privateKey)
}

func (tyr *TYRSA) PublicKeyToSPKI() string {
	spkiBytes, err := x509.MarshalPKIXPublicKey(&tyr.PublicKey)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(spkiBytes)
}

func (tyr *TYRSA) Decrypt(ciphertext []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, (*rsa.PrivateKey)(tyr), ciphertext)
}
