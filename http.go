package tienyik

import (
	"errors"
	"net/url"

	"github.com/fumiama/tienyik/internal/textio"
)

func (tya *AES) EUrlParams(params url.Values) string {
	if tya == nil {
		return params.Encode()
	}
	return url.Values{
		textio.FuncName(1, true): {textio.BytesToString(tya.Encrypt(
			textio.StringToBytes(params.Encode()),
		))},
	}.Encode()
}

func (tya *AES) ParseQuery(eparams string) (url.Values, error) {
	if tya == nil {
		return url.ParseQuery(eparams)
	}
	q, err := url.ParseQuery(eparams)
	if err != nil {
		return nil, err
	}
	if len(q) != 1 {
		return nil, errors.New("len(q) must be 1")
	}
	for _, v := range q {
		dec, err := tya.Decrypt(textio.StringToBytes(v[0]))
		if err != nil {
			return nil, err
		}
		return url.ParseQuery(textio.BytesToString(dec))
	}
	panic("unexpected")
}
