package x19_02_15_10

import (
	"context"

	"configcenter/src/common"
	"configcenter/src/common/condition"
	"configcenter/src/common/mapstr"
	"configcenter/src/common/metadata"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func fixEventSubscribeLastTime(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {
	cond := condition.CreateCondition()
	cond.Field(common.BKOwnerIDField).Eq(common.BKDefaultOwnerID)
	cond.Field(common.BKSubscriptionNameField).Like("process instance refresh")

	data := mapstr.MapStr{
		common.LastTimeField: metadata.Now(),
	}

	err := db.Table(common.BKTableNameSubscription).Update(ctx, cond.ToMapStr(), data)
	if err != nil {
		return err
	}
	return nil
}
