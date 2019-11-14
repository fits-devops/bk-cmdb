package x18_09_13_01

import (
	"context"

	"configcenter/src/common/blog"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func init() {
	upgrader.RegistUpgrader("x18.09.13.01", upgrade)
}

func upgrade(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	err = addOperationLogIndex(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x18.09.13.01] addOperationLogIndex error  %s", err.Error())
		return err
	}
	err = reconcileGroupPrivilege(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x18.09.13.01] reconcileGroupPrivilege error  %s", err.Error())
		return err
	}
	err = reconcileGroupPrivilege(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x18.09.13.01] reconcileGroupPrivilege error  %s", err.Error())
		return err
	}
	return
}
