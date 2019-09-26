package x08_09_17_01

import (
	"context"

	"configcenter/src/common/blog"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func init() {
	upgrader.RegistUpgrader("x08.09.17.01", upgrade)
}

func upgrade(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	err = createTable(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x08.09.17.01] create table netcollect error  %s", err.Error())
		return err
	}

	return nil
}
