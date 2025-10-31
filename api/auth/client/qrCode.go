package client

import (
	"net/url"

	"github.com/fumiama/tienyik"
	"github.com/fumiama/tienyik/hcli"
	"github.com/fumiama/tienyik/internal/hson"
	"github.com/fumiama/tienyik/internal/textio"
)

type ResponseGenData struct {
	QrCodeId       string  `json:"qrCodeId"`
	QrCodeEndpoint *string `json:"qrCodeEndpoint"`
	ServerHost     string  `json:"serverHost"`
}

func GenData(tya *tienyik.AES, cli *hcli.Client) (*ResponseGenData, error) {
	resp, err := cli.Post(textio.API(), textio.ContenTypeForm, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return hson.Unmarshal[*ResponseGenData](tya, resp.Body)
}

type ResponseGetStatusData struct {
	CodeId           string  `json:"codeId"`
	CodeStatus       string  `json:"codeStatus"`
	LoginToken       *string `json:"loginToken"`
	ClientCustomData *string `json:"clientCustomData"`
	SupplementData   *string `json:"supplementData"`
}

func GetStatus(tya *tienyik.AES, cli *hcli.Client, qrCodeId string) (*ResponseGetStatusData, error) {
	u, err := url.Parse(textio.API())
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("qrCodeId", qrCodeId)
	u.RawQuery = tya.EUrlParams(q)

	resp, err := cli.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return hson.Unmarshal[*ResponseGetStatusData](tya, resp.Body)
}
