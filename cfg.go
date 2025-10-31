package tienyik

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"strings"
)

const Version = "103010001"

const (
	DeviceTypePC  = "25"
	DeviceTypeMAC = "45"
	DeviceTypeWEB = "60"
)

const (
	AppModelTOC   = "1"
	AppModelTOB   = "2"
	AppModelPHONE = "3"
)

func NewDeviceCode() string {
	sb := &strings.Builder{}
	sb.WriteString("web_")
	enc := base64.NewEncoder(base64.RawURLEncoding, sb)
	io.CopyN(enc, rand.Reader, (32/6+1)*8)
	enc.Close()
	return sb.String()[:4+32]
}
