package x18_10_30_02

import (
	"context"

	"configcenter/src/common/blog"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func init() {
	upgrader.RegistUpgrader("x18.10.30.02", upgrade)
}
func upgrade(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	err = addBizSuupierID(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x18.10.30.02] addBizSuupierID error  %s", err.Error())
		return err
	}

	return
}
