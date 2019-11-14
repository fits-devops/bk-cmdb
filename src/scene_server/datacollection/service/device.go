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

package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os/exec"
	"strconv"

	restful "github.com/emicklei/go-restful"

	"configcenter/src/common"
	"configcenter/src/common/blog"
	"configcenter/src/common/errors"
	meta "configcenter/src/common/metadata"
	"configcenter/src/common/util"
)

// CreateDevice create device
func (s *Service) CreateDevice(req *restful.Request, resp *restful.Response) {
	pheader := req.Request.Header
	defErr := s.CCErr.CreateDefaultCCErrorIf(util.GetLanguage(pheader))

	deviceInfo := meta.NetcollectDevice{}
	if err := json.NewDecoder(req.Request.Body).Decode(&deviceInfo); nil != err {
		blog.Errorf("[NetDevice] add device failed with decode body err: %v", err)
		resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: defErr.Error(common.CCErrCommJSONUnmarshalFailed)})
		return
	}

	result, err := s.Logics.AddDevice(pheader, deviceInfo)
	if nil != err {
		if err.Error() == defErr.Error(common.CCErrCollectNetDeviceCreateFail).Error() {
			resp.WriteError(http.StatusInternalServerError, &meta.RespError{Msg: err})
			return
		}

		resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: err})
		return
	}

	resp.WriteEntity(meta.NewSuccessResp(result))
}

// UpdateDevice update device
func (s *Service) UpdateDevice(req *restful.Request, resp *restful.Response) {
	pheader := req.Request.Header
	defErr := s.CCErr.CreateDefaultCCErrorIf(util.GetLanguage(pheader))

	netDeviceID, err := checkDeviceIDPathParam(defErr, req.PathParameter("device_id"))
	if nil != err {
		resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: err})
		return
	}

	deviceInfo := meta.NetcollectDevice{}
	if err := json.NewDecoder(req.Request.Body).Decode(&deviceInfo); nil != err {
		blog.Errorf("[NetDevice] update device failed with decode body err: %v", err)
		resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: defErr.Error(common.CCErrCommJSONUnmarshalFailed)})
		return
	}

	if err = s.Logics.UpdateDevice(pheader, netDeviceID, deviceInfo); nil != err {
		if err.Error() == defErr.Error(common.CCErrCollectNetDeviceUpdateFail).Error() {
			resp.WriteError(http.StatusInternalServerError, &meta.RespError{Msg: err})
			return
		}

		resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: err})
		return
	}

	resp.WriteEntity(meta.NewSuccessResp(nil))
}

// BatchCreateDevice batch create device
func (s *Service) BatchCreateDevice(req *restful.Request, resp *restful.Response) {
	pheader := req.Request.Header
	defErr := s.CCErr.CreateDefaultCCErrorIf(util.GetLanguage(pheader))

	batchAddDevice := new(meta.BatchAddDevice)
	if err := json.NewDecoder(req.Request.Body).Decode(&batchAddDevice); nil != err {
		blog.Errorf("[NetDevice] batch add device failed with decode body err: %v", err)
		resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: defErr.Error(common.CCErrCommJSONUnmarshalFailed)})
		return
	}

	deviceList := batchAddDevice.Data
	resultList, hasError := s.Logics.BatchCreateDevice(pheader, deviceList)
	if hasError {
		resp.WriteEntity(meta.Response{
			BaseResp: meta.BaseResp{
				Result: false,
				Code:   common.CCErrCollectNetDeviceCreateFail,
				ErrMsg: defErr.Error(common.CCErrCollectNetDeviceCreateFail).Error()},
			Data: resultList,
		})
		return
	}

	resp.WriteEntity(meta.NewSuccessResp(resultList))
}

// SearchDevice search device
func (s *Service) SearchDevice(req *restful.Request, resp *restful.Response) {
	pheader := req.Request.Header
	defErr := s.CCErr.CreateDefaultCCErrorIf(util.GetLanguage(pheader))

	body := new(meta.NetCollSearchParams)
	if err := json.NewDecoder(req.Request.Body).Decode(body); nil != err {
		blog.Errorf("[NetDevice] search net device failed with decode body err: %v", err)
		resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: defErr.Error(common.CCErrCommJSONUnmarshalFailed)})
		return
	}

	devices, err := s.Logics.SearchDevice(pheader, body)
	if nil != err {
		blog.Errorf("[NetDevice] search net device failed, err: %v", err)
		resp.WriteError(http.StatusInternalServerError, &meta.RespError{Msg: defErr.Error(common.CCErrCollectNetDeviceGetFail)})
		return
	}

	resp.WriteEntity(meta.SearchNetDeviceResult{
		BaseResp: meta.SuccessBaseResp,
		Data:     devices,
	})
}

// DeleteDevice delete device
func (s *Service) DeleteDevice(req *restful.Request, resp *restful.Response) {
	pheader := req.Request.Header
	defErr := s.CCErr.CreateDefaultCCErrorIf(util.GetLanguage(pheader))

	deleteNetDeviceBatchOpt := new(meta.DeleteNetDeviceBatchOpt)
	if err := json.NewDecoder(req.Request.Body).Decode(deleteNetDeviceBatchOpt); nil != err {
		blog.Errorf("[NetDevice] delete net device batch, but decode body failed, err: %v", err)
		resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: defErr.Error(common.CCErrCommJSONUnmarshalFailed)})
		return
	}

	for _, deviceID := range deleteNetDeviceBatchOpt.DeviceIDs {
		if err := s.Logics.DeleteDevice(pheader, deviceID); nil != err {
			blog.Errorf("[NetDevice] delete net device failed, with device_id [%d], err: %v", deviceID, err)

			if defErr.Error(common.CCErrCollectNetDeviceHasPropertyDeleteFail).Error() == err.Error() {
				resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: err})
				return
			}

			resp.WriteError(http.StatusInternalServerError, &meta.RespError{Msg: defErr.Error(common.CCErrCollectNetDeviceDeleteFail)})
			return
		}
	}

	resp.WriteEntity(meta.NewSuccessResp(nil))
}

func checkDeviceIDPathParam(defErr errors.DefaultCCErrorIf, ID string) (uint64, error) {
	netDeviceID, err := strconv.ParseUint(ID, 10, 64)
	if nil != err {
		blog.Errorf("[NetDevice] update net device with id[%s] to parse the net device id, error: %v", ID, err)
		return 0, defErr.Errorf(common.CCErrCommParamsNeedInt, common.BKDeviceIDField)
	}
	if 0 == netDeviceID {
		blog.Errorf("[NetDevice] update net device with id[%d] should not be 0", netDeviceID)
		return 0, defErr.Error(common.CCErrCommHTTPInputInvalid)
	}

	return netDeviceID, nil
}

// CreateHost create host
func (s *Service) CreateHost(req *restful.Request, resp *restful.Response) {
	pheader := req.Request.Header
	defErr := s.CCErr.CreateDefaultCCErrorIf(util.GetLanguage(pheader))
	body := new(meta.HostParams)
	if err := json.NewDecoder(req.Request.Body).Decode(&body); err != nil {
		blog.Errorf("[CreateHost] decode body failed, err: %v", err)
		resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: defErr.Error(common.CCErrCommJSONUnmarshalFailed)})
		return
	}
	port := body.SshPort
	userName := body.HostUserName
	pwd := body.HostUserPassword
	for _,hostIp := range body.TargetHostIp {
		command := "./test.sh "+hostIp+" "+port+" "+ userName+" "+pwd//脚本的路径
		cmd := exec.Command("/bin/bash", "-c",command)
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			blog.Errorf("[CreateHost] failed, err: %v", stderr.String())
		}
		ok := out.String()
		if ok == "100" {
			blog.Errorf("[CreateHost] success, err: %v", ok)
			result := meta.AddHostResult{MinionId: hostIp}
			resp.WriteEntity(meta.NewSuccessResp(result))
		}else{
			blog.Info("exec sh result: %v", ok)
			resp.WriteError(http.StatusInternalServerError, &meta.RespError{Msg: defErr.Error(common.CCErrCollectCreateHostFail)})
		}

	}
	// 檢查ip 是否存在 或者 sh 脚本檢查 存在手动添加的主机 可能是agent 未安装 minionId 不存在
	// 可以查询 minionId 是否存在 再去调用脚本

	//{"inst_name":"haowan66607","ip":"192.168.31.102","asset_id":"423424"} // 添加主機

}