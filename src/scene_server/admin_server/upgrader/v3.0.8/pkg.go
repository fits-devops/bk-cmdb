package v3v0v8

import (
	"context"

	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func init() {
	upgrader.RegistUpgrader("v3.0.8", upgrade)
}

func upgrade(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	//创建表
	err = createTable(ctx, db, conf)
	if err != nil {
		return err
	}
	//增加内置对象数据
	err = addPresetObjects(ctx, db, conf)
	if err != nil {
		return err
	}
	//增加云区域数据
	err = addPlatData(ctx, db, conf)
	if err != nil {
		return err
	}
	//增加系统数据
	err = addSystemData(ctx, db, conf)
	if err != nil {
		return err
	}
	//增加默认的应用数据
	//err = addDefaultBiz(ctx, db, conf)
	//if err != nil {
	//	return err
	//}
	//增加后台数据
	err = addBKApp(ctx, db, conf)
	if err != nil {
		return err
	}

	return
}
