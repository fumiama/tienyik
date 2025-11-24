package desktop

import (
	"bytes"
	"net/url"
	"strconv"

	"github.com/fumiama/tienyik"
	"github.com/fumiama/tienyik/hcli"
	"github.com/fumiama/tienyik/internal/horm"
	"github.com/fumiama/tienyik/internal/hson"
	"github.com/fumiama/tienyik/internal/textio"
)

const (
	DefaultRequestPageDesktopSortType = "createTimeV1"
)

type RequestPageDesktop struct {
	GetCnt       int      `json:"getCnt"`
	DesktopTypes []string `json:"desktopTypes"`
	SortType     string   `json:"sortType"`
}

type ResponsePageDesktop struct {
	Timestamp int64 `json:"timestamp"`
	SortList  []struct {
		ObjID        string   `json:"objId"`
		ObjType      int      `json:"objType"`
		ObjValue     string   `json:"objValue"`
		DesktopTypes []string `json:"desktopTypes"`
	} `json:"sortList"`
	DesktopPoolList []any `json:"desktopPoolList"`
	DesktopList     []struct {
		ObjType            int      `json:"objType"`
		TenantID           int      `json:"tenantId"`
		ObjID              string   `json:"objId"`
		ConnectURL         []string `json:"connectUrl"`
		ObjName            string   `json:"objName"`
		Backupurl          []string `json:"backupurl"`
		OsType             string   `json:"osType"`
		OsName             string   `json:"osName"`
		ConnectMaster      int      `json:"connectMaster"`
		NeedLineUp         bool     `json:"needLineUp"`
		UserDesktopGroupID any      `json:"userDesktopGroupId"`
		Strategy           struct {
			ReconnectMsg        any    `json:"reconnectMsg"`
			RebootMsg           any    `json:"rebootMsg"`
			ShutoffMsg          any    `json:"shutoffMsg"`
			ShutdownStrategy    any    `json:"shutdownStrategy"`
			RebootStrategy      any    `json:"rebootStrategy"`
			ModifyComputerAllas string `json:"modifyComputerAllas"`
			CheckBeforeConnect  any    `json:"checkBeforeConnect"`
		} `json:"strategy"`
		CloudMobileType         any    `json:"cloudMobileType"`
		DesktopID               string `json:"desktopId"`
		DesktopName             string `json:"desktopName"`
		FlavorName              any    `json:"flavorName"`
		ImageName               string `json:"imageName"`
		OsBit                   string `json:"osBit"`
		CPUCore                 any    `json:"cpuCore"`
		MemoryGB                any    `json:"memoryGB"`
		RootDiskGB              any    `json:"rootDiskGB"`
		DataDiskGB              any    `json:"dataDiskGB"`
		Status                  string `json:"status"`
		Summary                 any    `json:"summary"`
		TanentCode              string `json:"tanentCode"`
		TanentName              string `json:"tanentName"`
		UseStatus               string `json:"useStatus"`
		DesktopCode             string `json:"desktopCode"`
		ForeignDesktopID        string `json:"foreignDesktopId"`
		ForbiddenConnect        bool   `json:"forbiddenConnect"`
		GpuType                 bool   `json:"gpuType"`
		GpuVirtualMethod        any    `json:"gpuVirtualMethod"`
		UserMode                int    `json:"userMode"`
		DefaultDesktop          bool   `json:"defaultDesktop"`
		ExpireDate              any    `json:"expireDate"`
		CreateDate              int64  `json:"createDate"`
		NowDate                 any    `json:"nowDate"`
		NoticeInterval          int    `json:"noticeInterval"`
		BandExpireDate          any    `json:"bandExpireDate"`
		BandNoticeInterval      int    `json:"bandNoticeInterval"`
		UpperResolution         any    `json:"upperResolution"`
		ProdType                string `json:"prodType"`
		ProdGroupType           int    `json:"prodGroupType"`
		ProdInstID              string `json:"prodInstId"`
		ProdGroupName           string `json:"prodGroupName"`
		DesktopMirrorTagSet     []any  `json:"desktopMirrorTagSet"`
		LicenseExpireDate       int64  `json:"licenseExpireDate"`
		LicenseNoticeInterval   int    `json:"licenseNoticeInterval"`
		AllowConnStartTime      any    `json:"allowConnStartTime"`
		AllowConnEndTime        any    `json:"allowConnEndTime"`
		OperationAuditSupported bool   `json:"operationAuditSupported"`
		ProjectionScreenState   any    `json:"projectionScreenState"`
		NickName                string `json:"nickName"`
		VMType                  int    `json:"vmType"`
		PayType                 string `json:"payType"`
		ProdSubType             any    `json:"prodSubType"`
		InstStatus              int    `json:"instStatus"`
		UseTimeVO               any    `json:"useTimeVO"`
		UseStatusShowActions    any    `json:"useStatusShowActions"`
		UseStatusText           string `json:"useStatusText"`
		UseStatusColor          string `json:"useStatusColor"`
		ModifyComputerAllas     any    `json:"modifyComputerAllas"`
		UsePrivateImageFile     bool   `json:"usePrivateImageFile"`
		ImageID                 int    `json:"imageId"`
		ImageCategoryID         int    `json:"imageCategoryId"`
		HaProdType              int    `json:"haProdType"`
		CtrlTypes               []any  `json:"ctrlTypes"`
		OrderProductData        struct {
			TimeLimitTotal  any `json:"timeLimitTotal"`
			TimeLimitUsed   any `json:"timeLimitUsed"`
			NextAcctTime    any `json:"nextAcctTime"`
			BusiChannelType any `json:"busiChannelType"`
			ManageData      any `json:"manageData"`
			ActiveDate      any `json:"activeDate"`
			KeepTime        any `json:"keepTime"`
		} `json:"orderProductData"`
		LicenseID     int    `json:"licenseId"`
		RegionID      int    `json:"regionId"`
		TenantCode    string `json:"tenantCode"`
		ConnectAPIURL struct {
			ConnectPath string `json:"connectPath"`
			StatusPath  string `json:"statusPath"`
			StatePath   string `json:"statePath"`
			AppendData  any    `json:"appendData"`
		} `json:"connectApiUrl"`
	} `json:"desktopList"`
	PreemptionDesktopList []any `json:"preemptionDesktopList"`
}

func PageDesktop(tya *tienyik.AES, cli *hcli.Client, r *RequestPageDesktop) (*ResponsePageDesktop, error) {
	resp, err := cli.Post(
		textio.API(), textio.ContenTypeJSON,
		bytes.NewReader(hson.Marshal(tya, r)),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return hson.Unmarshal[*ResponsePageDesktop](tya, resp.Body)
}

type ResponseFeature struct {
	CPUCore      int `json:"cpuCore"`
	MemoryGB     int `json:"memoryGB"`
	SystemDiskGB int `json:"systemDiskGB"`
	DataDiskGB   any `json:"dataDiskGB"`
	TotalDiskGB  int `json:"totalDiskGB"`
	SysDisk      struct {
		Size int    `json:"size"`
		Path string `json:"path"`
		Code string `json:"code"`
	} `json:"sysDisk"`
	DataDiskList       []any  `json:"dataDiskList"`
	MirrorVersion      string `json:"mirrorVersion"`
	MirrorCategoryName string `json:"mirrorCategoryName"`
	DesktopName        string `json:"desktopName"`
	GpuSliceRAM        any    `json:"gpuSliceRam"`
	GpuSliceRAMDesc    any    `json:"gpuSliceRamDesc"`
	ExpireDate         any    `json:"expireDate"`
	CreateDate         int64  `json:"createDate"`
	NowDate            int64  `json:"nowDate"`
	LinkInfo           any    `json:"linkInfo"`
	OrderProduct       struct {
		TimeLimitTotal  any    `json:"timeLimitTotal"`
		TimeLimitUsed   any    `json:"timeLimitUsed"`
		NextAcctTime    any    `json:"nextAcctTime"`
		BusiChannelType string `json:"busiChannelType"`
		ManageData      any    `json:"manageData"`
		ActiveDate      any    `json:"activeDate"`
		KeepTime        any    `json:"keepTime"`
	} `json:"orderProduct"`
	MirrorID         int    `json:"mirrorId"`
	MirrorCategoryID int    `json:"mirrorCategoryId"`
	ProductName      string `json:"productName"`
	PayType          string `json:"payType"`
	ProdType         string `json:"prodType"`
	ProdSubType      any    `json:"prodSubType"`
	TimePkgVOS       []any  `json:"timePkgVOS"`
	Os               string `json:"os"`
	MirrorType       string `json:"mirrorType"`
}

func Feature(tya *tienyik.AES, cli *hcli.Client, desktopId string, objType int, objId string) (*ResponseFeature, error) {
	u, err := url.Parse(textio.API())
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("desktopId", desktopId)
	q.Set("objType", strconv.Itoa(objType))
	q.Set("objId", objId)
	u.RawQuery = tya.EUrlParams(q)

	resp, err := cli.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return hson.Unmarshal[*ResponseFeature](tya, resp.Body)
}

type ResponseGetDesktopExtraInfo struct {
	Strategy struct {
		ReconnectMsg        any    `json:"reconnectMsg"`
		RebootMsg           any    `json:"rebootMsg"`
		ShutoffMsg          any    `json:"shutoffMsg"`
		ShutdownStrategy    string `json:"shutdownStrategy"`
		RebootStrategy      string `json:"rebootStrategy"`
		ModifyComputerAllas any    `json:"modifyComputerAllas"`
		CheckBeforeConnect  any    `json:"checkBeforeConnect"`
	} `json:"strategy"`
	UpperResolution      string `json:"upperResolution"`
	TimeLimitProductData any    `json:"timeLimitProductData"`
	HaProdType           int    `json:"haProdType"`
}

func GetDesktopExtraInfo(tya *tienyik.AES, cli *hcli.Client, objId string, objType int) (*ResponseGetDesktopExtraInfo, error) {
	u, err := url.Parse(textio.API())
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("objId", objId)
	q.Set("objType", strconv.Itoa(objType))
	u.RawQuery = tya.EUrlParams(q)

	resp, err := cli.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return hson.Unmarshal[*ResponseGetDesktopExtraInfo](tya, resp.Body)
}

type RequestConnect struct {
	ObjID                 string `form:"objId"`
	ObjType               int    `form:"objType"`
	OsType                string `form:"osType"`
	DeviceID              int    `form:"deviceId"`
	DeviceCode            string `form:"deviceCode"`
	DeviceName            string `form:"deviceName"`
	SysVersion            string `form:"sysVersion"`
	AppVersion            string `form:"appVersion"`
	HostName              string `form:"hostName"`
	VdCommand             string `form:"vdCommand"`
	IPAddress             string `form:"ipAddress"`
	MacAddress            string `form:"macAddress"`
	HardwareFeatureCode   string `form:"hardwareFeatureCode"`
	SpecifiedCertCategory int    `form:"specifiedCertCategory"`
}

type ResponseConnect struct {
	GoingRetry        bool `json:"goingRetry"`
	DesktopInfo       any  `json:"desktopInfo"`
	ShadowDesktopInfo struct {
		InHaMode         int `json:"inHaMode"`
		HaDesktopID      any `json:"haDesktopId"`
		HaConnectingTips any `json:"haConnectingTips"`
		HaConnectSucTips any `json:"haConnectSucTips"`
	} `json:"shadowDesktopInfo"`
	DesktopAnywhereInfo struct {
		AnywhereStatus       int  `json:"anywhereStatus"`
		MigrateStatus        any  `json:"migrateStatus"`
		AnywhereDesktopID    any  `json:"anywhereDesktopId"`
		SrcResPoolName       any  `json:"srcResPoolName"`
		TargetResPoolName    any  `json:"targetResPoolName"`
		EstimatedTime        int  `json:"estimatedTime"`
		ReminderDays         any  `json:"reminderDays"`
		RoamingDays          any  `json:"roamingDays"`
		NeedReserveRemind    int  `json:"needReserveRemind"`
		ShadowInfoDTO        any  `json:"shadowInfoDTO"`
		AnywhereOpen         bool `json:"anywhereOpen"`
		ConnectTargetDesktop bool `json:"connectTargetDesktop"`
	} `json:"desktopAnywhereInfo"`
	DesktopID  string `json:"desktopId"`
	PollingKey string `json:"pollingKey"`
	AuthInfo   any    `json:"authInfo"`
}

func Connect(tya *tienyik.AES, cli *hcli.Client, r *RequestConnect) (*ResponseConnect, error) {
	resp, err := cli.Post(
		textio.API(), textio.ContenTypeForm,
		bytes.NewReader(horm.Marshal(tya, r)),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return hson.Unmarshal[*ResponseConnect](tya, resp.Body)
}

type RequestState struct {
	ObjID   string `json:"objId"`
	ObjType int    `json:"objType"`
}

type ResponseState struct {
	ObjType         int    `json:"objType"`
	ObjID           string `json:"objId"`
	DesktopID       int    `json:"desktopId"`
	DesktopState    string `json:"desktopState"`
	RunningTask     int    `json:"runningTask"`
	RunningTaskName string `json:"runningTaskName"`
	TaskStartTime   int64  `json:"taskStartTime"`
	MirrorReady     any    `json:"mirrorReady"`
	UseStatus       string `json:"useStatus"`
	UseStatusText   string `json:"useStatusText"`
	UseStatusColor  string `json:"useStatusColor"`
}

func State(tya *tienyik.AES, cli *hcli.Client, r []RequestState) ([]ResponseState, error) {
	resp, err := cli.Post(
		textio.API(), textio.ContenTypeJSON,
		bytes.NewReader(hson.Marshal(tya, &r)),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return hson.Unmarshal[[]ResponseState](tya, resp.Body)
}
