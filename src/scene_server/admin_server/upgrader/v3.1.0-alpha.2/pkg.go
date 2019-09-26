package v3v0v1alpha2

import (
	"context"

	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func init() {
	upgrader.RegistUpgrader("v3.1.0-alpha.2", upgrade)
}

func upgrade(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	err = addBkStartParamRegex(ctx, db, conf)
	if err != nil {
		return err
	}
	err = updateLanguageField(ctx, db, conf)
	if err != nil {
		return err
	}
	return
}
