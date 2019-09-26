package v3v0v9beta1

import (
	"context"

	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func init() {
	upgrader.RegistUpgrader("v3.0.9-beta.1", upgrade)
}

func upgrade(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	err = fixesSupplierAccount(ctx, db, conf)
	if err != nil {
		return err
	}

	return
}
