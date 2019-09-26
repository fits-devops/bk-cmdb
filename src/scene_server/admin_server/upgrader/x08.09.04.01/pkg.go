package x08_09_04_01

import (
	"context"

	"configcenter/src/common/blog"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func init() {
	upgrader.RegistUpgrader("x08.09.04.01", upgrade)
}

func upgrade(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	err = updateSystemProperty(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x08.09.04.01] updateSystemProperty error  %s", err.Error())
		return err
	}
	err = updateIcon(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x08.09.04.01] updateIcon error  %s", err.Error())
		return err
	}
	err = fixesProcess(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x08.09.04.01] fixesProcess error  %s", err.Error())
		return err
	}
	return
}
