package v3v0v8

import (
	"context"
	"time"

	"configcenter/src/common"
	"configcenter/src/common/blog"
	"configcenter/src/common/metadata"
	mCommon "configcenter/src/scene_server/admin_server/common"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func addPresetObjects(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	//初始化模型分类基础数据
	err = addClassifications(ctx, db, conf)
	if err != nil {
		return err
	}
	//初始化属性分组基础数据
	err = addPropertyGroupData(ctx, db, conf)
	if err != nil {
		return err
	}
	//初始化模型基础数据
	err = addObjDesData(ctx, db, conf)
	if err != nil {
		return err
	}
	//初始化模型字段基础数据
	err = addObjAttDescData(ctx, db, conf)
	if err != nil {
		return err
	}
	//初始化模型关联基础数据
	err = addAsstData(ctx, db, conf)
	if err != nil {
		return err
	}

	//初始化模型拓扑基础数据
	err = addModeTopoData(ctx, db, conf)
	if err != nil {
		return err
	}

	return nil
}


func addModeTopoData(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {
	tablename := common.BKTableNameTopoGraphics
	blog.Errorf("add data for  %s table ", tablename)
	rows := getModeTopoData(conf.OwnerID)
	for _, row := range rows {
		if _, _, err := upgrader.Upsert(ctx, db, tablename, row, "", []string{common.BKObjIDField},[]string{}); err != nil {
			blog.Errorf("add data for  %s table error  %s", tablename, err)
			return err
		}
	}

	return nil
}


func addAsstData(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {
	tablename := common.BKTableNameObjAsst
	blog.Errorf("add data for  %s table ", tablename)
	rows := getAddAsstData(conf.OwnerID)
	for _, row := range rows {
		// topo mainline could be changed,so need to ignore asst_obj_id
		_, _, err := upgrader.Upsert(ctx, db, tablename, row, "id", []string{common.BKObjIDField, common.BKObjAttIDField, common.BKOwnerIDField}, []string{"id", "asst_obj_id"})
		if nil != err {
			blog.Errorf("add data for  %s table error  %s", tablename, err)
			return err
		}
	}

	blog.Errorf("add data for  %s table  ", tablename)
	return nil
}

func addObjAttDescData(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {
	tablename := common.BKTableNameObjAttDes
	blog.Infof("add data for  %s table ", tablename)
	rows := getObjAttDescData(conf.OwnerID)
	for _, row := range rows {
		_, _, err := upgrader.Upsert(ctx, db, tablename, row, "id", []string{common.BKObjIDField, common.BKPropertyIDField, common.BKOwnerIDField}, []string{})
		if nil != err {
			blog.Errorf("add data for  %s table error  %s", tablename, err)
			return err
		}
	}
	selector := map[string]interface{}{
		common.BKObjIDField: map[string]interface{}{
			common.BKDBIN: []string{"switch",
				"router",
				"load_balance",
				"firewall",
				"weblogic_service",
				"tomcat_service",
				"apache_service",
			},
		},
		common.BKPropertyIDField: "bk_name",
	}

	db.Table(tablename).Delete(ctx, selector)

	return nil
}

func addObjDesData(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {
	tablename := common.BKTableNameObjDes
	blog.Errorf("add data for  %s table ", tablename)
	rows := getObjectDesData(conf.OwnerID)
	for _, row := range rows {
		if _, _, err := upgrader.Upsert(ctx, db, tablename, row, "id", []string{common.BKObjIDField, common.BKClassificationIDField, common.BKOwnerIDField}, []string{"id"}); err != nil {
			blog.Errorf("add data for  %s table error  %s", tablename, err)
			return err
		}
	}

	return nil
}

func addClassifications(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	tablename := common.BKTableNameObjClassifiction
	blog.Infof("add %s rows", tablename)
	for _, row := range classificationRows {
		if _, _, err = upgrader.Upsert(ctx, db, tablename, row, "id", []string{common.BKClassificationIDField}, []string{"id"}); err != nil {
			blog.Errorf("add data for  %s table error  %s", tablename, err)
			return err
		}
	}
	return
}

func addPropertyGroupData(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {
	tablename := common.BKTableNamePropertyGroup
	blog.Errorf("add data for  %s table ", tablename)
	rows := getPropertyGroupData(conf.OwnerID)
	for _, row := range rows {
		if _, _, err := upgrader.Upsert(ctx, db, tablename, row, "id", []string{common.BKObjIDField, "group_id"}, []string{"id"}); err != nil {
			blog.Errorf("add data for  %s table error  %s", tablename, err)
			return err
		}
	}
	return nil
}
func getObjectDesData(ownerID string) []*metadata.Object {

	dataRows := []*metadata.Object{
		//bk_iaas（基础设施）
		&metadata.Object{ObjCls: "bk_iaas", ObjectID: common.BKInnerObjIDHost, ObjectName: "主机", IsPre: true, ObjIcon: "icon-cc-host", Position: `{"bk_iaas":{"x":-188,"y":18}}`},
		&metadata.Object{ObjCls: "bk_iaas", ObjectID: common.BKInnerObjIDStorage, ObjectName: "存储", IsPre: true, ObjIcon: "icon-cc-storage", Position: `{"bk_iaas":{"x":123,"y":-95}}`},
		&metadata.Object{ObjCls: "bk_iaas", ObjectID: common.BKInnerObjIDIdc, ObjectName: "机房", IsPre: true, ObjIcon: "icon-cc-engine-room", Position: `{"bk_iaas":{"x":-180,"y":417}}`},
		&metadata.Object{ObjCls: "bk_iaas", ObjectID: common.BKInnerObjIDIdcRack, ObjectName: "机柜", IsPre: true, ObjIcon: "icon-cc-cabinet", Position: `{"bk_iaas":{"x":-184,"y":211}}`},
		&metadata.Object{ObjCls: "bk_iaas", ObjectID: common.BKInnerObjIDRouter, ObjectName: "路由器", IsPre: true, ObjIcon: "icon-cc-router", Position: `{"bk_iaas":{"x":127,"y":27}}`},
		&metadata.Object{ObjCls: "bk_iaas", ObjectID: common.BKInnerObjIDSwitch, ObjectName: "交换机", IsPre: true, ObjIcon: "icon-cc-switch2", Position: `{"bk_iaas":{"x":133,"y":150}}`},
		&metadata.Object{ObjCls: "bk_iaas", ObjectID: common.BKInnerObjIDFirewall, ObjectName: "防火墙", IsPre: true, ObjIcon: "icon-cc-firewall", Position: `{"bk_iaas":{"x":137,"y":269}}`},
		&metadata.Object{ObjCls: "bk_iaas", ObjectID: common.BKInnerObjIDBlance, ObjectName: "负载均衡", IsPre: true, ObjIcon: "icon-cc-balance", Position: `{"bk_iaas":{"x":141,"y":401}}`},

		//bk_paas（平台资源）
		&metadata.Object{ObjCls: "bk_paas", ObjectID: common.BKInnerObjIDNginx, ObjectName: "Nginx服务", ObjIcon: "icon-cc-nginx", Position: `{"bk_paas":{"x":0,"y":0}}`},
		&metadata.Object{ObjCls: "bk_paas", ObjectID: common.BKInnerObjIDZookeeper, ObjectName: "Zookeeper服务", ObjIcon: "icon-cc-kafka", Position: `{"bk_paas":{"x":0,"y":0}}`},
		&metadata.Object{ObjCls: "bk_paas", ObjectID: common.BKInnerObjIDMysql, ObjectName: "Mysql服务", ObjIcon: "icon-cc-mysql", Position: `{"bk_paas":{"x":0,"y":0}}`},
		&metadata.Object{ObjCls: "bk_paas", ObjectID: common.BKInnerObjIDMongo, ObjectName: "Mongo服务", ObjIcon: "icon-cc-mongodb", Position: `{"bk_paas":{"x":0,"y":0}}`},
		&metadata.Object{ObjCls: "bk_paas", ObjectID: common.BKInnerObjIDWeblogic, ObjectName: "weblogic", ObjIcon: "icon-cc-weblogic", Position: `{"bk_paas":{"x":0,"y":0}}`},
		&metadata.Object{ObjCls: "bk_paas", ObjectID: common.BKInnerObjIDApache, ObjectName: "apache", ObjIcon: "icon-cc-apache", Position: `{"bk_paas":{"x":0,"y":0}}`},
		&metadata.Object{ObjCls: "bk_paas", ObjectID: common.BKInnerObjIDTomcat, ObjectName: "tomcat", ObjIcon: "icon-cc-tomcat", Position: `{"bk_paas":{"x":0,"y":0}}`},

		//bk_saas（应用资源）
		&metadata.Object{ObjCls: "bk_saas", ObjectID: common.BKInnerObjIDApp, ObjectName: "业务系统", IsPre: true, ObjIcon: "icon-cc-business", Position: `{"bk_saas":{"x":-782,"y":-191}}`},
		&metadata.Object{ObjCls: "bk_saas", ObjectID: common.BKInnerObjIDSet, ObjectName: "集群", IsPre: true, ObjIcon: "icon-cc-set", Position: `{"bk_saas":{"x":-472,"y":-191}}`},
		&metadata.Object{ObjCls: "bk_saas", ObjectID: common.BKInnerObjIDModule, ObjectName: "模块", IsPre: true, ObjIcon: "icon-cc-module", Position: `{"bk_saas":{"x":-470,"y":17}}`},
		&metadata.Object{ObjCls: "bk_saas", ObjectID: common.BKInnerObjIDProc, ObjectName: "进程服务", IsPre: true, ObjIcon: "icon-cc-process", Position: `{"bk_saas":{"x":0,"y":0}}`},
		&metadata.Object{ObjCls: "bk_saas", ObjectID: common.BKInnerObjIDPlat, ObjectName: "云区域", IsPre: true, ObjIcon: "icon-cc-subnet", Position: `{"bk_saas":{"x":-190,"y":-193}}`},

		//bk_organization (组织信息)
		&metadata.Object{ObjCls: "bk_organization", ObjectID: common.BKInnerObjIDUser, ObjectName: "用户", IsPre: true, ObjIcon: "icon-cc-group", Position: `{"bk_organization":{"x":-470,"y":209}}`},
		&metadata.Object{ObjCls: "bk_organization", ObjectID: common.BKInnerObjIDUserGroup, ObjectName: "用户组", IsPre: true, ObjIcon: "icon-cc-department", Position: `{"bk_organization":{"x":-780,"y":204}}`},
	}
	t := metadata.Now()
	for _, r := range dataRows {
		r.CreateTime = &t
		r.LastTime = &t
		r.IsPaused = false
		r.Creator = common.CCSystemOperatorUserName
		r.OwnerID = ownerID
		r.Description = ""
		r.Modifier = ""
	}

	return dataRows
}

// Association for purpose of this structure not change by other, copy here
type Association struct {
	ID               int64  `field:"id" json:"id" bson:"id"`
	ObjectID         string `field:"obj_id" json:"obj_id" bson:"obj_id"`
	OwnerID          string `field:"org_id" json:"org_id" bson:"org_id"`
	AsstForward      string `field:"bk_asst_forward" json:"bk_asst_forward" bson:"bk_asst_forward"`
	AsstObjID        string `field:"asst_obj_id" json:"asst_obj_id" bson:"asst_obj_id"`
	AsstName         string `field:"asst_name" json:"asst_name" bson:"asst_name"`
	ObjectAttID      string `field:"bk_object_att_id" json:"bk_object_att_id" bson:"bk_object_att_id"`
	ClassificationID string `field:"classification_id" bson:"-"`
	ObjectIcon       string `field:"obj_icon" bson:"-"`
	ObjectName       string `field:"obj_name" bson:"-"`
}

func getAddAsstData(ownerID string) []Association {
	dataRows := []Association{
		{OwnerID: ownerID, ObjectID: common.BKInnerObjIDSet, ObjectAttID: common.BKChildStr, AsstObjID: common.BKInnerObjIDApp},
		{OwnerID: ownerID, ObjectID: common.BKInnerObjIDModule, ObjectAttID: common.BKChildStr, AsstObjID: common.BKInnerObjIDSet},
		{OwnerID: ownerID, ObjectID: common.BKInnerObjIDHost, ObjectAttID: common.BKChildStr, AsstObjID: common.BKInnerObjIDModule},
		{OwnerID: ownerID, ObjectID: common.BKInnerObjIDHost, ObjectAttID: common.BKCloudIDField, AsstObjID: common.BKInnerObjIDPlat},
	}
	return dataRows
}

func getObjAttDescData(ownerID string) []*Attribute {

	predataRows := AppRow()
	predataRows = append(predataRows, SetRow()...)
	predataRows = append(predataRows, ModuleRow()...)
	predataRows = append(predataRows, HostRow()...)
	predataRows = append(predataRows, ProcRow()...)
	predataRows = append(predataRows, PlatRow()...)

	dataRows := SwitchRow()
	dataRows = append(dataRows, RouterRow()...)
	dataRows = append(dataRows, LoadBalanceRow()...)
	dataRows = append(dataRows, FirewallRow()...)
	dataRows = append(dataRows, StorageRow()...)
	dataRows = append(dataRows, IdcRow()...)
	dataRows = append(dataRows, IDCRackRow()...)

	dataRows = append(dataRows, NginxRow()...)
	dataRows = append(dataRows, ZookeeperRow()...)
	dataRows = append(dataRows, MysqlRow()...)
	dataRows = append(dataRows, MongoRow()...)
	dataRows = append(dataRows, WeblogicRow()...)
	dataRows = append(dataRows, ApacheRow()...)
	dataRows = append(dataRows, TomcatRow()...)

	dataRows = append(dataRows, UserRow()...)
	dataRows = append(dataRows, UserGroupRow()...)

	t := new(time.Time)
	*t = time.Now()
	for _, r := range predataRows {
		r.OwnerID = ownerID
		//r.IsPre = true
		if false != r.IsEditable {
			r.IsEditable = true
		}
		r.IsReadOnly = false
		r.CreateTime = t
		r.Creator = common.CCSystemOperatorUserName
		r.LastTime = r.CreateTime
		r.Description = ""
	}
	for _, r := range dataRows {
		r.OwnerID = ownerID
		if false != r.IsEditable {
			r.IsEditable = true
		}
		r.IsReadOnly = false
		r.CreateTime = t
		r.Creator = common.CCSystemOperatorUserName
		r.LastTime = r.CreateTime
		r.Description = ""
	}

	return append(predataRows, dataRows...)
}

func getPropertyGroupData(ownerID string) []*metadata.Group {
	objectIDs := make(map[string]map[string]string)

	dataRows := []*metadata.Group{
		//app
		&metadata.Group{ObjectID: common.BKInnerObjIDApp, GroupID: mCommon.BaseInfo, GroupName: mCommon.BaseInfoName, GroupIndex: 1, OwnerID: ownerID, IsDefault: true},
		&metadata.Group{ObjectID: common.BKInnerObjIDApp, GroupID: mCommon.AppRole, GroupName: mCommon.AppRoleName, GroupIndex: 2, OwnerID: ownerID, IsDefault: true},

		//set
		&metadata.Group{ObjectID: common.BKInnerObjIDSet, GroupID: mCommon.BaseInfo, GroupName: mCommon.BaseInfoName, GroupIndex: 1, OwnerID: ownerID, IsDefault: true},

		//module
		&metadata.Group{ObjectID: common.BKInnerObjIDModule, GroupID: mCommon.BaseInfo, GroupName: mCommon.BaseInfoName, GroupIndex: 1, OwnerID: ownerID, IsDefault: true},

		//host
		&metadata.Group{ObjectID: common.BKInnerObjIDHost, GroupID: mCommon.BaseInfo, GroupName: mCommon.BaseInfoName, GroupIndex: 1, OwnerID: ownerID, IsDefault: true},
		&metadata.Group{ObjectID: common.BKInnerObjIDHost, GroupID: mCommon.MoreInfo, GroupName: mCommon.MoreInfoName, GroupIndex: 3, OwnerID: ownerID, IsDefault: true},
		//storage
		&metadata.Group{ObjectID: common.BKInnerObjIDStorage, GroupID: mCommon.BaseInfo, GroupName: mCommon.BaseInfoName, GroupIndex: 1, OwnerID: ownerID, IsDefault: true},
		//idc
		&metadata.Group{ObjectID: common.BKInnerObjIDIdc, GroupID: mCommon.BaseInfo, GroupName: mCommon.BaseInfoName, GroupIndex: 1, OwnerID: ownerID, IsDefault: true},
		//idcrack
		&metadata.Group{ObjectID: common.BKInnerObjIDIdcRack, GroupID: mCommon.BaseInfo, GroupName: mCommon.BaseInfoName, GroupIndex: 1, OwnerID: ownerID, IsDefault: true},
		//router
		&metadata.Group{ObjectID: common.BKInnerObjIDRouter, GroupID: mCommon.BaseInfo, GroupName: mCommon.BaseInfoName, GroupIndex: 1, OwnerID: ownerID, IsDefault: true},
		//switch
		&metadata.Group{ObjectID: common.BKInnerObjIDSwitch, GroupID: mCommon.BaseInfo, GroupName: mCommon.BaseInfoName, GroupIndex: 1, OwnerID: ownerID, IsDefault: true},
		//firewall
		&metadata.Group{ObjectID: common.BKInnerObjIDFirewall, GroupID: mCommon.BaseInfo, GroupName: mCommon.BaseInfoName, GroupIndex: 1, OwnerID: ownerID, IsDefault: true},
		//bk_blance
		&metadata.Group{ObjectID: common.BKInnerObjIDBlance, GroupID: mCommon.BaseInfo, GroupName: mCommon.BaseInfoName, GroupIndex: 1, OwnerID: ownerID, IsDefault: true},

		//proc
		&metadata.Group{ObjectID: common.BKInnerObjIDProc, GroupID: mCommon.BaseInfo, GroupName: mCommon.BaseInfoName, GroupIndex: 1, OwnerID: ownerID, IsDefault: true},
		&metadata.Group{ObjectID: common.BKInnerObjIDProc, GroupID: mCommon.ProcPort, GroupName: mCommon.ProcPortName, GroupIndex: 2, OwnerID: ownerID, IsDefault: true},
		&metadata.Group{ObjectID: common.BKInnerObjIDProc, GroupID: mCommon.ProcGsekitBaseInfo, GroupName: mCommon.ProcGsekitBaseInfoName, GroupIndex: 3, OwnerID: ownerID, IsDefault: true},
		&metadata.Group{ObjectID: common.BKInnerObjIDProc, GroupID: mCommon.ProcGsekitManageInfo, GroupName: mCommon.ProcGsekitManageInfoName, GroupIndex: 4, OwnerID: ownerID, IsDefault: true},

		//plat
		&metadata.Group{ObjectID: common.BKInnerObjIDPlat, GroupID: mCommon.BaseInfo, GroupName: mCommon.BaseInfoName, GroupIndex: 1, OwnerID: ownerID, IsDefault: true},

		//nginx_service
		&metadata.Group{ObjectID: common.BKInnerObjIDNginx, GroupID: mCommon.BaseInfo, GroupName: mCommon.BaseInfoName, GroupIndex: 1, OwnerID: ownerID, IsDefault: true},
		//zookeeper_service
		&metadata.Group{ObjectID: common.BKInnerObjIDZookeeper, GroupID: mCommon.BaseInfo, GroupName: mCommon.BaseInfoName, GroupIndex: 1, OwnerID: ownerID, IsDefault: true},
		//mysql_service
		&metadata.Group{ObjectID: common.BKInnerObjIDMysql, GroupID: mCommon.BaseInfo, GroupName: mCommon.BaseInfoName, GroupIndex: 1, OwnerID: ownerID, IsDefault: true},
		//mongo_service
		&metadata.Group{ObjectID: common.BKInnerObjIDMongo, GroupID: mCommon.BaseInfo, GroupName: mCommon.BaseInfoName, GroupIndex: 1, OwnerID: ownerID, IsDefault: true},
		//weblogic_service
		&metadata.Group{ObjectID: common.BKInnerObjIDWeblogic, GroupID: mCommon.BaseInfo, GroupName: mCommon.BaseInfoName, GroupIndex: 1, OwnerID: ownerID, IsDefault: true},
		//tomcat_service
		&metadata.Group{ObjectID: common.BKInnerObjIDTomcat, GroupID: mCommon.BaseInfo, GroupName: mCommon.BaseInfoName, GroupIndex: 1, OwnerID: ownerID, IsDefault: true},
		//apache_service
		&metadata.Group{ObjectID: common.BKInnerObjIDApache, GroupID: mCommon.BaseInfo, GroupName: mCommon.BaseInfoName, GroupIndex: 1, OwnerID: ownerID, IsDefault: true},

		//user
		&metadata.Group{ObjectID: common.BKInnerObjIDUser, GroupID: mCommon.BaseInfo, GroupName: mCommon.BaseInfoName, GroupIndex: 1, OwnerID: ownerID, IsDefault: true},
		//usergroup
		&metadata.Group{ObjectID: common.BKInnerObjIDUserGroup, GroupID: mCommon.BaseInfo, GroupName: mCommon.BaseInfoName, GroupIndex: 1, OwnerID: ownerID, IsDefault: true},
	}
	for objID, kv := range objectIDs {
		index := int64(1)
		for id, name := range kv {
			row := &metadata.Group{ObjectID: objID, GroupID: id, GroupName: name, GroupIndex: index, OwnerID: ownerID, IsDefault: true}
			dataRows = append(dataRows, row)
			index++
		}

	}

	return dataRows

}

var classificationRows = []*metadata.Classification{
	&metadata.Classification{ClassificationID: "bk_iaas", ClassificationName: "基础设施", ClassificationType: "inner", ClassificationIcon: "icon-cc-iaas"},
	&metadata.Classification{ClassificationID: "bk_paas", ClassificationName: "平台资源", ClassificationType: "inner", ClassificationIcon: "icon-cc-paas"},
	&metadata.Classification{ClassificationID: "bk_saas", ClassificationName: "应用资源", ClassificationType: "inner", ClassificationIcon: "icon-cc-saas"},
	&metadata.Classification{ClassificationID: "bk_organization", ClassificationName: "组织信息", ClassificationType: "inner", ClassificationIcon: "icon-cc-organization"},
}


func getModeTopoData(ownerID string) []*metadata.TopoGraphics {

	create := func(num int64) *int64 {
		return &num;
	}

	dataRows := []*metadata.TopoGraphics{
		//bk_iaas（基础设施）
		&metadata.TopoGraphics{ObjID: common.BKInnerObjIDHost,Position: metadata.Position{X:create(23),Y:create(46)}},
		&metadata.TopoGraphics{ObjID: common.BKInnerObjIDStorage,Position:metadata.Position{X:create(-280),Y:create(435)}},
		&metadata.TopoGraphics{ObjID: common.BKInnerObjIDIdc,Position:metadata.Position{X:create(38),Y:create(423)}},
		&metadata.TopoGraphics{ObjID: common.BKInnerObjIDIdcRack,Position:metadata.Position{X:create(27),Y:create(228)}},
		&metadata.TopoGraphics{ObjID: common.BKInnerObjIDRouter,Position:metadata.Position{X:create(343),Y:create(-61)}},
		&metadata.TopoGraphics{ObjID: common.BKInnerObjIDSwitch,Position:metadata.Position{X:create(349),Y:create(406)}},
		&metadata.TopoGraphics{ObjID: common.BKInnerObjIDFirewall,Position:metadata.Position{X:create(342),Y:create(236)}},
		&metadata.TopoGraphics{ObjID: common.BKInnerObjIDBlance,Position:metadata.Position{X:create(345),Y:create(98)}},
		//bk_saas（应用资源）
		&metadata.TopoGraphics{ObjID: common.BKInnerObjIDSet,Position:metadata.Position{X:create(-304),Y:create(-143)}},
		&metadata.TopoGraphics{ObjID: common.BKInnerObjIDModule,Position:metadata.Position{X:create(-296),Y:create(45)}},
		&metadata.TopoGraphics{ObjID: common.BKInnerObjIDPlat,Position:metadata.Position{X:create(23),Y:create(-144)}},
		&metadata.TopoGraphics{ObjID: common.BKInnerObjIDApp,Position:metadata.Position{X:create(-637),Y:create(-145)}},

		//bk_organization (组织信息)
		&metadata.TopoGraphics{ObjID: common.BKInnerObjIDUserGroup,Position:metadata.Position{X:create(-630),Y:create(243)}},
		&metadata.TopoGraphics{ObjID: common.BKInnerObjIDUser,Position:metadata.Position{X:create(-300),Y:create(238)}},

	}

	for _, r := range dataRows {
		r.ScopeType = "global"
		r.ScopeID = "0"
		r.NodeType = "obj"
		r.IsPre = false
		r.InstID = 0
		r.SupplierAccount = ownerID
		r.Assts = []metadata.GraphAsst{}
	}

	return dataRows
}