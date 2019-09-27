package x19_04_16_02

import (
	"context"

	"configcenter/src/common"
	"configcenter/src/common/condition"
	"configcenter/src/common/mapstr"
	"configcenter/src/common/metadata"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/source_controller/coreservice/core/instances"
	"configcenter/src/storage/dal"
)

func updateFranceName(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {
	cond := condition.CreateCondition()
	cond.Field(common.BKObjIDField).Eq(common.BKInnerObjIDHost)
	cond.Field(common.BKPropertyIDField).Eq("state_name")
	state := metadata.Attribute{}
	err := db.Table(common.BKTableNameObjAttDes).Find(cond.ToMapStr()).One(ctx, &state)
	if err != nil {
		return err
	}

	enums, err := instances.ParseEnumOption(state.Option)
	if err != nil {
		return err
	}
	for index := range enums {
		if enums[index].ID == "FR" {
			enums[index].Name = "法国"
		}
	}

	return db.Table(common.BKTableNameObjAttDes).Update(ctx, cond.ToMapStr(), mapstr.MapStr{common.BKOptionField: enums})
}
