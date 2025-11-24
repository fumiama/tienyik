package tienyik

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"strings"
)

const (
	Version    = 103010001
	AppVersion = "3.1.0"
)

const (
	DeviceTypePC  = 25
	DeviceTypeMAC = 45
	DeviceTypeWEB = 60
)

const (
	AppModelTOC   = "1"
	AppModelTOB   = "2"
	AppModelPHONE = "3"
)

const (
	DeviceNameEdge = "Edge浏览器"
)

// alos sysVersion
const (
	DeviceModelMacOS = "Macintosh; Intel Mac OS X 10_15_7"
)

const (
	ArchX86 = "2001"
	ArchARM = "2002"
	ArchHW  = "2003"
)

func NewDeviceCode() string {
	sb := &strings.Builder{}
	sb.WriteString("web_")
	enc := base64.NewEncoder(base64.RawURLEncoding, sb)
	io.CopyN(enc, rand.Reader, (32/6+1)*8)
	enc.Close()
	return sb.String()[:4+32]
}
