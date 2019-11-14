package x18_09_11_01

import (
	"context"

	"configcenter/src/common/blog"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func init() {
	upgrader.RegistUpgrader("x18.09.11.01", upgrade)
}

func upgrade(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	err = addOperationLogIndex(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x18.09.11.01] updateSystemProperty error  %s", err.Error())
		return err
	}
	return
}
