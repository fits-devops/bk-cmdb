package app

import (
	"context"
	"fmt"
	"os"
	"time"

	"configcenter/src/common"
	"configcenter/src/common/backbone"
	cc "configcenter/src/common/backbone/configcenter"
	"configcenter/src/common/blog"
	"configcenter/src/common/rdapi"
	"configcenter/src/common/types"
	"configcenter/src/common/version"
	"configcenter/src/source_controller/testservice/app/options"
	testsvr "configcenter/src/source_controller/testservice/service"
	"configcenter/src/storage/dal/mongo"
	"configcenter/src/storage/dal/redis"
)

// 定义TestServer结构体
type TestServer struct {
	Core    *backbone.Engine
	Config  options.Config
	Service testsvr.TestServiceInterface
}

//从配置文件获取配置并更新
func (t *TestServer) onTestServiceConfigUpdate(previous, current cc.ProcessConfig) {

	t.Config.Mongo = mongo.ParseConfigFromKV("mongodb", current.ConfigMap)
	t.Config.Redis = redis.ParseConfigFromKV("redis", current.ConfigMap)

	blog.V(3).Infof("the new cfg:%#v the origin cfg:%#v", t.Config, current.ConfigMap)

}

// Run main主方法，启动服务Server的入口
func Run(ctx context.Context, op *options.ServerOption) error {
	svrInfo, err := newServerInfo(op)
	if err != nil {
		return fmt.Errorf("wrap server info failed, err: %v", err)
	}

	testSvr := new(TestServer)
	testService := testsvr.New()
	testSvr.Service = testService

	webhandler := testService.WebService()
	webhandler.ServiceErrorHandler(rdapi.ServiceErrorHandler)

	input := &backbone.BackboneParameter{
		ConfigUpdate: testSvr.onTestServiceConfigUpdate,
		ConfigPath:   op.ServConf.ExConfig,
		Regdiscv:     op.ServConf.RegDiscover,
		SrvInfo:      svrInfo,
	}

	engine, err := backbone.NewBackbone(ctx, input)
	if err != nil {
		return fmt.Errorf("new backbone failed, err: %v", err)
	}

	var configReady bool
	for sleepCnt := 0; sleepCnt < common.APPConfigWaitTime; sleepCnt++ {
		// redis not found
		if "" == testSvr.Config.Redis.Address {
			time.Sleep(time.Second)
			continue
		}
		// Mongo not found
		if "" == testSvr.Config.Mongo.Address {
			time.Sleep(time.Second)
			continue
		}
		configReady = true
		break

	}

	if false == configReady {
		return fmt.Errorf("Configuration item not found")
	}

	testSvr.Core = engine
	err = testService.SetConfig(testSvr.Config, engine, engine.CCErr, engine.Language)
	if err != nil {
		return err
	}
	if err := backbone.StartServer(ctx, engine, webhandler); err != nil {
		return err
	}
	select {
	case <-ctx.Done():
	}
	return nil
}

func newServerInfo(op *options.ServerOption) (*types.ServerInfo, error) {
	ip, err := op.ServConf.GetAddress()
	if err != nil {
		return nil, err
	}

	port, err := op.ServConf.GetPort()
	if err != nil {
		return nil, err
	}

	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	info := &types.ServerInfo{
		IP:       ip,
		Port:     port,
		HostName: hostname,
		Scheme:   "http",
		Version:  version.GetVersion(),
		Pid:      os.Getpid(),
	}
	return info, nil
}
