package v3v0v8

import (
	"time"

	"configcenter/src/common"
	mCommon "configcenter/src/scene_server/admin_server/common"
	"configcenter/src/scene_server/validator"
)

// default group
var (
	groupBaseInfo = mCommon.BaseInfo
)

// Distribution init revision
var Distribution = "community" // could be community or enterprise

/*
	&Attribute{ObjectID: objID, PropertyID: "", PropertyName: "", IsRequired: , IsOnly: , PropertyGroup: , PropertyType: , Option: ""},
*/

// AppRow app structure
func AppRow() []*Attribute {
	objID := common.BKInnerObjIDApp

	groupAppRole := mCommon.AppRole

	lifeCycleOption := []validator.EnumVal{{ID: "1", Name: "测试中", Type: "text"}, {ID: "2", Name: "已上线", Type: "text", IsDefault: true}, {ID: "3", Name: "停运", Type: "text"}}
	languageOption := []validator.EnumVal{{ID: "1", Name: "中文", Type: "text", IsDefault: true}, {ID: "2", Name: "English", Type: "text"}}
	dataRows := []*Attribute{
		&Attribute{ObjectID: objID, PropertyID: "biz_name", PropertyName: "业务名", IsRequired: true, IsOnly: true, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "life_cycle", PropertyName: "生命周期", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: lifeCycleOption},

		//role
		&Attribute{ObjectID: objID, PropertyID: common.BKMaintainersField, PropertyName: "运维人员", IsRequired: true, IsOnly: false, IsEditable: true, PropertyGroup: groupAppRole, PropertyType: common.FieldTypeUser, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: common.BKProductPMField, PropertyName: "产品人员", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupAppRole, PropertyType: common.FieldTypeUser, Option: ""},

		&Attribute{ObjectID: objID, PropertyID: common.BKTesterField, PropertyName: "测试人员", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupAppRole, PropertyType: common.FieldTypeUser, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "biz_developer", PropertyName: "开发人员", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupAppRole, PropertyType: common.FieldTypeUser, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: common.BKOperatorField, PropertyName: "操作人员", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupAppRole, PropertyType: common.FieldTypeUser, Option: ""},

		&Attribute{ObjectID: objID, PropertyID: "time_zone", PropertyName: "时区", IsRequired: true, IsOnly: false, IsEditable: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeTimeZone, Option: "", IsReadOnly: true},
		&Attribute{ObjectID: objID, PropertyID: "language", PropertyName: "语言", IsRequired: true, IsOnly: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: languageOption, IsReadOnly: true},
	}

	return dataRows

}

// SetRow set structure
func SetRow() []*Attribute {
	objID := common.BKInnerObjIDSet
	serviceStatusOption := []validator.EnumVal{{ID: "1", Name: "开放", Type: "text", IsDefault: true}, {ID: "2", Name: "关闭", Type: "text"}}

	dataRows := []*Attribute{
		&Attribute{ObjectID: objID, PropertyID: common.BKAppIDField, PropertyName: "业务ID", IsAPI: true, IsRequired: false, IsOnly: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Option: validator.MinMaxOption{}},
		&Attribute{ObjectID: objID, PropertyID: "set_name", PropertyName: "集群名字", IsRequired: true, IsOnly: true, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "set_desc", PropertyName: "集群描述", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "set_env", PropertyName: "环境类型", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: []validator.EnumVal{{ID: "1", Name: "测试", Type: "text"}, {ID: "2", Name: "体验", Type: "text"}, {ID: "3", Name: "正式", Type: "text", IsDefault: true}}},
		&Attribute{ObjectID: objID, PropertyID: "service_status", PropertyName: "服务状态", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: serviceStatusOption},
		&Attribute{ObjectID: objID, PropertyID: "description", PropertyName: "备注", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "capacity", PropertyName: "设计容量", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Option: validator.MinMaxOption{Min: "1", Max: "999999999"}},

		&Attribute{ObjectID: objID, PropertyID: common.BKChildStr, PropertyName: "", IsRequired: false, IsOnly: false, IsSystem: true, PropertyType: "", Option: ""},
		&Attribute{ObjectID: objID, PropertyID: common.BKInstParentStr, PropertyName: "", IsSystem: true, IsRequired: true, IsOnly: true, PropertyType: common.FieldTypeInt, Option: validator.MinMaxOption{}},
	}
	return dataRows
}

// ModuleRow module structure
func ModuleRow() []*Attribute {
	objID := common.BKInnerObjIDModule
	moduleTypeOption := []validator.EnumVal{{ID: "1", Name: "普通", Type: "text", IsDefault: true}, {ID: "2", Name: "数据库", Type: "text"}}

	dataRows := []*Attribute{
		&Attribute{ObjectID: objID, PropertyID: common.BKAppIDField, PropertyName: "业务ID", IsAPI: true, IsRequired: false, IsOnly: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Option: validator.MinMaxOption{}},
		&Attribute{ObjectID: objID, PropertyID: common.BKSetIDField, PropertyName: "集群ID", IsAPI: true, IsRequired: false, IsOnly: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Option: validator.MinMaxOption{}},
		&Attribute{ObjectID: objID, PropertyID: common.BKModuleNameField, PropertyName: "模块名", IsRequired: true, IsOnly: true, IsEditable: true, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: common.BKChildStr, PropertyName: "", IsRequired: false, IsOnly: false, IsSystem: true, PropertyType: "", Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "module_type", PropertyName: "模块类型", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: moduleTypeOption},
		&Attribute{ObjectID: objID, PropertyID: "operator", PropertyName: "主要维护人", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeUser, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bak_operator", PropertyName: "备份维护人", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeUser, Option: ""},
	}
	return dataRows
}

// PlatRow plat structure
func PlatRow() []*Attribute {
	objID := common.BKInnerObjIDPlat
	dataRows := []*Attribute{
		&Attribute{ObjectID: objID, PropertyID: common.BKCloudNameField, PropertyName: "云区域", IsRequired: true, IsOnly: true, IsPre: true, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: common.BKOwnerIDField, PropertyName: "供应商", IsRequired: true, IsOnly: true, IsPre: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
	}
	return dataRows
}

// HostRow host structure
func HostRow() []*Attribute {
	objID := common.BKInnerObjIDHost
	dataRows := []*Attribute{
		//基本信息分组
		&Attribute{ObjectID: objID, PropertyID: "hostname", PropertyName: "主机名称", IsRequired: true, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "asset_id", PropertyName: "资产编号", IsRequired: false, IsOnly: false, IsEditable: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "sn", PropertyName: "设备SN", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "agentVersion", PropertyName: "agent版本", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "agentStatus", PropertyName: "agent状态", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: []validator.EnumVal{{ID: "1", Name: "未安装", Type: "text"}, {ID: "2", Name: "异常", Type: "text"}, {ID: "3", Name: "正常", Type: "text"}}},
		&Attribute{ObjectID: objID, PropertyID: "ip", PropertyName: "IP", IsRequired: true, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: common.PatternMultipleIP},
		&Attribute{ObjectID: objID, PropertyID: "mac", PropertyName: "物理地址", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "status", PropertyName: "运营状态", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: statusEnum},
		&Attribute{ObjectID: objID, PropertyID: "environment", PropertyName: "主机环境", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: environmentEnum},
		&Attribute{ObjectID: objID, PropertyID: "type", PropertyName: "主机类型", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: []validator.EnumVal{{ID: "1", Name: "物理机", Type: "text"}, {ID: "2", Name: "虚拟机", Type: "text"}}},
		&Attribute{ObjectID: objID, PropertyID: "startU", PropertyName: "起始U位", IsRequired: false, IsOnly: false, IsEditable: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Option: validator.MinMaxOption{Min: "1", Max: "42"}},
		&Attribute{ObjectID: objID, PropertyID: "occupiedU", PropertyName: "占用U数", IsRequired: false, IsOnly: false, IsEditable: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Option: validator.MinMaxOption{Min: "1", Max: "42"}},
		&Attribute{ObjectID: objID, PropertyID: "provider", PropertyName: "供应商", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "comment", PropertyName: "备注", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "cpuSize", PropertyName: "CPU大小", IsRequired: false, IsOnly: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Unit: "核", Option: validator.MinMaxOption{Min: "1", Max: "512"}},
		&Attribute{ObjectID: objID, PropertyID: "memSize", PropertyName: "内存大小", IsRequired: false, IsOnly: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Unit: "GB", Option: validator.MinMaxOption{Min: "1", Max: "100000000"}},
		&Attribute{ObjectID: objID, PropertyID: "diskSize", PropertyName: "磁盘大小", IsRequired: false, IsOnly: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Unit: "GB", Option: validator.MinMaxOption{Min: "1", Max: "100000000"}},
		&Attribute{ObjectID: objID, PropertyID: common.BKOSTypeField, PropertyName: "操作系统类型", IsRequired: false, IsOnly: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: []validator.EnumVal{{ID: "1", Name: "Linux", Type: "text"}, {ID: "2", Name: "Windows", Type: "text"}}},
		&Attribute{ObjectID: objID, PropertyID: "os_version", PropertyName: "操作系统版本", IsRequired: false, IsOnly: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "os_bit", PropertyName: "操作系统位数", IsRequired: false, IsOnly: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: []validator.EnumVal{{ID: "1", Name: "32bit", Type: "text"}, {ID: "2", Name: "64bit", Type: "text"}}},
		&Attribute{ObjectID: objID, PropertyID: "service_term", PropertyName: "质保年限", IsRequired: false, IsOnly: false, IsEditable: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Option: validator.MinMaxOption{Min: "1", Max: "10"}},
		&Attribute{ObjectID: objID, PropertyID: "sla", PropertyName: "SLA级别", IsRequired: false, IsOnly: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: []validator.EnumVal{{ID: "1", Name: "L1", Type: "text"}, {ID: "2", Name: "L2", Type: "text"}, {ID: "3", Name: "L3", Type: "text"}}},
		&Attribute{ObjectID: objID, PropertyID: common.BKCloudIDField, PropertyName: "云区域", IsRequired: false, IsOnly: true, IsEditable: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleAsst, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "state_name", PropertyName: "所在国家", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: stateEnum},
		&Attribute{ObjectID: objID, PropertyID: "province_name", PropertyName: "所在省份", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: provincesEnum},
		&Attribute{ObjectID: objID, PropertyID: "isp_name", PropertyName: "所属运营商", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: ispNameEnum},
		//agent 没有分组
		&Attribute{ObjectID: objID, PropertyID: common.CreateTimeField, PropertyName: "录入时间", IsRequired: false, IsOnly: false, IsEditable: false, PropertyGroup: mCommon.GroupNone, PropertyType: common.FieldTypeTime, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "import_from", PropertyName: "录入方式", IsRequired: false, IsOnly: false, IsEditable: false, PropertyGroup: mCommon.GroupNone, PropertyType: common.FieldTypeEnum, Option: []validator.EnumVal{{ID: "1", Name: "excel", Type: "text"}, {ID: "2", Name: "agent", Type: "text"}, {ID: "3", Name: "api", Type: "text"}}},
	}

	return dataRows
}

// ProcRow proc structure
func ProcRow() []*Attribute {
	objID := common.BKInnerObjIDProc
	groupPort := mCommon.ProcPort
	// groupGsekit := mCommon.Proc_gsekit_base_info
	// groupGsekitManage := mCommon.Proc_gsekit_manage_info
	dataRows := []*Attribute{
		//base info
		//&Attribute{ObjectID: objID, PropertyID: "process_id", PropertyName: "进程ID", IsSystem: true, IsRequired: true, IsOnly: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Option: "{}"},
		&Attribute{ObjectID: objID, PropertyID: common.BKAppIDField, PropertyName: "业务ID", IsAPI: true, IsRequired: true, IsOnly: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Option: validator.MinMaxOption{}},
		&Attribute{ObjectID: objID, PropertyID: common.BKProcessNameField, PropertyName: "进程名称", IsRequired: true, IsOnly: true, IsPre: true, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "description", PropertyName: "进程描述", IsRequired: false, IsOnly: false, IsPre: true, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},

		//监听端口分组
		&Attribute{ObjectID: objID, PropertyID: "bind_ip", PropertyName: "绑定IP", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupPort, PropertyType: common.FieldTypeEnum, Option: []validator.EnumVal{{ID: "1", Name: "127.0.0.1", Type: "text"}, {ID: "2", Name: "0.0.0.0", Type: "text"}, {ID: "3", Name: "第一内网IP", Type: "text"}, {ID: "4", Name: "第一外网IP", Type: "text"}}},
		&Attribute{ObjectID: objID, PropertyID: "port", PropertyName: "端口", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupPort, PropertyType: common.FieldTypeSingleChar, Option: common.PatternMultiplePortRange, Placeholder: `单个端口：8080 </br>多个连续端口：8080-8089 </br>多个不连续端口：8080-8089,8199`},
		&Attribute{ObjectID: objID, PropertyID: "protocol", PropertyName: "协议", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupPort, PropertyType: common.FieldTypeEnum, Option: []validator.EnumVal{{ID: "1", Name: "TCP", Type: "text"}, {ID: "2", Name: "UDP", Type: "text"}}},

		//gsekit 基础信息
		&Attribute{ObjectID: objID, PropertyID: "func_id", PropertyName: "功能ID", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: mCommon.GroupNone, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "func_name", PropertyName: "功能名称", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: mCommon.GroupNone, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "work_path", PropertyName: "工作路径", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: mCommon.GroupNone, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "user", PropertyName: "启动用户", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: mCommon.GroupNone, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "proc_num", PropertyName: "启动数量", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: mCommon.GroupNone, PropertyType: common.FieldTypeInt, Option: validator.MinMaxOption{Min: "1", Max: "1000000"}},
		&Attribute{ObjectID: objID, PropertyID: "priority", PropertyName: "启动优先级", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: mCommon.GroupNone, PropertyType: common.FieldTypeInt, Option: validator.MinMaxOption{Min: "1", Max: "100"}},
		&Attribute{ObjectID: objID, PropertyID: "timeout", PropertyName: "操作超时时长", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: mCommon.GroupNone, PropertyType: common.FieldTypeInt, Option: validator.MinMaxOption{Min: "1", Max: "1000000"}},

		//gsekit 进程信息
		&Attribute{ObjectID: objID, PropertyID: "start_cmd", PropertyName: "启动命令", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: mCommon.GroupNone, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "stop_cmd", PropertyName: "停止命令", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: mCommon.GroupNone, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "restart_cmd", PropertyName: "重启命令", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: mCommon.GroupNone, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "face_stop_cmd", PropertyName: "强制停止命令", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: mCommon.GroupNone, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "reload_cmd", PropertyName: "进程重载命令", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: mCommon.GroupNone, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "pid_file", PropertyName: "PID文件路径", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: mCommon.GroupNone, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "auto_start", PropertyName: "是否自动拉起", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: mCommon.GroupNone, PropertyType: common.FieldTypeBool, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "auto_time_gap", PropertyName: "拉起间隔", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: mCommon.GroupNone, PropertyType: common.FieldTypeInt, Option: validator.MinMaxOption{Min: "1", Max: "1000000"}},
	}
	return dataRows
}

// SwitchRow proc structure
func SwitchRow() []*Attribute {
	objID := common.BKInnerObjIDSwitch
	dataRows := []*Attribute{
		&Attribute{ObjectID: objID, PropertyID: "inst_name", PropertyName: "设备名称", IsRequired: true, IsOnly: true, IsPre: true, IsEditable: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "asset_id", PropertyName: "资产编号", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "status", PropertyName: "状态", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: []validator.EnumVal{{ID: "1", Name: "使用中", Type: "text"}, {ID: "2", Name: "未使用", Type: "text"}}},
		&Attribute{ObjectID: objID, PropertyID: "sn", PropertyName: "设备SN", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "func", PropertyName: "设备用途", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "vendor", PropertyName: "设备厂商", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "model", PropertyName: "设备型号", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "admin_ip", PropertyName: "管理IP", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: common.PatternMultipleIP},
		&Attribute{ObjectID: objID, PropertyID: "operator", PropertyName: "维护人", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "os_detail", PropertyName: "操作系统详情", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "comment", PropertyName: "备注", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeLongChar, Option: ""},
	}
	return dataRows
}

// RouterRow proc structure
func RouterRow() []*Attribute {
	objID := common.BKInnerObjIDRouter
	dataRows := []*Attribute{
		&Attribute{ObjectID: objID, PropertyID: "inst_name", PropertyName: "设备名称", IsRequired: true, IsOnly: true, IsPre: true, IsEditable: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "admin_ip", PropertyName: "管理IP", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: common.PatternMultipleIP},
		&Attribute{ObjectID: objID, PropertyID: "status", PropertyName: "状态", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: []validator.EnumVal{{ID: "1", Name: "使用中", Type: "text"}, {ID: "2", Name: "未使用", Type: "text"}}},
		&Attribute{ObjectID: objID, PropertyID: "asset_id", PropertyName: "资产编号", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "startU", PropertyName: "起始U位", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "occupiedU", PropertyName: "占用U数", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "sn", PropertyName: "设备SN", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "func", PropertyName: "设备用途", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "vendor", PropertyName: "设备厂商", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "model", PropertyName: "设备型号", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "operator", PropertyName: "维护人", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "os_detail", PropertyName: "操作系统详情", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "comment", PropertyName: "备注", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeLongChar, Option: ""},
	}
	return dataRows
}

// LoadBalanceRow proc structure
func LoadBalanceRow() []*Attribute {
	objID := common.BKInnerObjIDBlance
	dataRows := []*Attribute{
		&Attribute{ObjectID: objID, PropertyID: "inst_name", PropertyName: "设备名称", IsRequired: true, IsOnly: true, IsPre: true, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "asset_id", PropertyName: "资产编号", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "status", PropertyName: "状态", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: []validator.EnumVal{{ID: "1", Name: "使用中", Type: "text"}, {ID: "2", Name: "未使用", Type: "text"}}},
		&Attribute{ObjectID: objID, PropertyID: "sn", PropertyName: "设备SN", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "func", PropertyName: "设备用途", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "vendor", PropertyName: "设备厂商", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "model", PropertyName: "设备型号", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "admin_ip", PropertyName: "管理IP", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: common.PatternMultipleIP},
		&Attribute{ObjectID: objID, PropertyID: "operator", PropertyName: "维护人", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "os_detail", PropertyName: "操作系统详情", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "comment", PropertyName: "备注", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeLongChar, Option: ""},
	}
	return dataRows
}

// FirewallRow proc structure
func FirewallRow() []*Attribute {
	objID := common.BKInnerObjIDFirewall
	dataRows := []*Attribute{
		&Attribute{ObjectID: objID, PropertyID: "inst_name", PropertyName: "设备名称", IsRequired: true, IsOnly: true, IsPre: true, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "asset_id", PropertyName: "固资编号", IsRequired: false, IsOnly: false, IsPre: true, IsEditable: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "status", PropertyName: "状态", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: []validator.EnumVal{{ID: "1", Name: "使用中", Type: "text"}, {ID: "2", Name: "未使用", Type: "text"}}},
		&Attribute{ObjectID: objID, PropertyID: "sn", PropertyName: "设备SN", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "func", PropertyName: "设备用途", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "vendor", PropertyName: "设备厂商", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "model", PropertyName: "设备型号", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "admin_ip", PropertyName: "管理IP", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: common.PatternMultipleIP},
		&Attribute{ObjectID: objID, PropertyID: "operator", PropertyName: "维护人", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "os_detail", PropertyName: "操作系统详情", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "comment", PropertyName: "备注", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeLongChar, Option: ""},
	}
	return dataRows
}

// IDC proc structure
func IdcRow() []*Attribute {
	objID := common.BKInnerObjIDIdc
	dataRows := []*Attribute{
		&Attribute{ObjectID: objID, PropertyID: "inst_name", PropertyName: "名称", IsRequired: true, IsOnly: true, IsPre: true, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "shortname", PropertyName: "简称", IsRequired: false, IsOnly: false, IsPre: true, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "telphone", PropertyName: "值班电话", IsRequired: false, IsOnly: false, IsPre: true, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "address", PropertyName: "地址", IsRequired: false, IsOnly: false, IsPre: true, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "comment", PropertyName: "备注", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "status", PropertyName: "状态", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: []validator.EnumVal{{ID: "1", Name: "使用中", Type: "text"}, {ID: "2", Name: "未使用", Type: "text"}}},
	}
	return dataRows
}

// IDCRack proc structure
func IDCRackRow() []*Attribute {
	objID := common.BKInnerObjIDIdcRack
	dataRows := []*Attribute{
		&Attribute{ObjectID: objID, PropertyID: "inst_name", PropertyName: "名称", IsRequired: true, IsOnly: true, IsPre: true, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "status", PropertyName: "状态", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: []validator.EnumVal{{ID: "1", Name: "使用中", Type: "text"}, {ID: "2", Name: "未使用", Type: "text"}}},
		&Attribute{ObjectID: objID, PropertyID: "unum", PropertyName: "机柜U数", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Option: common.PatternInt},
		&Attribute{ObjectID: objID, PropertyID: "comment", PropertyName: "备注", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
	}
	return dataRows
}

// WeblogicRow proc structure
func WeblogicRow() []*Attribute {
	objID := common.BKInnerObjIDWeblogic
	dataRows := []*Attribute{
		&Attribute{ObjectID: objID, PropertyID: "inst_key", PropertyName: "中间件标识", IsRequired: true, IsOnly: true, IsPre: true, IsEditable: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "inst_name", PropertyName: "名称", IsRequired: true, IsOnly: false, IsPre: true, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "version", PropertyName: "版本", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "patch_version", PropertyName: "补丁版本", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "main_path", PropertyName: "主目录", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "log_path", PropertyName: "日志路径", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "vendor", PropertyName: "厂商", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "ip", PropertyName: "IP地址", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: common.PatternMultipleIP},
		&Attribute{ObjectID: objID, PropertyID: "port", PropertyName: "端口", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "detail", PropertyName: "详细描述", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "jdk_version", PropertyName: "JDK版本", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "jvm_free_mem", PropertyName: "JVM配置的最大空闲内存", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "jvm_capacity", PropertyName: "JVM堆的当前大小", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "jvm_used_mem", PropertyName: "JVM堆的当前可用的内存", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
	}
	return dataRows
}

// TomcatRow proc structure
func TomcatRow() []*Attribute {
	objID := common.BKInnerObjIDTomcat
	dataRows := []*Attribute{
		&Attribute{ObjectID: objID, PropertyID: "inst_key", PropertyName: "中间件标识", IsRequired: true, IsOnly: true, IsPre: true, IsEditable: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "inst_name", PropertyName: "名称", IsRequired: true, IsOnly: false, IsPre: true, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "version", PropertyName: "版本", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "patch_version", PropertyName: "补丁版本", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "main_path", PropertyName: "主目录", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "log_path", PropertyName: "日志路径", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "vendor", PropertyName: "厂商", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "ip", PropertyName: "IP地址", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: common.PatternMultipleIP},
		&Attribute{ObjectID: objID, PropertyID: "port", PropertyName: "端口", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "detail", PropertyName: "详细描述", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "jdk_version", PropertyName: "JDK版本", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeLongChar, Option: ""},
	}
	return dataRows
}

// ApacheRow proc structure
func ApacheRow() []*Attribute {
	objID := common.BKInnerObjIDApache
	dataRows := []*Attribute{
		&Attribute{ObjectID: objID, PropertyID: "inst_key", PropertyName: "中间件标识", IsRequired: true, IsOnly: true, IsPre: true, IsEditable: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "inst_name", PropertyName: "名称", IsRequired: true, IsOnly: false, IsPre: true, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "version", PropertyName: "版本", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "patch_version", PropertyName: "补丁版本", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "main_path", PropertyName: "主目录", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "log_path", PropertyName: "日志路径", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "vendor", PropertyName: "厂商", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "ip", PropertyName: "IP地址", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: common.PatternMultipleIP},
		&Attribute{ObjectID: objID, PropertyID: "port", PropertyName: "端口", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "detail", PropertyName: "详细描述", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "max_connect", PropertyName: "最大连接请求数", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Option: validator.MinMaxOption{}},
		&Attribute{ObjectID: objID, PropertyID: "max_keepalive", PropertyName: "最大keepAlive请求数", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Option: validator.MinMaxOption{}},
	}
	return dataRows
}

var stateEnum = []validator.EnumVal{
	{ID: "CN", Name: "中国", Type: "text"},
	{ID: "JP", Name: "日本", Type: "text"},
	{ID: "US", Name: "美国", Type: "text"},
}

var provincesEnum = []validator.EnumVal{
	{ID: "110000", Name: "北京市", Type: "text"},
	{ID: "120000", Name: "天津市", Type: "text"},
	{ID: "130000", Name: "河北省", Type: "text"},
	{ID: "140000", Name: "山西省", Type: "text"},
	{ID: "150000", Name: "内蒙古自治区", Type: "text"},
	{ID: "210000", Name: "辽宁省", Type: "text"},
	{ID: "220000", Name: "吉林省", Type: "text"},
	{ID: "230000", Name: "黑龙江省", Type: "text"},
	{ID: "310000", Name: "上海市", Type: "text"},
	{ID: "320000", Name: "江苏省", Type: "text"},
	{ID: "330000", Name: "浙江省", Type: "text"},
	{ID: "340000", Name: "安徽省", Type: "text"},
	{ID: "350000", Name: "福建省", Type: "text"},
	{ID: "360000", Name: "江西省", Type: "text"},
	{ID: "370000", Name: "山东省", Type: "text"},
	{ID: "410000", Name: "河南省", Type: "text"},
	{ID: "420000", Name: "湖北省", Type: "text"},
	{ID: "430000", Name: "湖南省", Type: "text"},
	{ID: "440000", Name: "广东省", Type: "text"},
	{ID: "450000", Name: "广西壮族自治区", Type: "text"},
	{ID: "460000", Name: "海南省", Type: "text"},
	{ID: "500000", Name: "重庆市", Type: "text"},
	{ID: "510000", Name: "四川省", Type: "text"},
	{ID: "520000", Name: "贵州省", Type: "text"},
	{ID: "530000", Name: "云南省", Type: "text"},
	{ID: "540000", Name: "西藏自治区", Type: "text"},
	{ID: "610000", Name: "陕西省", Type: "text"},
	{ID: "620000", Name: "甘肃省", Type: "text"},
	{ID: "630000", Name: "青海省", Type: "text"},
	{ID: "640000", Name: "宁夏回族自治区", Type: "text"},
	{ID: "650000", Name: "新疆维吾尔自治区", Type: "text"},
	{ID: "710000", Name: "台湾省", Type: "text"},
	{ID: "810000", Name: "香港特别行政区", Type: "text"},
	{ID: "820000", Name: "澳门特别行政区", Type: "text"},
}

var ispNameEnum = []validator.EnumVal{
	{ID: "0", Name: "其他", Type: "text"},
	{ID: "1", Name: "电信", Type: "text"},
	{ID: "2", Name: "联通", Type: "text"},
	{ID: "3", Name: "移动", Type: "text"},
}

var statusEnum = []validator.EnumVal{
	{ID: "1", Name: "运营中", Type: "text"},
	{ID: "2", Name: "故障中", Type: "text"},
	{ID: "3", Name: "未上线", Type: "text"},
	{ID: "4", Name: "下线隔离中", Type: "text"},
	{ID: "5", Name: "开发机", Type: "text"},
	{ID: "6", Name: "测试机", Type: "text"},
	{ID: "7", Name: "维修中", Type: "text"},
	{ID: "8", Name: "报废", Type: "text"},
}

var environmentEnum = []validator.EnumVal{
	{ID: "1", Name: "无", Type: "text"},
	{ID: "2", Name: "开发", Type: "text"},
	{ID: "3", Name: "测试", Type: "text"},
	{ID: "4", Name: "预发布", Type: "text"},
	{ID: "5", Name: "生产", Type: "text"},
	{ID: "6", Name: "灾备", Type: "text"},
}

type Attribute struct {
	ID                int64       `field:"id" json:"id" bson:"id"`
	OwnerID           string      `field:"org_id" json:"org_id" bson:"org_id"`
	ObjectID          string      `field:"obj_id" json:"obj_id" bson:"obj_id"`
	PropertyID        string      `field:"property_id" json:"property_id" bson:"property_id"`
	PropertyName      string      `field:"property_name" json:"property_name" bson:"property_name"`
	PropertyGroup     string      `field:"property_group" json:"property_group" bson:"property_group"`
	PropertyGroupName string      `field:"property_group_name,ignoretomap" json:"property_group_name" bson:"-"`
	PropertyIndex     int64       `field:"property_index" json:"property_index" bson:"property_index"`
	Unit              string      `field:"unit" json:"unit" bson:"unit"`
	Placeholder       string      `field:"placeholder" json:"placeholder" bson:"placeholder"`
	IsEditable        bool        `field:"editable" json:"editable" bson:"editable"`
	IsPre             bool        `field:"ispre" json:"ispre" bson:"ispre"`
	IsRequired        bool        `field:"isrequired" json:"isrequired" bson:"isrequired"`
	IsReadOnly        bool        `field:"isreadonly" json:"isreadonly" bson:"isreadonly"`
	IsOnly            bool        `field:"isonly" json:"isonly" bson:"isonly"`
	IsSystem          bool        `field:"issystem" json:"issystem" bson:"issystem"`
	IsAPI             bool        `field:"isapi" json:"isapi" bson:"isapi"`
	PropertyType      string      `field:"property_type" json:"property_type" bson:"property_type"`
	Option            interface{} `field:"option" json:"option" bson:"option"`
	Description       string      `field:"description" json:"description" bson:"description"`
	Creator           string      `field:"creator" json:"creator" bson:"creator"`
	CreateTime        *time.Time  `json:"create_time" bson:"create_time"`
	LastTime          *time.Time  `json:"last_time" bson:"last_time"`
}
