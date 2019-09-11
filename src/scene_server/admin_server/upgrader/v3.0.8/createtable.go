package v3v0v8

import (
	"context"

	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"

	"gopkg.in/mgo.v2"
)

func createTable(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	for tablename, indexs := range tables {
		exists, err := db.HasTable(tablename)
		if err != nil {
			return err
		}
		if !exists {
			if err = db.CreateTable(tablename); err != nil && !mgo.IsDup(err) {
				return err
			}
		}
		for index := range indexs {
			if err = db.Table(tablename).CreateIndex(ctx, indexs[index]); err != nil && !db.IsDuplicatedError(err) {
				return err
			}
		}
	}
	return nil
}

var tables = map[string][]dal.Index{
	"cc_ApplicationBase": []dal.Index{
		dal.Index{Name: "", Keys: map[string]int32{"biz_id": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"biz_name": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"default": 1}, Background: true},
	},

	"cc_HostBase": []dal.Index{
		dal.Index{Name: "", Keys: map[string]int32{"host_id": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"host_name": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"host_innerip": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"host_outerip": 1}, Background: true},
	},
	"cc_ModuleBase": []dal.Index{
		dal.Index{Name: "", Keys: map[string]int32{"module_id": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"module_name": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"default": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"biz_id": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"org_id": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"set_id": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"bk_parent_id": 1}, Background: true},
	},
	"cc_ModuleHostConfig": []dal.Index{
		dal.Index{Name: "", Keys: map[string]int32{"biz_id": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"host_id": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"module_id": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"set_id": 1}, Background: true},
	},
	"cc_ObjAsst": []dal.Index{
		dal.Index{Name: "", Keys: map[string]int32{"obj_id": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"bk_asst_obj_id": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"org_id": 1}, Background: true},
	},
	"cc_ObjAttDes": []dal.Index{
		dal.Index{Name: "", Keys: map[string]int32{"obj_id": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"org_id": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"id": 1}, Background: true},
	},
	"cc_ObjClassification": []dal.Index{
		dal.Index{Name: "", Keys: map[string]int32{"classification_id": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"classification_name": 1}, Background: true},
	},
	"cc_ObjDes": []dal.Index{
		dal.Index{Name: "", Keys: map[string]int32{"obj_id": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"classification_id": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"obj_name": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"org_id": 1}, Background: true},
	},
	"cc_ObjectBase": []dal.Index{
		dal.Index{Name: "", Keys: map[string]int32{"obj_id": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"org_id": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"inst_id": 1}, Background: true},
	},
	"cc_OperationLog": []dal.Index{
		dal.Index{Name: "", Keys: map[string]int32{"op_target": 1, "inst_id": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"org_id": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"biz_id": 1, "org_id": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"ext_key": 1, "org_id": 1}, Background: true},
	},
	"cc_PlatBase": []dal.Index{
		dal.Index{Name: "", Keys: map[string]int32{"org_id": 1}, Background: true},
	},
	"cc_Proc2Module": []dal.Index{
		dal.Index{Name: "", Keys: map[string]int32{"biz_id": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"process_id": 1}, Background: true},
	},
	"cc_Process": []dal.Index{
		dal.Index{Name: "", Keys: map[string]int32{"process_id": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"biz_id": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"org_id": 1}, Background: true},
	},
	"cc_PropertyGroup": []dal.Index{
		dal.Index{Name: "", Keys: map[string]int32{"obj_id": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"org_id": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"group_id": 1}, Background: true},
	},
	"cc_SetBase": []dal.Index{
		dal.Index{Name: "", Keys: map[string]int32{"set_id": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"bk_parent_id": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"biz_id": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"org_id": 1}, Background: true},
		dal.Index{Name: "", Keys: map[string]int32{"set_name": 1}, Background: true},
	},
	"cc_Subscription": []dal.Index{
		dal.Index{Name: "", Keys: map[string]int32{"subscription_id": 1}, Background: true},
	},
	"cc_TopoGraphics": []dal.Index{
		dal.Index{Name: "", Keys: map[string]int32{"scope_type": 1, "scope_id": 1, "node_type": 1, "obj_id": 1, "inst_id": 1}, Background: true, Unique: true},
	},
	"cc_InstAsst": []dal.Index{
		dal.Index{Name: "", Keys: map[string]int32{"obj_id": 1, "inst_id": 1}, Background: true},
	},

	"cc_Privilege":          []dal.Index{},
	"cc_History":            []dal.Index{},
	"cc_HostFavourite":      []dal.Index{},
	"cc_UserAPI":            []dal.Index{},
	"cc_UserCustom":         []dal.Index{},
	"cc_UserGroup":          []dal.Index{},
	"cc_UserGroupPrivilege": []dal.Index{},
	"cc_idgenerator":        []dal.Index{},
	"cc_System":             []dal.Index{},
}
