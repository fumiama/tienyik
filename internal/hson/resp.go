package hson

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/fumiama/tienyik"
	"github.com/fumiama/tienyik/internal/log"
)

type responseBase[T any] struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Data  T      `json:"data"`
	EData string `json:"edata"`
}

func (rb *responseBase[T]) ok() error {
	if rb.Code != 0 {
		return errors.New("[" + strconv.Itoa(rb.Code) + "] " + rb.Msg)
	}
	return nil
}

func Unmarshal[T any](tya *tienyik.AES, r io.Reader) (data T, err error) {
	var rsp responseBase[T]
	err = json.NewDecoder(r).Decode(&rsp)
	if err == nil {
		err = rsp.ok()
	}
	if err != nil {
		return
	}
	if len(rsp.EData) > 0 && tya != nil {
		var d []byte
		d, err = base64.StdEncoding.DecodeString(rsp.EData)
		if err != nil {
			return
		}
		d, err = tya.Decrypt(d)
		if err != nil {
			return
		}
		log.Debugln("decrypted data:", string(d))
		err = json.Unmarshal(d, &rsp)
		if err != nil {
			return
		}
		err = rsp.ok()
	}
	data = rsp.Data
	return
}
