package x18_09_04_01

import (
	"context"

	"configcenter/src/common"
	"configcenter/src/common/metadata"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/scene_server/validator"
	"configcenter/src/storage/dal"
)

func updateSystemProperty(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	objs := []metadata.Object{}
	condition := map[string]interface{}{
		"classification_id": "bk_biz_topo",
	}
	err = db.Table(common.BKTableNameObjDes).Find(condition).All(ctx, &objs)
	if err != nil {
		return err
	}

	objIDs := []string{}
	for _, obj := range objs {
		objIDs = append(objIDs, obj.ObjectID)
	}

	tablename := common.BKTableNameObjAttDes
	condition = map[string]interface{}{
		"property_id": map[string]interface{}{"$in": []string{common.BKChildStr, common.BKInstParentStr}},
		"obj_id":      map[string]interface{}{"$in": objIDs},
	}
	data := map[string]interface{}{
		"issystem": true,
	}

	return db.Table(tablename).Update(ctx, condition, data)
}

func updateIcon(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	condition := map[string]interface{}{
		"obj_id": common.BKInnerObjIDTomcat,
	}
	data := map[string]interface{}{
		"obj_icon": "icon-cc-tomcat",
	}
	err = db.Table(common.BKTableNameObjDes).Update(ctx, condition, data)
	if err != nil {
		return err
	}
	condition = map[string]interface{}{
		"obj_id": common.BKInnerObjIDApache,
	}
	data = map[string]interface{}{
		"obj_icon": "icon-cc-apache",
	}

	return db.Table(common.BKTableNameObjDes).Update(ctx, condition, data)
}

func fixesProcess(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	condition := map[string]interface{}{
		common.BKObjIDField:      common.BKInnerObjIDProc,
		common.BKPropertyIDField: map[string]interface{}{"$in": []string{"priority", "proc_num", "auto_time_gap", "timeout"}},
	}
	data := map[string]interface{}{
		"option": validator.MinMaxOption{Min: "1", Max: "10000"},
	}
	err = db.Table(common.BKTableNameObjAttDes).Update(ctx, condition, data)
	if nil != err {
		return err
	}
	return nil
}
