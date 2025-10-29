package tienyik

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"strings"
	"testing"
)

func TestRSANegotiationEncKey(t *testing.T) {
	const (
		priv = "MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQC77sSeEJkai9l7PkWbIw+ayFucW/+hbLCYFV4+2kcffjpyWqC9p9Knmwm7iivFixmdRkYSBkdkEJe1ywoa4n/mBdwXWXHEKomjejzvuqeQTfnf6QdRPK1pjVsglqj130qLnC1zgFFDA9yfGIFu6ke9dtOkq84mGQwdUVUinFCo3d/RKHwYhBGxoPaY0aZvfQlzt7TPqx0CZTj4VtEujjIGkpWptmlDkVm2n09aRBV8MLJn+3fm1+agHvCpJHfllTq+2DM4n20X+FzdppMFFj1iU4GHlDoTHxe+UEJQSEo7BjRis6UGMeRA1YV565xfPHPx2J1lNFpW4QIfQEp5c+gDAgMBAAECggEAWg20EboKY3oYIMJaQFiwpY7UtzwuZn/ar0Wt/5hA9IpcH2fmFntQqhfUthUjnGlnQMHn2cNAemapC+fbU76tYt/z8wxdQ47OnZN5l7ZtjOyQjAbyBq94uVePDzKijA8PfY69CNe4GNDE60em3itNbTB/gi0Bf6gI8hODJC6bSA8WhjeIx2Heg6hfs4CRi6hBYo3CQaQ5nkpzVVQG/IU/7s0vYe2daSiIw4FjedgKeIHKP5Gge4KHDHW/g9FouXqxq6tQoUeiI3squrdX5u2rq4Ef8wdF91xBvQzHcnFZR3XEwMxjBwj8x1jH6U2kMAFEnC9f7teRiPbyrDXRxdXPgQKBgQDxFrixINnlCvX1c9e+6VFktoT4fBjB3qaFAebDrr235yZGNL9FCmNvBlF9ZeHr53Ww063VUTh46lfpCbugy4oKS2LZKnftldkpNmoIUw4dMerk76g142k2wL+GhxTYlShxzrssBvC/BhqM63kecyzdnUUo8pTZuZsIeX4HJJHLUQKBgQDHjmZo/cX5FiD7ooMbf5GyJn6CKO8IMt/JQcxk4eUt0juDFN0PoCMsXAd2PIwgHtIBJPNHoibsdoNFoQhJGx/eha2IH+P5G7ERoOTOiGEfXUUS/UWUnYOGM3oypVvpcMoQQqN/2W4MFiHaYK+/tdK0UTeyaopeqnFYiqh8sM6BEwKBgQDX/i5L6w2sSYygcj+5N9mHLIqnAK9BidOQWGrBqB0q1PmSbpFqLmt1Pajmes/UhRMI8CzOb6zzj6hhDSo/Xft3S6DsxUKa5eSgKrMGcDq151H198yxMuPBfSBaS01e2QtaIwfH4xPvYG2LES/7Gt74UX1zuRmoksQV7Jr/lUDVcQKBgQCsiqJfmzSGpyyDhkMoDogR6hiuP+hVRW+bGyo3+91lXgYY03xD22kuHLBS+g0KZLudQ66ZvEk9YUcleBOq6ioHA7xEG5bIt7nFDUFoliCrtsBXp+d2lS64ZV/91F4BHIWJw6SMkZoGF0jUAY9UCkkRobuvp1DWgzaXoOQHU/RpQwKBgQC7MIrzy8x81m4y5sLX7HGmE607fNZZIg/3AoAE5JDLF4Z2fU5GUSASjU01ddmGB6gkc8Yjxe0JAH7WaZ+tWji2ywNOcJv/rnsdgYTIgFfWaIUF++K55Ouhe2pdGmNAXFdJZIlw7dma2aLjEVzU2IDPHYULXT8ci2jfGvLsOoH+iw=="
		resp = `{
    "code": 0,
    "data": {
        "encData": "gMm1F5m7nboBdxmAI4gkoP2Voc3W8zgUxoR7pUBF5C2jKeCddqNnEIgWCFloU0mzcnCqXBVeyCAiU4NY6FmTimjKaOaxN4taX2bzzRmqW/gBwU4Rtz30SdG2VbPzqLqg6fZpUrFXWgR0oxjENyn1oKya0VVH6kLazQ28dbYI88WTqHWcoG/guiB9cjIAEXjaKRzHOxH80bd1AEUtEJiUfg==",
        "encKey": "A9nj1emiuEpXvEb00TL/iKDKsTonyD/oUbqs5IN/rJtlY58UemUWgYxNSiGDlZazGKhgyi7AnUHguj1prSJqTRKDLKIej/et6xVB1K5rQQI3KnBE6Q7S5OsnAquhV9MuEoSDHTMC6V94vtYhqFVKAl08cCaAE0oYgZRDGpZ6Q30blu7iMrUonohz3bbZ8/qamtDEQsQ8igeK+ZFyMu4x3SrDqEIQfrwbgjCTdAxIrH78/+x1JGwqrD79yD2YiW3531Z8uHI9L+mUEBL4cp4FmfD4sQyUvn2/nhWb2DXbZIMVb41d3mKAV/wJIg/RGFhhRBtSxRv/UgQZVurppyW31Q=="
    }
}`
		deks = "ugajrabhuq1mhaav8uwtuukqjbvej08t"
		dats = `{"evalue":"dwwitacs6whdl2n2fk9btiejlv1jhfnu","eid":"c359e8288500cccaf2d321a92bbe662b8c13558e0d00f50155d83527dae70dd2311ce4460c95495b3000bec80f732ee8"}`
	)

	der, err := base64.StdEncoding.DecodeString(priv)
	if err != nil {
		t.Fatal(err)
	}

	k, err := x509.ParsePKCS8PrivateKey(der)
	if err != nil {
		t.Fatal(err)
	}
	tyr := (*TYRSA)(k.(*rsa.PrivateKey))
	tyr.E = 0x010001

	w := bytes.NewBuffer(make([]byte, 0, 1024))
	json.NewEncoder(w).Encode(&struct {
		CertData string `json:"certData"`
		CertType string `json:"certType"`
		Etype    string `json:"etype"`
	}{
		CertData: tyr.PublicKeyToSPKI(),
		CertType: ETPYE_AES_CBC,
		Etype:    ETPYE_AES_CBC,
	})

	type respd struct {
		EncData string `json:"encData"`
		EncKey  string `json:"encKey"`
	}

	type respb struct {
		Code int   `json:"code"`
		Data respd `json:"data"`
	}

	var body respb
	if err := json.NewDecoder(strings.NewReader(resp)).Decode(&body); err != nil {
		t.Fatal(err)
	}

	if body.Code != 0 {
		t.Fatalf("Expected code 0, got %d", body.Code)
	}

	t.Logf("EncData: %s", body.Data.EncData)
	t.Logf("EncKey: %s", body.Data.EncKey)

	v, err := base64.StdEncoding.DecodeString(body.Data.EncKey)
	if err != nil {
		t.Fatal(err)
	}
	aesk, err := tyr.Decrypt(v)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(aesk))
	if string(aesk) != deks {
		t.Fatal("expect deks", deks, "got", string(aesk))
	}

	v, err = base64.StdEncoding.DecodeString(body.Data.EncData)
	if err != nil {
		t.Fatal(err)
	}
	v, err = NewTYAES([]byte(aesk)).Decrypt(v)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(v))
	if string(v) != dats {
		t.Fatal("expect", dats, "got", string(v))
	}
}
