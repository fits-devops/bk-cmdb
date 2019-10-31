package x19_10_31_01

import (
	"context"

	"configcenter/src/common/blog"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func init() {
	upgrader.RegistUpgrader("x19.10.31.01", upgrade)
}
func upgrade(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	blog.Infof("from now on, the cmdb version will be v4.0.x")
	err = addAssociation(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x19.10.31.01] addAssociation error  %s", err.Error())
		return err
	}

	return
}
