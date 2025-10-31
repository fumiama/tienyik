package hson

import (
	"bytes"
	"encoding/base64"
	"encoding/json"

	"github.com/fumiama/tienyik"
	"github.com/fumiama/tienyik/internal/log"
	"github.com/fumiama/tienyik/internal/op"
)

type reqbody struct {
	EData string `json:"edata"`
}

func Marshal(tya *tienyik.AES, v any) []byte {
	w := bytes.NewBuffer(make([]byte, 0, 1024))
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		panic(err)
	}
	log.Debugln("Marshal JSON:", w.String())
	if tya != nil {
		return tya.Encrypt(op.Must(json.Marshal(&reqbody{
			EData: base64.StdEncoding.EncodeToString(w.Bytes()),
		})))
	}
	return w.Bytes()
}
