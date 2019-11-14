package x18_09_18_01

import (
	"context"

	"configcenter/src/common/blog"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func init() {
	upgrader.RegistUpgrader("x18.09.18.01", upgrade)
}

func upgrade(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	err = fixedHostPlatAssocateRelation(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x18.09.18.01] fixedHostPlatAssocateRelation error  %s", err.Error())
		return err
	}

	return nil
}
