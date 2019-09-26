package x18_12_12_02

import (
	"context"

	"configcenter/src/common/blog"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func init() {
	upgrader.RegistUpgrader("x18.12.12.02", upgrade)
}

func upgrade(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	err = fixModuleNamePropertyGroup(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x18.12.12.02] addAIXProperty error  %s", err.Error())
		return err
	}
	return
}
