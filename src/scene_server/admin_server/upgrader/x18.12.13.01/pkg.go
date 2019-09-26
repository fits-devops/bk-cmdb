package x18_12_13_01

import (
	"context"

	"configcenter/src/common/blog"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func init() {
	upgrader.RegistUpgrader("x18.12.13.01", upgrade)
}
func upgrade(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	blog.Infof("from now on, the cmdb version will be v3.4.x")
	err = addswitchAssociation(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x18.12.13.01] addswitchAssociation error  %s", err.Error())
		return err
	}
	err = changeNetDeviceTableName(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x18.12.13.01] changeNetDeviceTableName error  %s", err.Error())
		return err
	}

	return
}
