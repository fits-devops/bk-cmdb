package x08_09_17_01

import (
	"context"

	"configcenter/src/common"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func createTable(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	for tablename, indexs := range tables {
		exists, err := db.HasTable(tablename)
		if err != nil {
			return err
		}
		if !exists {
			if err = db.CreateTable(tablename); err != nil && !db.IsDuplicatedError(err) {
				return err
			}
		}
		for index := range indexs {
			if err = db.Table(tablename).CreateIndex(ctx, indexs[index]); err != nil && !db.IsDuplicatedError(err) {
				return err
			}
		}
	}
	return nil
}

var tables = map[string][]dal.Index{
	common.BKTableNameNetcollectDevice: []dal.Index{
		{Keys: map[string]int32{"device_id": 1}, Background: true},
		{Keys: map[string]int32{"device_name": 1}, Background: true},
		{Keys: map[string]int32{"org_id": 1}, Background: true},
	},

	common.BKTableNameNetcollectProperty: []dal.Index{
		{Keys: map[string]int32{"netcollect_property_id": 1}, Background: true},
		{Keys: map[string]int32{"org_id": 1}, Background: true},
	},
}
