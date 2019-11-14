package x18_09_13_01

import (
	"context"

	"configcenter/src/common"
	"configcenter/src/common/mapstr"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func reconcileGroupPrivilege(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	all := []mapstr.MapStr{}

	if err = db.Table(common.BKTableNameUserGroupPrivilege).Find(nil).All(ctx, &all); nil != err {
		return err
	}
	flag := "updateflag"
	expectM := map[string]mapstr.MapStr{}
	for _, privilege := range all {
		groupID, err := privilege.String("group_id")
		if err != nil {
			return err
		}
		privilege.Set(flag, true)
		expectM[groupID] = privilege
	}

	for _, privilege := range expectM {
		if err = db.Table(common.BKTableNameUserGroupPrivilege).Insert(ctx, privilege); nil != err {
			return err
		}
	}

	if err = db.Table(common.BKTableNameUserGroupPrivilege).Delete(ctx, map[string]interface{}{
		flag: map[string]interface{}{
			"$ne": true,
		},
	}); err != nil {
		return err
	}

	if err = db.Table(common.BKTableNameUserGroupPrivilege).DropColumn(ctx, flag); err != nil {
		return err
	}

	return nil

}
