package x19_01_18_01

import (
	"context"

	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func init() {
	upgrader.RegistUpgrader("x19.01.18.01", upgrade)
}

func upgrade(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	err = dropIndex(ctx, db, conf)
	if err != nil {
		return err
	}

	return
}
