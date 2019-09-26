package x19_02_15_10

import (
	"context"

	"configcenter/src/common"
	"configcenter/src/common/condition"
	"configcenter/src/common/mapstr"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func fixAssociationTypeName(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {

	nameKV := map[string]string{
		"run":         "运行",
		"group":       "组成",
		"default":     "默认关联",
		"cover":       "覆盖",
		"connect":     "上联",
		"bk_mainline": "拓扑组成",
		"belong":      "属于",
	}

	for id, name := range nameKV {
		cond := condition.CreateCondition()
		cond.Field(common.BKOwnerIDField).Eq(common.BKDefaultOwnerID)
		cond.Field(common.AssociationKindIDField).Eq(id)

		data := mapstr.MapStr{
			common.AssociationKindNameField: name,
		}

		err := db.Table(common.BKTableNameAsstDes).Update(ctx, cond.ToMapStr(), data)
		if err != nil {
			return err
		}
	}
	return nil
}
