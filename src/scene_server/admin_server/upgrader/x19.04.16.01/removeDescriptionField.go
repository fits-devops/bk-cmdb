package x19_04_16_01

import (
	"context"

	"configcenter/src/common"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func removeDescriptionField(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {
	err := db.Table(common.BKTableNameObjAttDes).DropColumn(ctx, "description")
	if err != nil {
		return err
	}
	return nil
}
