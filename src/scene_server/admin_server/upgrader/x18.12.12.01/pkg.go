package x18_12_12_01

import (
	"context"

	"configcenter/src/common/blog"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func init() {
	upgrader.RegistUpgrader("x18.12.12.01", upgrade)
}

func upgrade(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	err = addAIXProperty(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x18.12.12.01] addAIXProperty error  %s", err.Error())
		return err
	}
	err = setTCPDefault(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x18.12.12.01] setTCPDefault error  %s", err.Error())
		return err
	}
	err = reconcilUnique(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x18.12.12.01] reconcilUnique error  %s", err.Error())
		return err
	}

	return
}
