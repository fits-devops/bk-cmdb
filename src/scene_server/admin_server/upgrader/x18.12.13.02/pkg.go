package x18_12_13_02

import (
	"context"

	"configcenter/src/common/blog"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func init() {
	upgrader.RegistUpgrader("x18_12_13_02", upgrade)
}

func upgrade(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {

	err = addCloudTaskTable(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x18_12_13_02] addCloudTaskTable error  %s", err.Error())
		return err
	}

	err = addCloudResourceConfirmTable(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x18_12_13_02] addCloudResourceConfirmTable error  %s", err.Error())
		return err
	}

	err = addCloudSyncHistoryTable(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x18_12_13_02] addCloudSyncHistoryTable error  %s", err.Error())
		return err
	}

	err = addCloudConfirmHistoryTable(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade x18_12_13_02] addCloudConfirmHistoryTable error  %s", err.Error())
		return err
	}

	return
}
