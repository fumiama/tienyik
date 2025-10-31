package auth

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/fumiama/tienyik"
	"github.com/fumiama/tienyik/hcli"
	"github.com/fumiama/tienyik/internal/horm"
	"github.com/fumiama/tienyik/internal/hson"
	"github.com/fumiama/tienyik/internal/textio"
)

type RequestNegotiationEncKey struct {
	CertData string `json:"certData"`
	CertType string `json:"certType"`
	Etype    string `json:"etype"`
}

type ResponseNegotiationEncKey struct {
	EncData string `json:"encData"`
	EncKey  string `json:"encKey"`
}

type ResponseNegotiationEncKeyData struct {
	EValue string `json:"evalue"`
	EID    string `json:"eid"`
}

func (r *ResponseNegotiationEncKey) Unwrap(tyr *tienyik.RSA) (*ResponseNegotiationEncKeyData, error) {
	v, err := base64.StdEncoding.DecodeString(r.EncKey)
	if err != nil {
		return nil, err
	}
	aesk, err := tyr.Decrypt(v)
	if err != nil {
		return nil, err
	}

	v, err = base64.StdEncoding.DecodeString(r.EncData)
	if err != nil {
		return nil, err
	}
	v, err = tienyik.NewAES(aesk).Decrypt(v)
	if err != nil {
		return nil, err
	}

	var rsp ResponseNegotiationEncKeyData
	err = json.Unmarshal(v, &rsp)
	if err != nil {
		return nil, err
	}
	return &rsp, nil
}

func NegotiationEncKey(r *RequestNegotiationEncKey) (*ResponseNegotiationEncKey, error) {
	resp, err := hcli.NoClient.Post(
		textio.API(), textio.ContenTypeJSON,
		bytes.NewReader(hson.Marshal(nil, r)),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return hson.Unmarshal[*ResponseNegotiationEncKey](nil, resp.Body)
}

type ResponseGenChallengeData struct {
	EffectiveSeconds int    `json:"effectiveSeconds"`
	ChallengeID      string `json:"challengeId"`
	ChallengeCode    string `json:"challengeCode"`
}

func GenChallengeData(tya *tienyik.AES, cli *hcli.Client) (*ResponseGenChallengeData, error) {
	resp, err := cli.Post(
		textio.API(), textio.ContenTypeJSON,
		strings.NewReader("{}"),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return hson.Unmarshal[*ResponseGenChallengeData](tya, resp.Body)
}

type RequestLogin struct {
	UserAccount    string `form:"userAccount"`
	Password       string `form:"password"`
	SHA256Password string `form:"sha256Password"`
	ChallengeID    string `form:"challengeId"`
	DeviceCode     string `form:"deviceCode"`
	DeviceName     string `form:"deviceName"`
	DeviceType     string `form:"deviceType"`
	DeviceModel    string `form:"deviceModel"`
	AppVersion     string `form:"appVersion"`
	SysVersion     string `form:"sysVersion"`
	ClientVersion  string `form:"clientVersion"`
}

type ResponseLogin struct {
	Timestamp             int64   `json:"timestamp"`
	UserID                int64   `json:"userId"`
	UserEid               string  `json:"userEid"`
	UserName              string  `json:"userName"`
	UserAccount           string  `json:"userAccount"`
	Email                 *string `json:"email"`
	Mobilephone           string  `json:"mobilephone"`
	TenantID              int64   `json:"tenantId"`
	SecretKey             string  `json:"secretKey"`
	DeviceType            string  `json:"deviceType"`
	TenantName            string  `json:"tenantName"`
	BondedDevice          bool    `json:"bondedDevice"`
	RealNameStatus        int64   `json:"realNameStatus"`
	HasPassword           int64   `json:"hasPassword"`
	PID                   *string `json:"pid"`
	NeedSmsValidate       bool    `json:"needSmsValidate"`
	NeedUpdatePassword    bool    `json:"needUpdatePassword"`
	ForceUpdateInitialPwd bool    `json:"forceUpdateInitialPwd"`
	AdminUser             bool    `json:"adminUser"`
	LoginNotify           *string `json:"loginNotify"`
	CommonLoginReqHeader  string  `json:"commonLoginReqHeader"`
	RecordSn              bool    `json:"recordSn"`
	Token                 *string `json:"token"`
	TwoFaValidateType     *string `json:"twoFaValidateType"`
	Random                *string `json:"random"`
	NeedBindVirtualMfa    *bool   `json:"needBindVirtualMfa"`
	Tryout                *string `json:"tryout"`
	CtqEncID              *string `json:"ctqEncId"`
}

func (r *ResponseLogin) SetClient(cli *hcli.Client) {
	cli.Tenantid = strconv.FormatInt(r.TenantID, 10)
	cli.Usereid = r.UserEid
}

func Login(tya *tienyik.AES, cli *hcli.Client, r *RequestLogin) (*ResponseLogin, error) {
	resp, err := cli.Post(
		textio.API(), textio.ContenTypeForm,
		bytes.NewReader(horm.Marshal(tya, r)),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return hson.Unmarshal[*ResponseLogin](tya, resp.Body)
}
