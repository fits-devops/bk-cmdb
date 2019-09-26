package x18_12_12_01

import (
	"context"

	"configcenter/src/common"
	"configcenter/src/common/condition"
	"configcenter/src/common/mapstr"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/scene_server/validator"
	"configcenter/src/storage/dal"
)

func setTCPDefault(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {

	cond := condition.CreateCondition()
	cond.Field(common.BKOwnerIDField).Eq(common.BKDefaultOwnerID)
	cond.Field(common.BKPropertyIDField).Eq(common.BKProtocol)

	ostypeProperty := Attribute{}
	err := db.Table(common.BKTableNameObjAttDes).Find(cond.ToMapStr()).One(ctx, &ostypeProperty)
	if err != nil {
		return err
	}

	enumOpts := validator.ParseEnumOption(ostypeProperty.Option)
	for index := range enumOpts {
		if enumOpts[index].Name == "TCP" {
			enumOpts[index].IsDefault = true
		}
	}

	data := mapstr.MapStr{
		common.BKOptionField: enumOpts,
	}

	err = db.Table(common.BKTableNameObjAttDes).Update(ctx, cond.ToMapStr(), data)
	if err != nil {
		return err
	}

	procCond := condition.CreateCondition()
	procCond.Field(common.BKProtocol).Eq(nil)

	procData := mapstr.MapStr{
		common.BKProtocol: "1",
	}
	err = db.Table(common.BKTableNameBaseProcess).Update(ctx, procCond.ToMapStr(), procData)
	if err != nil {
		return err
	}

	return nil

}
