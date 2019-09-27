/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package api

// plat
const (
	fieldObjectID = "obj_id"
	plat          = "plat"
)

// set fields
const (
	fieldParentID        = "bk_parent_id"
	fieldSetID           = "set_id"
	fieldSetName         = "set_name"
	fieldPlatID          = "cloud_id"
	fieldPlatName        = "cloud_name"
	fieldSupplierAccount = "org_id"
	fieldSupplierID      = "supplier_id"
	fieldBusinessID      = "biz_id"
	fieldCapacity        = "capacity"
	fieldServiceStatus   = "service_status"
	fieldSetDesc         = "set_desc"
	fieldSetEnv          = "set_env"
	fieldObjID           = "obj_id"
	fieldDescription     = "description"
)

// module fields
const (
	fieldModuleID    = "module_id"
	fieldModuleName  = "module_name"
	fieldBakOperator = "bak_operator"
	fieldModuleTYpe  = "module_type"
	fieldOperator    = "operator"
)

// business fields
const (
	fieldBizDeveloper  = "biz_developer"
	fieldBizID         = "biz_id"
	fieldBizMaintainer = "biz_maintainer"
	fieldBizName       = "biz_name"
	fieldBizProductor  = "biz_productor"
	fieldBizTester     = "biz_tester"
	fieldLifeCycle     = "life_cycle"
	fieldBizOperator   = "operator"
)

// host fields
const (
	fieldOsBit        = "os_bit"
	fieldSLA          = "sla"
	fieldCloudID      = "cloud_id"
	fieldHostInnerIP  = "host_innerip"
	fieldCPU          = "cpu"
	fieldCPUMhz       = "cpu_mhz"
	fieldOsType       = "os_type"
	fieldDisk         = "disk"
	fieldHostID       = "host_id"
	fieldHostOuterIP  = "host_outerip"
	fieldAssetID      = "asset_id"
	fieldMac          = "mac"
	fieldProvinceName = "bk_provinceName"
	fieldSN           = "sn"
	fieldCPUModule    = "cpu_module"
	fieldHostName     = "host_name"
	fieldISPName      = "isp_name"
	fieldOuterMac     = "outer_mac"
	fieldServiceTerm  = "service_term"
	fieldComment      = "comment"
	fieldMem          = "mem"
	fieldOsName       = "os_name"
	fieldOsVersion    = "os_version"
	fieldImportFrom   = "import_from"
	fieldHostOperator = "operator"
)

// Enum definition
const (
	HostSLALevel1            = "1"
	HostSLALevel2            = "2"
	HostSLALevel3            = "3"
	HostOSTypeLinux          = "1"
	HostOSTypeWindows        = "2"
	HostImportFromExcel      = "1"
	HostImportFromAgent      = "2"
	HostImportFromAPI        = "3"
	BusinessLifeCycleTesting = "1"
	BusinessLifeCycleOnLine  = "2"
	BusinessLifeCycleStopped = "3"
	SetEnvTesting            = "1"
	SetEnvGuest              = "2"
	SetEnvNormal             = "3"
	SetServiceOpen           = "1"
	SetServiceClose          = "2"
)

type HostModuleActionType string

const (
	HostAppendModule  HostModuleActionType = "append"
	HostReplaceModule HostModuleActionType = "replace"
)
