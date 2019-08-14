package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/emicklei/go-restful"
	redis "gopkg.in/redis.v5"

	"configcenter/src/common"
	"configcenter/src/common/backbone"
	"configcenter/src/common/blog"
	"configcenter/src/common/errors"
	"configcenter/src/common/http/httpserver"
	"configcenter/src/common/language"
	"configcenter/src/common/mapstr"
	"configcenter/src/common/metadata"
	"configcenter/src/common/rdapi"
	"configcenter/src/common/util"
	"configcenter/src/source_controller/testservice/app/options"
	"configcenter/src/source_controller/testservice/core"
	"configcenter/src/source_controller/testservice/core/student"
	"configcenter/src/storage/dal"
	"configcenter/src/storage/dal/mongo/local"
	"configcenter/src/storage/dal/mongo/remote"
	dalredis "configcenter/src/storage/dal/redis"
)

// TestServiceInterface the topo service methods used to init
type TestServiceInterface interface {
	WebService() *restful.Container
	SetConfig(cfg options.Config, engin *backbone.Engine, err errors.CCErrorIf, language language.CCLanguageIf) error
}

// New create test service instance
func New() TestServiceInterface {
	return &testService{}
}

// testService test service
type testService struct {
	engin    *backbone.Engine
	language language.CCLanguageIf
	err      errors.CCErrorIf
	actions  []action
	cfg      options.Config
	core     core.Core
	db       dal.RDB
	cahce    *redis.Client
}

func (s *testService) SetConfig(cfg options.Config, engin *backbone.Engine, err errors.CCErrorIf, language language.CCLanguageIf) error {

	s.cfg = cfg
	s.engin = engin

	if nil != err {
		s.err = err
	}

	if nil != language {
		s.language = language
	}

	var db dal.DB
	var dbErr error
	if cfg.Mongo.Transaction == "enable" {
		blog.Infof("connecting to transaction manager")
		db, dbErr = remote.NewWithDiscover(engin.ServiceManageInterface.TMServer().GetServers, cfg.Mongo)
		if dbErr != nil {
			blog.Errorf("failed to connect the txc server, error info is %s", dbErr.Error())
			return dbErr
		}
	} else {
		db, dbErr = local.NewMgo(cfg.Mongo.BuildURI(), time.Minute)
		if dbErr != nil {
			blog.Errorf("failed to connect the remote server(%s), error info is %s", cfg.Mongo.BuildURI(), dbErr.Error())
			return dbErr
		}
	}
	cache, cacheRrr := dalredis.NewFromConfig(cfg.Redis)
	if cacheRrr != nil {
		blog.Errorf("new redis client failed, err: %v", cacheRrr)
		return cacheRrr
	}

	s.db = db
	s.cahce = cache

	// connect the remote mongodb
	s.core = core.New(
		student.New(db),
	)
	return nil
}

// WebService the web service
func (s *testService) WebService() *restful.Container {

	container := restful.NewContainer()

	// init service actions
	s.initService()

	api := new(restful.WebService)
	getErrFunc := func() errors.CCErrorIf {
		return s.err
	}
	api.Path("/api/v3").Filter(rdapi.AllGlobalFilter(getErrFunc)).Produces(restful.MIME_JSON)

	innerActions := s.Actions()

	for _, actionItem := range innerActions {
		switch actionItem.Verb {
		case http.MethodPost:
			api.Route(api.POST(actionItem.Path).To(actionItem.Handler))
		case http.MethodDelete:
			api.Route(api.DELETE(actionItem.Path).To(actionItem.Handler))
		case http.MethodPut:
			api.Route(api.PUT(actionItem.Path).To(actionItem.Handler))
		case http.MethodGet:
			api.Route(api.GET(actionItem.Path).To(actionItem.Handler))
		default:
			blog.Errorf(" the url (%s), the http method (%s) is not supported", actionItem.Path, actionItem.Verb)
		}
	}

	container.Add(api)

	healthzAPI := new(restful.WebService).Produces(restful.MIME_JSON)
	healthzAPI.Route(healthzAPI.GET("/healthz").To(s.Healthz))
	container.Add(healthzAPI)

	return container
}

func (s *testService) createAPIRspStr(errcode int, info interface{}) (string, error) {

	rsp := metadata.Response{
		BaseResp: metadata.SuccessBaseResp,
		Data:     nil,
	}

	if common.CCSuccess != errcode {
		rsp.Code = errcode
		rsp.Result = false
		rsp.ErrMsg = fmt.Sprintf("%v", info)
	} else {
		rsp.ErrMsg = common.CCSuccessStr
		rsp.Data = info
	}

	data, err := json.Marshal(rsp)
	return string(data), err
}

func (s *testService) createCompleteAPIRspStr(errcode int, errmsg string, info interface{}) (string, error) {

	rsp := metadata.Response{
		BaseResp: metadata.SuccessBaseResp,
		Data:     nil,
	}

	if common.CCSuccess != errcode {
		rsp.Code = errcode
		rsp.Result = false
		rsp.ErrMsg = errmsg
	} else {
		rsp.ErrMsg = common.CCSuccessStr
	}
	rsp.Data = info
	data, err := json.Marshal(rsp)
	return string(data), err
}

func (s *testService) sendResponse(resp *restful.Response, errorCode int, dataMsg interface{}) {
	resp.Header().Set("Content-Type", "application/json")
	if rsp, rspErr := s.createAPIRspStr(errorCode, dataMsg); nil == rspErr {
		io.WriteString(resp, rsp)
	} else {
		blog.Errorf("failed to send response , error info is %s", rspErr.Error())
	}
}

func (s *testService) sendCompleteResponse(resp *restful.Response, errorCode int, errMsg string, info interface{}) {
	resp.Header().Set("Content-Type", "application/json")
	rsp, rspErr := s.createCompleteAPIRspStr(errorCode, errMsg, info)
	if nil == rspErr {
		io.WriteString(resp, rsp)
		return
	}
	blog.Errorf("failed to send response , error info is %s", rspErr.Error())

}

func (s *testService) addAction(method string, path string, handlerFunc LogicFunc, handlerParseOriginDataFunc ParseOriginDataFunc) {
	actionObject := action{
		Method:                     method,
		Path:                       path,
		HandlerFunc:                handlerFunc,
		HandlerParseOriginDataFunc: handlerParseOriginDataFunc,
	}
	s.actions = append(s.actions, actionObject)
}

// Actions return the all actions
func (s *testService) Actions() []*httpserver.Action {

	var httpactions []*httpserver.Action
	for _, a := range s.actions {

		func(act action) {

			httpactions = append(httpactions, &httpserver.Action{Verb: act.Method, Path: act.Path, Handler: func(req *restful.Request, resp *restful.Response) {
				rid := util.GetHTTPCCRequestID(req.Request.Header)

				ownerID := util.GetOwnerID(req.Request.Header)
				user := util.GetUser(req.Request.Header)

				// get the language
				language := util.GetLanguage(req.Request.Header)

				defLang := s.language.CreateDefaultCCLanguageIf(language)

				// get the error info by the language
				defErr := s.err.CreateDefaultCCErrorIf(language)

				value, err := ioutil.ReadAll(req.Request.Body)
				if err != nil {
					blog.Errorf("read http request body failed, err: %+v, rid: %s", err, rid)
					errStr := defErr.Error(common.CCErrCommHTTPReadBodyFailed)
					s.sendResponse(resp, common.CCErrCommHTTPReadBodyFailed, errStr)
					return
				}

				mData := mapstr.MapStr{}
				if nil == act.HandlerParseOriginDataFunc {
					if err := json.Unmarshal(value, &mData); nil != err && 0 != len(value) {
						blog.Errorf("failed to unmarshal the data, err: %+v, rid: %s", err, rid)
						errStr := defErr.Error(common.CCErrCommJSONUnmarshalFailed)
						s.sendResponse(resp, common.CCErrCommJSONUnmarshalFailed, errStr)
						return
					}
				} else {
					mData, err = act.HandlerParseOriginDataFunc(value)
					if nil != err {
						blog.Errorf("failed to unmarshal the data, err: %+v, rid: %s", err, rid)
						errStr := defErr.Error(common.CCErrCommJSONUnmarshalFailed)
						s.sendResponse(resp, common.CCErrCommJSONUnmarshalFailed, errStr)
						return
					}
				}

				data, dataErr := act.HandlerFunc(core.ContextParams{
					Context:         util.GetDBContext(context.Background(), req.Request.Header),
					Error:           defErr,
					Lang:            defLang,
					Header:          req.Request.Header,
					SupplierAccount: ownerID,
					ReqID:           rid,
					User:            user,
				},
					req.PathParameter,
					req.QueryParameter,
					mData)

				if nil != dataErr {
					switch e := dataErr.(type) {
					default:
						s.sendCompleteResponse(resp, common.CCSystemBusy, dataErr.Error(), data)
					case errors.CCErrorCoder:
						s.sendCompleteResponse(resp, e.GetCode(), dataErr.Error(), data)
					}
					return
				}

				s.sendResponse(resp, common.CCSuccess, data)

			}})
		}(a)

	}
	return httpactions
}
