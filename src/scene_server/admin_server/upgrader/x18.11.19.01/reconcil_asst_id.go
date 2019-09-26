package x18_11_19_01

// import (
// 	"context"

// 	"github.com/rs/xid"

// 	"configcenter/src/common"
// 	"configcenter/src/common/condition"
// 	"configcenter/src/common/mapstr"
// 	"configcenter/src/scene_server/admin_server/upgrader"
// 	"configcenter/src/storage/dal"
// )

// func reconcilAsstID(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {
// 	start, limit := uint64(0), uint64(100)

// 	type HostInst struct {
// 		HostID  uint64 `bson:"host_id"`
// 		AssetID string `bson:"asset_id"`
// 		OwnerID string `bson:"org_id"`
// 	}

// 	cond := condition.CreateCondition()
// 	or := cond.NewOR()
// 	or.Item(mapstr.MapStr{common.BKAssetIDField: nil})
// 	or.Item(mapstr.MapStr{common.BKAssetIDField: ""})

// 	for {
// 		hosts := []HostInst{}

// 		err := db.Table(common.BKTableNameBaseHost).Find(cond.ToMapStr()).
// 			Sort(common.BKHostIDField).
// 			Start(start).Limit(limit).All(ctx, &hosts)
// 		if err != nil {
// 			return err
// 		}

// 		for index := range hosts {
// 			if hosts[index].AssetID == "" {
// 				data := mapstr.MapStr{
// 					common.BKAssetIDField: xid.New().String(),
// 				}
// 				updateCond := condition.CreateCondition()
// 				updateCond.Field(common.BKHostIDField).Eq(hosts[index].HostID)
// 				updateCond.Field(common.BKOwnerIDField).Eq(hosts[index].OwnerID)

// 				if err := db.Table(common.BKTableNameBaseHost).
// 					Update(ctx, updateCond.ToMapStr(), data); err != nil {
// 					return err
// 				}
// 			}
// 		}

// 		if uint64(len(hosts)) < limit {
// 			break
// 		}

// 		start += limit
// 	}

// 	return nil
// }
