package v3v0v8

import (
	"context"
	"fmt"

	"strings"
	"time"

	"configcenter/src/common"
	"configcenter/src/common/auditoplog"
	"configcenter/src/common/blog"
	"configcenter/src/common/mapstr"
	"configcenter/src/common/metadata"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

// 进程:功能:port
var prc2port = []string{
	"cmdb_coreservice:coreservice:30001",
	"cmdb_toposerver:toposerver:40001",
	"cmdb_adminserver:adminserver:40002",
	"cmdb_datacollection:datacollection:40003",
	"cmdb_eventserver:eventserver:40004",
	"cmdb_apiserver:apiserver:8080",
	"cmdb_webserver:webserver:8083"}

// 集群:模块:进程
var setModuleKv = map[string]map[string]string{
	"配置平台": {
		"coreservice":		"cmdb_coreservice",
		"toposerver":       "cmdb_toposerver",
		"adminserver":      "cmdb_adminserver",
		"datacollection":   "cmdb_datacollection",
		"eventserver":      "cmdb_eventserver",
		"apiserver":        "cmdb_apiserver",
		"webserver":        "cmdb_webserver",
	},
	"公共组件": {
		"mysql": "common_mysql",
		"redis": "common_redis",
		"redis_cluster": "redis_cluster",
		"zookeeper": "zk_java",
		"etcd": "etcd",
		"mongodb": "mongodb"},
}

var procName2ID map[string]uint64

//addBKApp add bk app
func addBKApp(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {

	if count, err := db.Table("cc_ApplicationBase").Find(mapstr.MapStr{common.BKAppNameField: common.BKAppName}).Count(ctx); err != nil {
		return err
	} else if count >= 1 {
		return nil
	}

	// add bk app
	appModelData := map[string]interface{}{}
	appModelData[common.BKAppNameField] = common.BKAppName
	appModelData[common.BKMaintainersField] = admin
	appModelData[common.BKTimeZoneField] = "Asia/Shanghai"
	appModelData[common.BKLanguageField] = "1" //"中文"
	appModelData[common.BKLifeCycleField] = common.DefaultAppLifeCycleNormal
	appModelData[common.BKOwnerIDField] = conf.OwnerID
	appModelData[common.BKDefaultField] = 0
	appModelData[common.BKSupplierIDField] = conf.SupplierID
	filled := fillEmptyFields(appModelData, AppRow())
	var preData map[string]interface{}
	bizID, preData, err := upgrader.Upsert(ctx, db, "cc_ApplicationBase", appModelData, common.BKAppIDField, []string{common.BKAppNameField, common.BKOwnerIDField}, append(filled, common.BKAppIDField))
	if err != nil {
		blog.Error("add addBKApp error ", err.Error())
		return err
	}

	// add audit log
	headers := []metadata.Header{}
	for _, item := range AppRow() {
		headers = append(headers, metadata.Header{
			PropertyID:   item.PropertyID,
			PropertyName: item.PropertyName,
		})
	}
	auditContent := metadata.Content{
		CurData: appModelData,
		Headers: headers,
	}
	logRow := &metadata.OperationLog{
		OwnerID:       conf.OwnerID,
		ApplicationID: int64(bizID),
		OpType:        int(auditoplog.AuditOpTypeAdd),
		OpTarget:      "biz",
		User:          conf.User,
		ExtKey:        "",
		OpDesc:        "create app",
		Content:       auditContent,
		CreateTime:    time.Now(),
		InstID:        int64(bizID),
	}
	if preData != nil {
		logRow.OpDesc = "update process"
		logRow.OpType = int(auditoplog.AuditOpTypeModify)
	}
	if err = db.Table(logRow.TableName()).Insert(ctx, logRow); err != nil {
		blog.Error("add audit log error ", err.Error())
		return err
	}

	//// add bk app default set
	//inputSetInfo := make(map[string]interface{})
	//inputSetInfo[common.BKAppIDField] = bizID
	//inputSetInfo[common.BKInstParentStr] = bizID
	//inputSetInfo[common.BKSetNameField] = common.DefaultResSetName
	//inputSetInfo[common.BKDefaultField] = common.DefaultResSetFlag
	//inputSetInfo[common.BKOwnerIDField] = conf.OwnerID
	//filled = fillEmptyFields(inputSetInfo, SetRow())
	//setID, _, err := upgrader.Upsert(ctx, db, "cc_SetBase", inputSetInfo, common.BKSetIDField, []string{common.BKOwnerIDField, common.BKAppIDField, common.BKSetNameField}, append(filled, common.BKSetIDField))
	//if err != nil {
	//	blog.Error("add defaultSet error ", err.Error())
	//	return err
	//}

	//// add bk app default module
	//inputResModuleInfo := make(map[string]interface{})
	//inputResModuleInfo[common.BKSetIDField] = setID
	//inputResModuleInfo[common.BKInstParentStr] = setID
	//inputResModuleInfo[common.BKAppIDField] = bizID
	//inputResModuleInfo[common.BKModuleNameField] = common.DefaultResModuleName
	//inputResModuleInfo[common.BKDefaultField] = common.DefaultResModuleFlag
	//inputResModuleInfo[common.BKOwnerIDField] = conf.OwnerID
	//filled = fillEmptyFields(inputResModuleInfo, ModuleRow())
	//_, _, err = upgrader.Upsert(ctx, db, "cc_ModuleBase", inputResModuleInfo, common.BKModuleIDField, []string{common.BKOwnerIDField, common.BKModuleNameField, common.BKAppIDField, common.BKSetIDField}, append(filled, common.BKModuleIDField))
	//if err != nil {
	//	blog.Error("add defaultResModule error ", err.Error())
	//	return err
	//}

	//inputFaultModuleInfo := make(map[string]interface{})
	//inputFaultModuleInfo[common.BKSetIDField] = setID
	//inputFaultModuleInfo[common.BKInstParentStr] = setID
	//inputFaultModuleInfo[common.BKAppIDField] = bizID
	//inputFaultModuleInfo[common.BKModuleNameField] = common.DefaultFaultModuleName
	//inputFaultModuleInfo[common.BKDefaultField] = common.DefaultFaultModuleFlag
	//inputFaultModuleInfo[common.BKOwnerIDField] = conf.OwnerID
	//filled = fillEmptyFields(inputFaultModuleInfo, ModuleRow())
	//_, _, err = upgrader.Upsert(ctx, db, "cc_ModuleBase", inputFaultModuleInfo, common.BKModuleIDField, []string{common.BKOwnerIDField, common.BKModuleNameField, common.BKAppIDField, common.BKSetIDField}, append(filled, common.BKModuleIDField))
	//if err != nil {
	//	blog.Error("add defaultFaultModule error ", err.Error())
	//	return err
	//}

	if err := addBKProcess(ctx, db, conf, bizID); err != nil {
		blog.Error("add addBKProcess error ", err.Error())
	}
	if err := addSetInBKApp(ctx, db, conf, bizID); err != nil {
		blog.Error("add addSetInBKApp error ", err.Error())
	}

	return nil
}

//addBKProcess add bk process
func addBKProcess(ctx context.Context, db dal.RDB, conf *upgrader.Config, bizID uint64) error {
	procName2ID = make(map[string]uint64)

	for _, procStr := range prc2port {
		procArr := strings.Split(procStr, ":")
		procName := procArr[0]
		funcName := procArr[1]
		portStr := procArr[2]
		var protocol string
		if len(procArr) > 3 {
			protocol = procArr[3]
		}
		procModelData := map[string]interface{}{}
		procModelData[common.BKProcessNameField] = procName
		procModelData[common.BKFuncName] = funcName
		procModelData[common.BKPort] = portStr
		procModelData[common.BKWorkPath] = "/data/bkee"
		procModelData[common.BKOwnerIDField] = conf.OwnerID
		procModelData[common.BKAppIDField] = bizID

		protocol = strings.ToLower(protocol)
		switch protocol {
		case "udp":
			procModelData[common.BKProtocol] = "2"
		case "tcp":
			procModelData[common.BKProtocol] = "1"
		default:
			procModelData[common.BKProtocol] = "1"
		}

		filled := fillEmptyFields(procModelData, ProcRow())
		var preData map[string]interface{}
		processID, preData, err := upgrader.Upsert(ctx, db, "cc_Process", procModelData, common.BKProcessIDField, []string{common.BKProcessNameField, common.BKAppIDField, common.BKOwnerIDField}, append(filled, common.BKProcessIDField))
		if err != nil {
			blog.Error("add addBKProcess error ", err.Error())
			return err
		}
		procName2ID[procName] = processID

		// add audit log
		headers := []metadata.Header{}
		for _, item := range ProcRow() {
			headers = append(headers, metadata.Header{
				PropertyID:   item.PropertyID,
				PropertyName: item.PropertyName,
			})
		}
		auditContent := metadata.Content{
			CurData: procModelData,
			Headers: headers,
		}
		logRow := &metadata.OperationLog{
			OwnerID:       conf.OwnerID,
			ApplicationID: int64(bizID),
			OpType:        int(auditoplog.AuditOpTypeAdd),
			OpTarget:      "process",
			User:          conf.User,
			ExtKey:        "",
			OpDesc:        "create process",
			Content:       auditContent,
			CreateTime:    time.Now(),
			InstID:        int64(processID),
		}
		if preData != nil {
			logRow.OpDesc = "update process"
			logRow.OpType = int(auditoplog.AuditOpTypeModify)
		}
		if err = db.Table(logRow.TableName()).Insert(ctx, logRow); err != nil {
			blog.Error("add audit log error ", err.Error())
			return err
		}

	}

	return nil
}

//addSetInBKApp add set in bk app
func addSetInBKApp(ctx context.Context, db dal.RDB, conf *upgrader.Config, bizID uint64) error {
	for setName, moduleArr := range setModuleKv {
		setModelData := map[string]interface{}{}
		setModelData[common.BKSetNameField] = setName
		setModelData[common.BKAppIDField] = bizID
		setModelData[common.BKOwnerIDField] = conf.OwnerID
		setModelData[common.BKInstParentStr] = bizID
		setModelData[common.BKDefaultField] = 0
		setModelData[common.CreateTimeField] = time.Now()
		setModelData[common.LastTimeField] = time.Now()
		filled := fillEmptyFields(setModelData, SetRow())
		var preData map[string]interface{}
		setID, preData, err := upgrader.Upsert(ctx, db, "cc_SetBase", setModelData, common.BKSetIDField, []string{common.BKSetNameField, common.BKOwnerIDField, common.BKAppIDField}, append(filled, common.BKSetIDField))
		if err != nil {
			blog.Error("add addSetInBKApp error ", err.Error())
			return err
		}

		// add audit log
		headers := []metadata.Header{}
		for _, item := range SetRow() {
			headers = append(headers, metadata.Header{
				PropertyID:   item.PropertyID,
				PropertyName: item.PropertyName,
			})
		}
		auditContent := metadata.Content{
			CurData: setModelData,
			Headers: headers,
		}
		logRow := &metadata.OperationLog{
			OwnerID:       conf.OwnerID,
			ApplicationID: int64(bizID),
			OpType:        int(auditoplog.AuditOpTypeAdd),
			OpTarget:      "set",
			User:          conf.User,
			ExtKey:        "",
			OpDesc:        "create set",
			Content:       auditContent,
			CreateTime:    time.Now(),
			InstID:        int64(setID),
		}
		if preData != nil {
			logRow.OpDesc = "update set"
			logRow.OpType = int(auditoplog.AuditOpTypeModify)
		}
		if err = db.Table(logRow.TableName()).Insert(ctx, logRow); err != nil {
			blog.Error("add audit log error ", err.Error())
			return err
		}

		// add module in set
		if err := addModuleInSet(ctx, db, conf, moduleArr, setID, bizID); err != nil {
			return err
		}
	}
	return nil
}

//addModuleInSet add module in set
func addModuleInSet(ctx context.Context, db dal.RDB, conf *upgrader.Config, moduleArr map[string]string, setID, bizID uint64) error {
	for moduleName, processNameStr := range moduleArr {
		moduleModelData := map[string]interface{}{}
		moduleModelData[common.BKModuleNameField] = moduleName
		moduleModelData[common.BKAppIDField] = bizID
		moduleModelData[common.BKSetIDField] = setID
		moduleModelData[common.BKOwnerIDField] = conf.OwnerID
		moduleModelData[common.BKInstParentStr] = setID
		moduleModelData[common.BKDefaultField] = 0
		var preData map[string]interface{}
		filled := fillEmptyFields(moduleModelData, ModuleRow())
		moduleID, preData, err := upgrader.Upsert(ctx, db, "cc_ModuleBase", moduleModelData, common.BKModuleIDField, []string{common.BKModuleNameField, common.BKOwnerIDField, common.BKAppIDField, common.BKSetIDField},
			append(filled, common.BKModuleIDField))
		if err != nil {
			blog.Error("add addModuleInSet error ", err.Error())
			return err
		}

		// add audit log
		headers := []metadata.Header{}
		for _, item := range ModuleRow() {
			headers = append(headers, metadata.Header{
				PropertyID:   item.PropertyID,
				PropertyName: item.PropertyName,
			})
		}
		auditContent := metadata.Content{
			CurData: moduleModelData,
			PreData: preData,
			Headers: headers,
		}
		logRow := &metadata.OperationLog{
			OwnerID:       conf.OwnerID,
			ApplicationID: int64(bizID),
			OpType:        int(auditoplog.AuditOpTypeAdd),
			OpTarget:      "module",
			User:          conf.User,
			ExtKey:        "",
			OpDesc:        "create module",
			Content:       auditContent,
			CreateTime:    time.Now(),
			InstID:        int64(moduleID),
		}
		if preData != nil {
			logRow.OpDesc = "update module"
			logRow.OpType = int(auditoplog.AuditOpTypeModify)
		}
		if err = db.Table(logRow.TableName()).Insert(ctx, logRow); err != nil {
			blog.Error("add audit log error ", err.Error())
			return err
		}

		//add module process config
		if err := addModule2Process(ctx, db, conf, processNameStr, moduleName, bizID); err != nil {
			return err
		}

	}
	return nil
}

//addModule2Process add process 2 module
func addModule2Process(ctx context.Context, db dal.RDB, conf *upgrader.Config, processNameStr string, moduleName string, bizID uint64) (err error) {
	processNameArr := strings.Split(processNameStr, ",")
	for _, processName := range processNameArr {
		processID, ok := procName2ID[processName]
		if false == ok {
			continue
		}
		module2Process := map[string]interface{}{}
		module2Process[common.BKAppIDField] = bizID
		module2Process[common.BKModuleNameField] = moduleName
		module2Process[common.BKProcessIDField] = processID

		if _, _, err = upgrader.Upsert(ctx, db, "cc_Proc2Module", module2Process, "", []string{common.BKModuleNameField, common.BKAppIDField, common.BKProcessIDField}, nil); err != nil {
			blog.Error("add addModuleInSet error ", err.Error())
			return err
		}

		// add audit log
		headers := []metadata.Header{}
		for _, item := range ModuleRow() {
			headers = append(headers, metadata.Header{
				PropertyID:   item.PropertyID,
				PropertyName: item.PropertyName,
			})
		}
		logRow := &metadata.OperationLog{
			OwnerID:       conf.OwnerID,
			ApplicationID: int64(bizID),
			OpType:        int(auditoplog.AuditOpTypeModify),
			OpTarget:      "module",
			User:          conf.User,
			ExtKey:        "",
			OpDesc:        fmt.Sprintf("bind module [%s]", moduleName),
			Content:       "",
			CreateTime:    time.Now(),
			InstID:        int64(bizID),
		}
		if err = db.Table(logRow.TableName()).Insert(ctx, logRow); err != nil {
			blog.Error("add audit log error ", err.Error())
			return err
		}
	}
	return nil
}
