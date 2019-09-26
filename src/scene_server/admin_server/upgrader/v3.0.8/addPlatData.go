package v3v0v8

import (
	"context"
	"time"

	"configcenter/src/common"
	"configcenter/src/common/blog"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func addPlatData(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {
	tablename := "cc_PlatBase"
	blog.Errorf("add data for  %s table ", tablename)
	rows := []map[string]interface{}{
		map[string]interface{}{
			common.BKCloudNameField: "default area",
			common.BKOwnerIDField:   common.BKDefaultOwnerID,
			common.BKCloudIDField:   common.BKDefaultDirSubArea,
			common.CreateTimeField:  time.Now(),
			common.LastTimeField:    time.Now(),
		},
	}
	for _, row := range rows {
		// ensure id plug > 1, 1Reserved
		platID, err := db.NextSequence(ctx, tablename)
		if err != nil {
			return err
		}
		// Direct connecting area id = 1
		if common.BKDefaultDirSubArea == row[common.BKCloudIDField] {
			platID = common.BKDefaultDirSubArea
		}

		row[common.BKCloudIDField] = platID
		_, _, err = upgrader.Upsert(ctx, db, tablename, row, "", []string{common.BKCloudNameField, common.BKOwnerIDField}, []string{common.BKCloudIDField})
		if nil != err {
			blog.Errorf("add data for  %s table error  %s", tablename, err)
			return err
		}

		return nil

	}

	blog.Errorf("add data for  %s table  ", tablename)
	return nil
}
