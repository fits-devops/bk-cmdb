package v3v0v8

import (
	"context"

	"configcenter/src/common"
	"configcenter/src/common/blog"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func addSystemData(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {
	tablename := "cc_System"
	blog.V(3).Infof("add data for  %s table ", tablename)
	data := map[string]interface{}{
		common.HostCrossBizField: common.HostCrossBizValue}
	isExist, err := db.Table(tablename).Find(data).Count(ctx)
	if nil != err {
		blog.Errorf("add data for  %s table error  %s", tablename, err)
		return err
	}
	if isExist > 0 {
		return nil
	}
	err = db.Table(tablename).Insert(ctx, data)
	if nil != err {
		blog.Errorf("add data for  %s table error  %s", tablename, err)
		return err
	}

	return nil
}
