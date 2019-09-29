package x19_09_29_01

import (
	"context"

	"configcenter/src/common/blog"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func init() {
	upgrader.RegistUpgrader("x19.09.29.01", upgrade)
}
func upgrade(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	blog.Infof("from now on, the cmdb version will be v3.4.x")
	err = UpdateIaasAssociation(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x19.09.29.01] UpdateIaasAssociation error  %s", err.Error())
		return err
	}

	return
}
