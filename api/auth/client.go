package auth

import (
	"bytes"

	"github.com/fumiama/tienyik/internal/hcli"
	"github.com/fumiama/tienyik/internal/hson"
	"github.com/fumiama/tienyik/internal/textio"
)

type RequestNegotiationEncKey struct {
	CertData string `json:"certData"`
	CertType string `json:"certType"`
	Etype    string `json:"etype"`
}

type ResponseNegotiationEncKey struct {
	EncData string `json:"encData"`
	EncKey  string `json:"encKey"`
}

func NegotiationEncKey(r *RequestNegotiationEncKey) (*ResponseNegotiationEncKey, error) {
	resp, err := hcli.Post(textio.API(), bytes.NewReader(hson.Marshal(nil, r)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return hson.Unmarshal[*ResponseNegotiationEncKey](nil, resp.Body)
}
