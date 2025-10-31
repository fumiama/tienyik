package horm

import (
	"net/url"
	"reflect"

	"github.com/fumiama/tienyik"
	"github.com/fumiama/tienyik/internal/textio"
)

func Marshal(tya *tienyik.AES, x any) []byte {
	q := url.Values{}
	v := reflect.ValueOf(x).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		formTag := fieldType.Tag.Get("form")
		if formTag == "" || formTag == "-" {
			continue
		}

		if field.Kind() == reflect.String {
			q.Set(formTag, field.String())
		}
	}

	s := q.Encode()
	if tya != nil {
		return tya.Encrypt(textio.StringToBytes(s))
	}
	return textio.StringToBytes(s)
}
