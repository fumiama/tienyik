package hcli

import (
	"net/http"
	"reflect"
	"strconv"
	"sync/atomic"
	"time"

	base14 "github.com/fumiama/go-base16384"
	"github.com/fumiama/tienyik"
	"golang.org/x/net/http2"
)

var NoClient = (*Client)(nil)

// DefaultClient is the default HTTP2 client.
var DefaultClient = http.Client{
	Transport: &http2.Transport{},
}

type Client struct {
	rcnt       uintptr
	Appmodel   string
	Devicecode string
	Devicetype string
	Servernode string
	Tenantid   string
	Usereid    string
	Version    string
}

func NewClient() *Client {
	return &Client{
		Appmodel:   tienyik.AppModelTOB,
		Devicecode: tienyik.NewDeviceCode(),
		Devicetype: tienyik.DeviceTypeWEB,
		Version:    tienyik.Version,
	}
}

func (c *Client) setHeaders(req *http.Request) {
	if c == nil {
		return
	}

	v := reflect.ValueOf(c).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		if !field.IsExported() {
			continue
		}

		fieldValue := v.Field(i)

		if fieldValue.Kind() == reflect.String {
			req.Header.Set(
				base14.DecodeString("廝呲舀㴄")+field.Name,
				fieldValue.String(),
			)
		}
	}

	if c.Appmodel != "" {
		ts := time.Now().UnixMilli()
		timestamp := strconv.FormatInt(ts, 10)
		requestid := strconv.FormatUint(
			uint64(atomic.AddUintptr(&c.rcnt, 1))+uint64(ts), 10,
		)

		req.Header.Set(base14.DecodeString("廝呲草獱歙攷徥爀㴆"), requestid)
		req.Header.Set(base14.DecodeString("廝呲荑睭杜蕆厵縀㴆"), timestamp)
	}

	if c.Servernode != "" {
		//TODO: gensign
	}
}
