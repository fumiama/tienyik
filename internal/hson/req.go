package hson

import (
	"bytes"
	"encoding/json"

	"github.com/fumiama/tienyik"
	"github.com/fumiama/tienyik/internal/log"
)

func Marshal(tya *tienyik.AES, v any) []byte {
	w := bytes.NewBuffer(make([]byte, 0, 1024))
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		panic(err)
	}
	log.Debugln("Marshal JSON:", w.String())
	if tya != nil {
		return tya.Encrypt(w.Bytes())
	}
	return w.Bytes()
}
