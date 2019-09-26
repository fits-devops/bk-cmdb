package x18_10_10_01

import (
	"context"

	"configcenter/src/common/blog"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func init() {
	upgrader.RegistUpgrader("x18.10.10.01", upgrade)
}
func upgrade(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	err = addProcOpTaskTable(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x18.10.10.01] addProcOpTaskTable error  %s", err.Error())
		return err
	}
	err = addProcInstanceModelTable(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x18.10.10.01] addProcInstanceModelTable error  %s", err.Error())
		return err
	}
	err = addProcInstanceDetailTable(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x18.10.10.01] addProcInstanceDetailTable error  %s", err.Error())
		return err
	}
	err = addProcFreshInstance(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x18.10.10.01] addProcFreshInstance error  %s", err.Error())
		return err
	}
	return
}
