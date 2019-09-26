package x08_09_11_01

import (
	"context"

	"configcenter/src/common"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func addOperationLogIndex(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {

	index := dal.Index{Name: "", Keys: map[string]int32{"opt_time": 1}, Background: true}

	if err = db.Table(common.BKTableNameOperationLog).CreateIndex(ctx, index); err != nil && db.IsDuplicatedError(err) {
		return err
	}

	return nil
}
