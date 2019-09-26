package x19_04_16_02

import (
	"context"

	"configcenter/src/common/blog"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func init() {
	upgrader.RegistUpgrader("x19.04.16.02", upgrade)
}

func upgrade(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	err = updateFranceName(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x19.04.16.02] fixAssociationTypeName error  %s", err.Error())
		return err
	}
	return nil
}
