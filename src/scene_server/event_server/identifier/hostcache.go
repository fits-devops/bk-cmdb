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

package identifier

import (
	"context"
	"encoding/json"
	"sort"

	redis "gopkg.in/redis.v5"

	"configcenter/src/common"
	"configcenter/src/common/blog"
	"configcenter/src/common/util"
	"configcenter/src/storage/dal"
)

type HostIdentifier struct {
	HostID          int64              `json:"host_id" bson:"host_id"`
	HostName        string             `json:"host_name" bson:"host_name"`
	SupplierID      int64              `json:"supplier_id"`
	SupplierAccount string             `json:"org_id"`
	CloudID         int64              `json:"cloud_id" bson:"cloud_id"`
	CloudName       string             `json:"cloud_name" bson:"cloud_name"`
	InnerIP         string             `json:"host_innerip" bson:"host_innerip"`
	OuterIP         string             `json:"host_outerip" bson:"host_outerip"`
	OSType          string             `json:"os_type" bson:"os_type"`
	OSName          string             `json:"os_name" bson:"os_name"`
	Memory          int64              `json:"mem" bson:"mem"`
	CPU             int64              `json:"cpu" bson:"cpu"`
	Disk            int64              `json:"disk" bson:"disk"`
	Module          map[string]*Module `json:"associations" bson:"associations"`
	Process         []Process          `json:"process" bson:"process"`
}

type PorcessSorter []Process

func (p PorcessSorter) Len() int      { return len(p) }
func (p PorcessSorter) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PorcessSorter) Less(i, j int) bool {
	sort.Sort(util.Int64Slice(p[i].BindModules))
	return p[i].ProcessID < p[j].ProcessID
}

type Process struct {
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

type Module struct {
	BizID      int64  `json:"biz_id"`
	BizName    string `json:"biz_name"`
	SetID      int64  `json:"set_id"`
	SetName    string `json:"set_name"`
	ModuleID   int64  `json:"module_id"`
	ModuleName string `json:"module_name"`
	SetStatus  string `json:"service_status"`
	SetEnv     string `json:"set_env"`
}

func (iden *HostIdentifier) MarshalBinary() (data []byte, err error) {
	sort.Sort(PorcessSorter(iden.Process))
	return json.Marshal(iden)
}

func (iden *HostIdentifier) fillIden(ctx context.Context, cache *redis.Client, db dal.RDB) *HostIdentifier {
	// fill cloudName
	cloud, err := getCache(ctx, cache, db, common.BKInnerObjIDPlat, iden.CloudID, false)
	if err != nil {
		blog.Errorf("identifier: getCache error %s", err.Error())
		return iden
	}
	iden.CloudName = getString(cloud.data[common.BKCloudNameField])

	// fill module
	for moduleID := range iden.Module {
		biz, err := getCache(ctx, cache, db, common.BKInnerObjIDApp, iden.Module[moduleID].BizID, false)
		if err != nil {
			blog.Errorf("identifier: getCache error %s", err.Error())
			continue
		}
		iden.Module[moduleID].BizName = getString(biz.data[common.BKAppNameField])
		iden.SupplierAccount = getString(biz.data[common.BKOwnerIDField])
		iden.SupplierID = getInt(biz.data, common.BKSupplierIDField)

		set, err := getCache(ctx, cache, db, common.BKInnerObjIDSet, iden.Module[moduleID].SetID, false)
		if err != nil {
			blog.Errorf("identifier: getCache error %s", err.Error())
			continue
		}
		iden.Module[moduleID].SetName = getString(set.data[common.BKSetNameField])
		iden.Module[moduleID].SetEnv = getString(set.data[common.BKSetEnvField])
		iden.Module[moduleID].SetStatus = getString(set.data[common.BKSetStatusField])

		module, err := getCache(ctx, cache, db, common.BKInnerObjIDModule, iden.Module[moduleID].ModuleID, false)
		if err != nil {
			blog.Errorf("identifier: getCache error %s", err.Error())
			continue
		}
		iden.Module[moduleID].ModuleName = getString(module.data[common.BKModuleNameField])
	}

	// fill process
	for procindex := range iden.Process {
		process := &iden.Process[procindex]
		proc, err := getCache(ctx, cache, db, common.BKInnerObjIDProc, process.ProcessID, false)
		if err != nil {
			blog.Errorf("identifier: getCache for %s %d error %s", common.BKInnerObjIDProc, process.ProcessID, err.Error())
			continue
		}
		process.ProcessName = getString(proc.data[common.BKProcessNameField])
		process.FuncID = getString(proc.data[common.BKFuncIDField])
		process.FuncName = getString(proc.data[common.BKFuncName])
		process.BindIP = getString(proc.data[common.BKBindIP])
		process.PROTOCOL = getString(proc.data[common.BKProtocol])
		process.PORT = getString(proc.data[common.BKPort])
		process.StartParamRegex = getString(proc.data["start_param_regex"])
	}

	return iden
}
