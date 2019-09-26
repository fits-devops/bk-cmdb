package x18_09_30_01

import (
	"context"

	"configcenter/src/common/blog"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func init() {
	upgrader.RegistUpgrader("x18.09.30.01", upgrade)
}

func upgrade(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	err = cleanBKCloud(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x18.09.30.01] cleanBKCloud error  %s", err.Error())
		return err
	}

	return nil
}
