package x18_12_12_02

import (
	"context"

	"configcenter/src/common"
	"configcenter/src/common/condition"
	"configcenter/src/common/mapstr"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func fixModuleNamePropertyGroup(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {

	cond := condition.CreateCondition()
	cond.Field(common.BKOwnerIDField).Eq(common.BKDefaultOwnerID)
	cond.Field(common.BKObjIDField).Eq(common.BKInnerObjIDModule)
	cond.Field(common.BKPropertyIDField).Eq(common.BKModuleNameField)

	data := mapstr.MapStr{
		common.BKPropertyGroupField: "default",
	}

	err := db.Table(common.BKTableNameObjAttDes).Update(ctx, cond.ToMapStr(), data)
	if err != nil {
		return err
	}
	return nil
}
