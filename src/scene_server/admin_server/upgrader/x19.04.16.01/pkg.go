package x19_04_16_01

import (
	"context"

	"configcenter/src/common/blog"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func init() {
	upgrader.RegistUpgrader("x19.04.16.01", upgrade)
}

func upgrade(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	err = removeDescriptionField(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x19.04.16.01] fixAssociationTypeName error  %s", err.Error())
		return err
	}
	return nil
}
