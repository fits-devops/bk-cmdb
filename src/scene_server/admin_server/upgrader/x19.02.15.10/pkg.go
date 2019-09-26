package x19_02_15_10

import (
	"context"

	"configcenter/src/common/blog"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func init() {
	upgrader.RegistUpgrader("x19.02.15.10", upgrade)
}

func upgrade(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	err = fixAssociationTypeName(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x19.02.15.10] fixAssociationTypeName error  %s", err.Error())
		return err
	}
	err = fixEventSubscribeLastTime(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x19.02.15.10] fixEventSubscribeLastTime error  %s", err.Error())
		return err
	}
	return
}
