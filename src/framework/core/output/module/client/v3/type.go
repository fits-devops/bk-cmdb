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

package v3

const (
	// Plat plat name
	Plat = "plat"

	// IsIncrement is increment
	IsIncrement = "is_increment"

	// HostID host id
	HostID = "host_id"

	// PlatID  plat id
	PlatID = "cloud_id"
	// BusinessID the business id
	BusinessID = "biz_id"
	// SetID the set id
	SetID = "set_id"
	// ModuleID the module id
	ModuleID = "module_id"
	// ObjectID the object identifier name
	ObjectID = "obj_id"
	// CommonInstID the common inst id
	CommonInstID = "inst_id"
	// SupplierAccount the business id
	SupplierAccount = "org_id"
)

const (
	HostInfoField = "host_info"
)

// v3Resp v3 api response data struct
type v3Resp struct {
	Result  bool        `json:"result"`
	Code    int         `json:"error_code"`
	Message string      `json:"error_msg"`
	Data    interface{} `json:"data"`
}
