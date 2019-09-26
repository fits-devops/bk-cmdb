package x19_04_16_03

import (
	"context"
	"time"

	"configcenter/src/common"
	"configcenter/src/common/condition"
	"configcenter/src/common/mapstr"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func updateAttributeCreateTime(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {
	var start uint64

	type Attribute struct {
		CreaetTime *time.Time `bson:"creaet_time"`
		CreateTime *time.Time `bson:"create_time"`
		ID         uint64     `bson:"id"`
	}

	now := time.Now()
	for {
		attrs := []Attribute{}
		err := db.Table(common.BKTableNameObjAttDes).Find(nil).Start(start).Limit(50).All(ctx, &attrs)
		if err != nil {
			return err
		}
		if len(attrs) <= 0 {
			break
		}
		start += 50

		for _, attr := range attrs {
			if attr.CreateTime == nil {
				createTime := attr.CreaetTime
				if createTime == nil {
					createTime = &now
				}
				if attr.CreateTime != nil {
					createTime = attr.CreateTime
				}

				cond := condition.CreateCondition()
				cond.Field(common.BKFieldID).Eq(attr.ID)
				err := db.Table(common.BKTableNameObjAttDes).Update(ctx, cond.ToMapStr(), mapstr.MapStr{common.CreateTimeField: createTime})
				if err != nil {
					return err
				}
			}
		}
	}

	return db.Table(common.BKTableNameObjAttDes).DropColumn(ctx, "creaet_time")
}
