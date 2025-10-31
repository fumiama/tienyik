package cdserv

import (
	"github.com/fumiama/tienyik/hcli"
	"github.com/fumiama/tienyik/internal/hson"
	"github.com/fumiama/tienyik/internal/textio"
)

type GlobalSwitches struct {
	BodyMsgEType        []string `json:"bodyMsgEType"`
	PasswordRulesEnable bool     `json:"passwordRulesEnable"`
}

type ResponseGetServData struct {
	GlobalSwitches GlobalSwitches `json:"globalSwitches"`
	ServerNodeId   string         `json:"serverNodeId"`
	Timestamp      int64          `json:"timestamp"`
	NetAccessType  int            `json:"netAccessType"`
}

func (r *ResponseGetServData) SetClient(cli *hcli.Client) {
	cli.Servernode = r.ServerNodeId
}

func GetServData() (*ResponseGetServData, error) {
	resp, err := hcli.NoClient.Get(textio.API())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return hson.Unmarshal[*ResponseGetServData](nil, resp.Body)
}
