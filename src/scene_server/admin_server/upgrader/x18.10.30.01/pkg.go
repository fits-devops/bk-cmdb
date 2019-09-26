package x18_10_30_01

import (
	"context"

	"configcenter/src/common/blog"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func init() {
	upgrader.RegistUpgrader("x18.10.30.01", upgrade)
}
func upgrade(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	err = createAssociationTable(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x18.10.30.01] createAssociationTable error  %s", err.Error())
		return err
	}
	err = addPresetAssociationType(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x18.10.30.01] addPresetAssociationType error  %s", err.Error())
		return err
	}
	err = reconcilAsstData(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x18.10.30.01] reconcilAsstData error  %s", err.Error())
		return err
	}
	return
}
