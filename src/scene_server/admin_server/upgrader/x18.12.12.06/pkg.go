package x18_12_12_06

import (
	"context"

	"configcenter/src/common/blog"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func init() {
	upgrader.RegistUpgrader("x18.12.12.06", upgrade)
}

func upgrade(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	err = reconcilUnique(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x18.12.12.06] fixBKObjAsstID error  %s", err.Error())
		return err
	}
	return
}
