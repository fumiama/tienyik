package hcli

import (
	"io"
	"net/http"

	base14 "github.com/fumiama/go-base16384"
)

func setCommonHeaders(req *http.Request, contentType string) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Pragma", "no-cache")
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	req.Header.Set("Origin", base14.DecodeString("栝啇俌蠯姜吲融艹歛烦宸㴅"))
	req.Header.Set("Referer", base14.DecodeString("栝啇俌蠯姜吲融艹歛烦宸紀㴆"))
	req.Header.Set(
		"User-Agent",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36 Edg/141.0.0.0",
	)
}

func (cli *Client) Get(path string) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodGet, ep(path), nil)
	if err != nil {
		return nil, err
	}
	setCommonHeaders(req, "")
	cli.setHeaders(req)
	return DefaultClient.Do(req)
}

func (cli *Client) Post(path string, contentType string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodPost, ep(path), body)
	if err != nil {
		return nil, err
	}
	setCommonHeaders(req, contentType)
	cli.setHeaders(req)
	return DefaultClient.Do(req)
}

func (cli *Client) Put(path string, contentType string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodPut, ep(path), body)
	if err != nil {
		return nil, err
	}
	setCommonHeaders(req, contentType)
	cli.setHeaders(req)
	return DefaultClient.Do(req)
}

func (cli *Client) Delete(path string) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodDelete, ep(path), nil)
	if err != nil {
		return nil, err
	}
	setCommonHeaders(req, "")
	cli.setHeaders(req)
	return DefaultClient.Do(req)
}

func (cli *Client) Patch(path string, contentType string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodPatch, ep(path), body)
	if err != nil {
		return nil, err
	}
	setCommonHeaders(req, contentType)
	cli.setHeaders(req)
	return DefaultClient.Do(req)
}
