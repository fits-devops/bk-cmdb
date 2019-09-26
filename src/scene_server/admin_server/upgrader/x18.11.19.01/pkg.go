package x18_11_19_01

import (
	"context"

	"configcenter/src/common/blog"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func init() {
	upgrader.RegistUpgrader("x18.11.19.01", upgrade)
}
func upgrade(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	err = createObjectUnitTable(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x18.11.19.01] createObjectUnitTable error  %s", err.Error())
		return err
	}

	err = reconcilUnique(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x18.11.19.01] reconcilUnique error  %s", err.Error())
		return err
	}
	// 产品调整，回撤
	// err = reconcilAsstID(ctx, db, conf)
	// if err != nil {
	// 	blog.Errorf("[upgrade x18.11.19.01] reconcilAsstID error  %s", err.Error())
	// 	return err
	// }
	return
}
