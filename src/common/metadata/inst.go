package metadata

import (
	"sort"

	"configcenter/src/common/util"
)

type SetInst struct {
	SetID     int64  `bson:"set_id"`
	SetName   string `bson:"set_name"`
	SetStatus string `bson:"service_status"`
	SetEnv    string `bson:"set_env"`
}
type ModuleInst struct {
	BizID      int64  `bson:"biz_id"`
	ModuleID   int64  `bson:"module_id"`
	ModuleName string `bson:"module_name"`
}
type BizInst struct {
	BizID           int64  `bson:"biz_id"`
	BizName         string `bson:"biz_name"`
	SupplierID      int64  `bson:"supplier_id"`
	SupplierAccount string `bson:"org_id"`
}
type CloudInst struct {
	CloudID   int64  `bson:"cloud_id"`
	CloudName string `bson:"cloud_name"`
}
type ProcessInst struct {
	ProcessID       int64  `json:"process_id" bson:"process_id"`               // 进程名称
	ProcessName     string `json:"process_name" bson:"process_name"`           // 进程名称
	BindIP          string `json:"bind_ip" bson:"bind_ip"`                     // 绑定IP, 枚举: [{ID: "1", Name: "127.0.0.1"}, {ID: "2", Name: "0.0.0.0"}, {ID: "3", Name: "第一内网IP"}, {ID: "4", Name: "第一外网IP"}]
	PORT            string `json:"port" bson:"port"`                           // 端口, 单个端口："8080", 多个连续端口："8080-8089", 多个不连续端口："8080-8089,8199"
	PROTOCOL        string `json:"protocol" bson:"protocol"`                   // 协议, 枚举: [{ID: "1", Name: "TCP"}, {ID: "2", Name: "UDP"}],
	FuncID          string `json:"func_id" bson:"func_id"`                     // 功能ID
	FuncName        string `json:"func_name" bson:"func_name"`                 // 功能名称
	StartParamRegex string `json:"start_param_regex" bson:"start_param_regex"` // 启动参数匹配规则
}

type HostIdentifier struct {
	HostID          int64                       `json:"host_id" bson:"host_id"`           // 主机ID(host_id)								数字
	HostName        string                      `json:"host_name" bson:"host_name"`       // 主机名称
	SupplierID      int64                       `json:"supplier_id"`                      // 开发商ID（supplier_id）				数字
	SupplierAccount string                      `json:"org_id"`                           // 开发商帐号（org_id）	数字
	CloudID         int64                       `json:"cloud_id" bson:"cloud_id"`         // 所属云区域id(cloud_id)				数字
	CloudName       string                      `json:"cloud_name" bson:"cloud_name"`     // 所属云区域名称(cloud_name)		字符串（最大长度25）
	InnerIP         string                      `json:"host_innerip" bson:"host_innerip"` // 内网IP
	OuterIP         string                      `json:"host_outerip" bson:"host_outerip"` // 外网IP
	OSType          string                      `json:"os_type" bson:"os_type"`           // 操作系统类型
	OSName          string                      `json:"os_name" bson:"os_name"`           // 操作系统名称
	Memory          int64                       `json:"mem" bson:"mem"`                   // 内存容量
	CPU             int64                       `json:"cpu" bson:"cpu"`                   // CPU逻辑核心数
	Disk            int64                       `json:"disk" bson:"disk"`                 // 磁盘容量
	HostIdentModule map[string]*HostIdentModule `json:"associations" bson:"associations"`
	Process         []HostIdentProcess          `json:"process" bson:"process"`
}

type HostIdentProcess struct {
	ProcessID       int64   `json:"process_id" bson:"process_id"`               // 进程名称
	ProcessName     string  `json:"process_name" bson:"process_name"`           // 进程名称
	BindIP          string  `json:"bind_ip" bson:"bind_ip"`                     // 绑定IP, 枚举: [{ID: "1", Name: "127.0.0.1"}, {ID: "2", Name: "0.0.0.0"}, {ID: "3", Name: "第一内网IP"}, {ID: "4", Name: "第一外网IP"}]
	PORT            string  `json:"port" bson:"port"`                           // 端口, 单个端口："8080", 多个连续端口："8080-8089", 多个不连续端口："8080-8089,8199"
	PROTOCOL        string  `json:"protocol" bson:"protocol"`                   // 协议, 枚举: [{ID: "1", Name: "TCP"}, {ID: "2", Name: "UDP"}],
	FuncID          string  `json:"func_id" bson:"func_id"`                     // 功能ID
	FuncName        string  `json:"func_name" bson:"func_name"`                 // 功能名称
	StartParamRegex string  `json:"start_param_regex" bson:"start_param_regex"` // 启动参数匹配规则
	BindModules     []int64 `json:"bind_modules" bson:"bind_modules"`           // 进程绑定的模块ID，数字数组
}

type HostIdentProcessSorter []HostIdentProcess

func (p HostIdentProcessSorter) Len() int      { return len(p) }
func (p HostIdentProcessSorter) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p HostIdentProcessSorter) Less(i, j int) bool {
	sort.Sort(util.Int64Slice(p[i].BindModules))
	return p[i].ProcessID < p[j].ProcessID
}

// HostIdentModule HostIdentifier module define
type HostIdentModule struct {
	BizID      int64  `json:"biz_id"`         // 业务ID
	BizName    string `json:"biz_name"`       // 业务名称
	SetID      int64  `json:"set_id"`         // 所属集群(set_id)：						数字
	SetName    string `json:"set_name"`       // 所属集群名称(set_name)：			字符串（最大长度25）
	ModuleID   int64  `json:"module_id"`      // 所属模块(module_id)：				数字
	ModuleName string `json:"module_name"`    // 所属模块(module_name)：			字符串（最大长度25）
	SetStatus  string `json:"service_status"` // 集群服务状态（bk_set_status）			数字
	SetEnv     string `json:"set_env"`        // 环境类型（bk_set_type）					数字
}

// SearchIdentifierParam defines the param
type SearchIdentifierParam struct {
	IP   IPParam `json:"ip"`
	Page BasePage
}

type IPParam struct {
	Data    []string `json:"data"`
	CloudID *int64   `json:"cloud_id"`
}

type SearchHostIdentifierResult struct {
	BaseResp `json:",inline"`
	Data     struct {
		Count int              `json:"count"`
		Info  []HostIdentifier `json:"info"`
	} `json:"data"`
}
