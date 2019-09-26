package v3v0v9beta3

import (
	"context"

	"configcenter/src/common"
	"configcenter/src/common/blog"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/scene_server/validator"
	"configcenter/src/storage/dal"
)

func fixesProcessPortPattern(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	condition := map[string]interface{}{
		common.BKObjIDField:      common.BKInnerObjIDProc,
		common.BKPropertyIDField: "port",
	}
	data := map[string]interface{}{
		"option": common.PatternMultiplePortRange,
	}
	err = db.Table(common.BKTableNameObjAttDes).Update(ctx, condition, data)
	if nil != err {
		blog.Errorf("[upgrade v3.0.9-beta.3] fixesPortPattern error  %s", err.Error())
		return err
	}
	return nil
}

func fixesProcessPriorityPattern(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	condition := map[string]interface{}{
		common.BKObjIDField:      common.BKInnerObjIDProc,
		common.BKPropertyIDField: "priority",
	}
	data := map[string]interface{}{
		"option": validator.MinMaxOption{Min: "1", Max: "10000"},
	}
	err = db.Table(common.BKTableNameObjAttDes).Update(ctx, condition, data)
	if nil != err {
		blog.Errorf("[upgrade v3.0.9-beta.3] fixesPortPattern error  %s", err.Error())
		return err
	}
	return nil
}
