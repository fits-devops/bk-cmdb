package v3v0v9beta1

import (
	"context"

	"configcenter/src/common"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func fixesSupplierAccount(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	for _, tablename := range shouldAddSupplierAccountFieldTables {
		condition := map[string]interface{}{
			common.BKOwnerIDField: map[string]interface{}{
				"$in": []interface{}{nil, ""},
			},
		}
		data := map[string]interface{}{
			common.BKOwnerIDField: common.BKDefaultOwnerID,
		}
		err := db.Table(tablename).Update(ctx, condition, data)
		if nil != err {
			return err
		}
	}
	return nil
}

var shouldAddSupplierAccountFieldTables = []string{
	"cc_ApplicationBase",
	"cc_HostBase",
	"cc_ModuleBase",
	"cc_ModuleHostConfig",
	"cc_ObjAsst",
	"cc_ObjAttDes",
	"cc_ObjClassification",
	"cc_ObjDes",
	"cc_ObjectBase",
	"cc_OperationLog",
	"cc_PlatBase",
	"cc_Proc2Module",
	"cc_Process",
	"cc_PropertyGroup",
	"cc_SetBase",
	"cc_Subscription",
	"cc_TopoGraphics",
	"cc_InstAsst",
	"cc_Privilege",
	"cc_History",
	"cc_HostFavourite",
	"cc_UserAPI",
	"cc_UserCustom",
	"cc_UserGroup",
	"cc_UserGroupPrivilege",
}
