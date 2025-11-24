package horm

import (
	"net/url"
	"reflect"
	"strconv"

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

		switch field.Kind() {
		case reflect.String:
			q.Set(formTag, field.String())
		case reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8, reflect.Uint, reflect.Uintptr:
			q.Set(formTag, strconv.FormatUint(field.Uint(), 10))
		case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int:
			q.Set(formTag, strconv.FormatInt(field.Int(), 10))
		}
	}

	s := q.Encode()
	if tya != nil {
		return tya.Encrypt(textio.StringToBytes(s))
	}
	return textio.StringToBytes(s)
}
