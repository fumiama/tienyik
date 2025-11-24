package hcli

import (
	"context"
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
	sg         tienyik.Signer
	secretKey  string
	offsetTime int64

	Appmodel   string
	Devicecode string
	Devicetype uint64
	Servernode string
	Tenantid   string
	Usereid    string
	Version    uint64
}

func NewClient() *Client {
	return &Client{
		sg:         tienyik.NewSigner(context.TODO()),
		Appmodel:   tienyik.AppModelTOB,
		Devicecode: tienyik.NewDeviceCode(),
		Devicetype: tienyik.DeviceTypeWEB,
		Version:    tienyik.Version,
	}
}

func (c *Client) SetSecretKey(k string) {
	c.secretKey = k
}

func (c *Client) SetTimestamp(ts int64) {
	n := time.Now().UnixMilli()
	c.offsetTime = n - ts
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
		if fieldValue.IsZero() {
			continue
		}
		k := base14.DecodeString("廝呲舀㴄") + field.Name

		switch fieldValue.Kind() {
		case reflect.String:
			req.Header.Set(k, fieldValue.String())
		case reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8, reflect.Uint, reflect.Uintptr:
			req.Header.Set(k, strconv.FormatUint(fieldValue.Uint(), 10))
		case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int:
			req.Header.Set(k, strconv.FormatInt(fieldValue.Int(), 10))
		default:
			panic("unsupported field " + field.Name + " value type " + fieldValue.Type().String())
		}
	}

	if c.Appmodel != "" {
		ts := time.Now().UnixMilli()
		rid := uint64(atomic.AddUintptr(&c.rcnt, 1)) + uint64(ts)
		requestid := strconv.FormatUint(rid, 10)

		ts -= c.offsetTime
		timestamp := strconv.FormatInt(ts, 10)

		req.Header.Set(base14.DecodeString("廝呲草獱歙攷徥爀㴆"), requestid)
		req.Header.Set(base14.DecodeString("廝呲荑睭杜蕆厵縀㴆"), timestamp)

		if c.secretKey != "" {
			req.Header.Set(base14.DecodeString("廝呲荍睧榘敇揉獳欜渀㴂"), c.sg.GenKeyNew(
				context.TODO(), c.Devicetype, uint64(ts), rid, c.secretKey,
				c.Usereid, req.URL.EscapedPath(), c.Servernode, c.Version,
			))
		}
	}
}
