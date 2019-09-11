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
	"context"
	"net/http"
	"time"

	"configcenter/src/auth/extensions"
	"configcenter/src/common"
	"configcenter/src/common/backbone"
	cfnc "configcenter/src/common/backbone/configcenter"
	"configcenter/src/common/errors"
	"configcenter/src/common/language"
	"configcenter/src/common/metadata"
	"configcenter/src/common/metric"
	"configcenter/src/common/rdapi"
	"configcenter/src/common/types"
	"configcenter/src/common/util"
	"configcenter/src/scene_server/proc_server/app/options"
	"configcenter/src/scene_server/proc_server/logics"
	ccRedis "configcenter/src/storage/dal/redis"
	"configcenter/src/thirdpartyclient/esbserver"
	"configcenter/src/thirdpartyclient/esbserver/esbutil"

	"github.com/emicklei/go-restful"
	redis "gopkg.in/redis.v5"
)

type srvComm struct {
	header        http.Header
	rid           string
	ccErr         errors.DefaultCCErrorIf
	ccLang        language.DefaultCCLanguageIf
	ctx           context.Context
	ctxCancelFunc context.CancelFunc
	user          string
	ownerID       string
	lgc           *logics.Logics
}

type ProcServer struct {
	*backbone.Engine
	EsbConfigChn       chan esbutil.EsbConfig
	Config             *options.Config
	EsbServ            esbserver.EsbClientInterface
	Cache              *redis.Client
	procHostInstConfig logics.ProcHostInstConfig
	ConfigMap          map[string]string
	AuthManager        *extensions.AuthManager
}

func (s *ProcServer) newSrvComm(header http.Header) *srvComm {
	rid := util.GetHTTPCCRequestID(header)
	lang := util.GetLanguage(header)
	user := util.GetUser(header)
	ctx, cancel := s.Engine.CCCtx.WithCancel()
	ctx = context.WithValue(ctx, common.ContextRequestIDField, rid)
	ctx = context.WithValue(ctx, common.ContextRequestUserField, user)

	return &srvComm{
		header:        header,
		rid:           rid,
		ccErr:         s.CCErr.CreateDefaultCCErrorIf(lang),
		ccLang:        s.Language.CreateDefaultCCLanguageIf(lang),
		ctx:           ctx,
		ctxCancelFunc: cancel,
		user:          util.GetUser(header),
		ownerID:       util.GetOwnerID(header),
		lgc:           logics.NewLogics(s.Engine, header, s.Cache, s.EsbServ, &s.procHostInstConfig),
	}
}

func (ps *ProcServer) WebService() *restful.Container {

	container := restful.NewContainer()

	getErrFunc := func() errors.CCErrorIf {
		return ps.Engine.CCErr
	}

	// v3
	api := new(restful.WebService)
	api.Path("/process/{version}").Filter(rdapi.AllGlobalFilter(getErrFunc)).Produces(restful.MIME_JSON)
	restful.DefaultRequestContentType(restful.MIME_JSON)
	restful.DefaultResponseContentType(restful.MIME_JSON)

	api.Route(api.POST("/{org_id}/{biz_id}").To(ps.CreateProcess))
	api.Route(api.DELETE("/{org_id}/{biz_id}/{process_id}").To(ps.DeleteProcess))
	api.Route(api.POST("/search/{org_id}/{biz_id}").To(ps.SearchProcess))
	api.Route(api.PUT("/{org_id}/{biz_id}/{process_id}").To(ps.UpdateProcess))
	api.Route(api.PUT("/{org_id}/{biz_id}").To(ps.BatchUpdateProcess))

	api.Route(api.GET("/module/{org_id}/{biz_id}/{process_id}").To(ps.GetProcessBindModule))
	api.Route(api.PUT("/module/{org_id}/{biz_id}/{process_id}/{module_name}").To(ps.BindModuleProcess))
	api.Route(api.DELETE("/module/{org_id}/{biz_id}/{process_id}/{module_name}").To(ps.DeleteModuleProcessBind))

	api.Route(api.GET("/{" + common.BKOwnerIDField + "}/{" + common.BKAppIDField + "}/{" + common.BKProcessIDField + "}").To(ps.GetProcessDetailByID))

	api.Route(api.POST("/operate/process").To(ps.OperateProcessInstance))
	api.Route(api.GET("/operate/process/taskresult/{taskID}").To(ps.QueryProcessOperateResult))

	api.Route(api.POST("/template/{org_id}/{biz_id}").To(ps.CreateTemplate))
	api.Route(api.PUT("/template/{org_id}/{biz_id}/{template_id}").To(ps.UpdateTemplate))
	api.Route(api.DELETE("/template/{org_id}/{biz_id}/{template_id}").To(ps.DeleteTemplate))
	api.Route(api.POST("/template/search/{org_id}/{biz_id}").To(ps.SearchTemplate))
	api.Route(api.POST("/template/version/search/{org_id}/{biz_id}/{template_id}").To(ps.SearchTemplateVersion))
	api.Route(api.POST("/template/version/{org_id}/{biz_id}/{template_id}").To(ps.CreateTemplateVersion))
	api.Route(api.PUT("/template/vesrion/{org_id}/{biz_id}/{template_id}/{version_id}").To(ps.UpdateTemplateVersion))
	api.Route(api.GET("/template/proc/{org_id}/{biz_id}/{process_id}").To(ps.GetProcBindTemplate))
	api.Route(api.PUT("/template/proc/{org_id}/{biz_id}/{process_id}/{template_id}").To(ps.BindProc2Template))
	api.Route(api.DELETE("/template/proc/{org_id}/{biz_id}/{process_id}/{template_id}").To(ps.DeleteProc2Template))
	api.Route(api.POST("/template/preview/{org_id}/{biz_id}/{template_id}").To(ps.PreviewCfg))
	api.Route(api.POST("/template/create/{org_id}/{biz_id}/{template_id}").To(ps.CreateCfg))
	api.Route(api.POST("/template/push/{org_id}/{biz_id}/{template_id}").To(ps.PushCfg))
	api.Route(api.POST("/template/getremote/{org_id}/{biz_id}/{template_id}").To(ps.GetRemoteCfg))
	api.Route(api.GET("/template/group/{org_id}/{biz_id}").To(ps.GetTemplateGroup))

	//v2
	api.Route(api.POST("/openapi/GetProcessPortByApplicationID/{" + common.BKAppIDField + "}").To(ps.GetProcessPortByApplicationID))
	api.Route(api.POST("/openapi/GetProcessPortByIP").To(ps.GetProcessPortByIP))

	api.Route(api.POST("/process/refresh/hostinstnum").To(ps.RefreshProcHostInstByEvent))

	container.Add(api)

	healthzAPI := new(restful.WebService).Produces(restful.MIME_JSON)
	healthzAPI.Route(healthzAPI.GET("/healthz").To(ps.Healthz))
	container.Add(healthzAPI)

	return container
}

func (s *ProcServer) Healthz(req *restful.Request, resp *restful.Response) {
	meta := metric.HealthMeta{IsHealthy: true}

	// zk health status
	zkItem := metric.HealthItem{IsHealthy: true, Name: types.CCFunctionalityServicediscover}
	if err := s.Engine.Ping(); err != nil {
		zkItem.IsHealthy = false
		zkItem.Message = err.Error()
	}
	meta.Items = append(meta.Items, zkItem)

	// object controller
	objCtr := metric.HealthItem{IsHealthy: true, Name: types.CC_MODULE_OBJECTCONTROLLER}
	if _, err := s.Engine.CoreAPI.Healthz().HealthCheck(types.CC_MODULE_OBJECTCONTROLLER); err != nil {
		objCtr.IsHealthy = false
		objCtr.Message = err.Error()
	}
	meta.Items = append(meta.Items, objCtr)

	// host controller
	hostCtrl := metric.HealthItem{IsHealthy: true, Name: types.CC_MODULE_HOSTCONTROLLER}
	if _, err := s.Engine.CoreAPI.Healthz().HealthCheck(types.CC_MODULE_HOSTCONTROLLER); err != nil {
		hostCtrl.IsHealthy = false
		hostCtrl.Message = err.Error()
	}

	// host controller
	procCtrl := metric.HealthItem{IsHealthy: true, Name: types.CC_MODULE_PROCCONTROLLER}
	if _, err := s.Engine.CoreAPI.Healthz().HealthCheck(types.CC_MODULE_PROCCONTROLLER); err != nil {
		procCtrl.IsHealthy = false
		procCtrl.Message = err.Error()
	}
	meta.Items = append(meta.Items, procCtrl)

	// coreservice
	coreSrv := metric.HealthItem{IsHealthy: true, Name: types.CC_MODULE_CORESERVICE}
	if _, err := s.Engine.CoreAPI.Healthz().HealthCheck(types.CC_MODULE_CORESERVICE); err != nil {
		coreSrv.IsHealthy = false
		coreSrv.Message = err.Error()
	}
	meta.Items = append(meta.Items, coreSrv)

	for _, item := range meta.Items {
		if item.IsHealthy == false {
			meta.IsHealthy = false
			meta.Message = "proc server is unhealthy"
			break
		}
	}

	info := metric.HealthInfo{
		Module:     types.CC_MODULE_HOST,
		HealthMeta: meta,
		AtTime:     metadata.Now(),
	}

	answer := metric.HealthResponse{
		Code:    common.CCSuccess,
		Data:    info,
		OK:      meta.IsHealthy,
		Result:  meta.IsHealthy,
		Message: meta.Message,
	}
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteEntity(answer)
}

func (ps *ProcServer) OnProcessConfigUpdate(previous, current cfnc.ProcessConfig) {

	//
	esbAddr, addrOk := current.ConfigMap["esb.addr"]
	esbAppCode, appCodeOk := current.ConfigMap["esb.appCode"]
	esbAppSecret, appSecretOk := current.ConfigMap["esb.appSecret"]
	if addrOk && appCodeOk && appSecretOk {
		go func() {
			ps.EsbConfigChn <- esbutil.EsbConfig{Addrs: esbAddr, AppCode: esbAppCode, AppSecret: esbAppSecret}
		}()
	}

	cfg := ccRedis.ParseConfigFromKV("redis", current.ConfigMap)
	ps.Config = &options.Config{
		Redis: &cfg,
	}

	hostInstPrefix := "host instance"
	procHostInstConfig := &ps.procHostInstConfig
	if val, ok := current.ConfigMap[hostInstPrefix+".maxEventCount"]; ok {
		eventCount, err := util.GetIntByInterface(val)
		if nil == err {
			procHostInstConfig.MaxEventCount = eventCount
		}
	}
	if val, ok := current.ConfigMap[hostInstPrefix+".maxModuleIDCount"]; ok {
		mid_count, err := util.GetIntByInterface(val)
		if nil == err {
			procHostInstConfig.MaxRefreshModuleCount = mid_count
		}
	}
	if val, ok := current.ConfigMap[hostInstPrefix+".getModuleIDInterval"]; ok {
		get_mid_interval, err := util.GetIntByInterface(val)
		if nil == err {
			procHostInstConfig.GetModuleIDInterval = time.Duration(get_mid_interval) * time.Second
		}
	}
	ps.ConfigMap = current.ConfigMap
}
