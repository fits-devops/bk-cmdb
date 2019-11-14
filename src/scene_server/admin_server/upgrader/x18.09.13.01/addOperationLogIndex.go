package x18_09_13_01

import (
	"context"

	"configcenter/src/common"
	"configcenter/src/common/blog"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func addOperationLogIndex(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	indexs, err := db.Table(common.BKTableNameOperationLog).Indexes(ctx)
	if err != nil {
		return err
	}

	existIndexMap := map[string]bool{}
	for _, index := range indexs {
		if index.Name == "_id_" {
			continue
		}
		existIndexMap[index.Name] = true
	}
	blog.Infof("existing index %v", existIndexMap)

	expectIndexs := []dal.Index{
		{Name: "op_target_1_inst_id_1_op_time_-1", Keys: map[string]int32{"op_target": 1, "inst_id": 1, "op_time": -1}, Background: true},
		{Name: "org_id_1_op_time_-1", Keys: map[string]int32{"org_id": 1, "op_time": -1}, Background: true},
		{Name: "biz_id_1_org_id_1_op_time_-1", Keys: map[string]int32{"biz_id": 1, "org_id": 1, "op_time": -1}, Background: true},
		{Name: "ext_key_1_org_id_1_op_time_-1", Keys: map[string]int32{"ext_key": 1, "org_id": 1, "op_time": -1}, Background: true},
	}
	for _, idx := range expectIndexs {
		blog.Infof("creating index %s", idx.Name)
		if !existIndexMap[idx.Name] {
			if err = db.Table(common.BKTableNameOperationLog).CreateIndex(ctx, idx); err != nil {
				blog.Infof("creat index %s failed, %v", idx.Name, err)
				return err
			}
		}
		existIndexMap[idx.Name] = false
		blog.Infof("creat index %s success", idx.Name)
	}

	for idxname, shouldDelete := range existIndexMap {
		if shouldDelete {
			blog.Infof("droping index %s", idxname)
			if err = db.Table(common.BKTableNameOperationLog).DropIndex(ctx, idxname); err != nil {
				blog.Infof("drop index %s failed, %v", idxname, err)
				return err
			}
		}
	}
	return nil
}
