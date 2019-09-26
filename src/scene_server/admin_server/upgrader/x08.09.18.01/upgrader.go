package x08_09_18_01

import (
	"context"
	"time"

	"configcenter/src/common"
	"configcenter/src/common/mapstr"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func fixedHostPlatAssocateRelation(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {

	type instAsstStruct struct {
		ID           int64     `bson:"id"`
		InstID       int64     `bson:"inst_id"`
		ObjectID     string    `bson:"obj_id"`
		AsstInstID   int64     `bson:"asst_inst_id"`
		AsstObjectID string    `bson:"asst_obj_id"`
		OwnerID      string    `bson:"org_id"`
		CreateTime   time.Time `bson:"create_time"`
		LastTime     time.Time `bson:"last_time"`
	}

	instAsstArr := make([]instAsstStruct, 0)
	instAsstConditionMap := mapstr.MapStr{
		common.BKObjIDField:      common.BKInnerObjIDHost,
		common.BKAsstInstIDField: common.BKInnerObjIDPlat,
	}

	err = db.Table(common.BKTableNameInstAsst).Find(instAsstConditionMap).Fields(common.BKHostIDField).All(ctx, &instAsstArr)
	if nil != err && !db.IsNotFoundError(err) {
		return err
	}
	var exitsAsstHostIDArr []int64
	for _, instAsst := range instAsstArr {
		exitsAsstHostIDArr = append(exitsAsstHostIDArr, instAsst.InstID)
	}

	type hostInfoStruct struct {
		HostID  int64  `bson:"host_id"`
		PlatID  int64  `bson:"cloud_id"`
		OwnerID string `bson:"org_id"`
	}
	hostInfoMap := make([]hostInfoStruct, 0)
	fields := []string{common.BKHostIDField, common.BKCloudIDField, common.BKOwnerIDField}
	hostCondition := make(mapstr.MapStr)
	if 0 < len(exitsAsstHostIDArr) {
		hostCondition[common.BKHostIDField] = mapstr.MapStr{common.BKDBNIN: exitsAsstHostIDArr}
	}

	err = db.Table(common.BKTableNameBaseHost).Find(hostCondition).Fields(fields...).All(ctx, &hostInfoMap)
	if err != nil && !db.IsNotFoundError(err) {
		return err
	}

	nowTime := time.Now().UTC()
	for _, host := range hostInfoMap {
		instAsstConditionMap := mapstr.MapStr{
			common.BKObjIDField:     common.BKInnerObjIDHost,
			common.BKInstIDField:    host.HostID,
			common.BKAsstObjIDField: common.BKInnerObjIDPlat,
		}
		cnt, err := db.Table(common.BKTableNameInstAsst).Find(instAsstConditionMap).Count(ctx)
		if nil != err {
			return err
		}
		if 0 == cnt {
			id, err := db.NextSequence(ctx, common.BKTableNameInstAsst)
			if nil != err {
				return err
			}
			addAsstInst := instAsstStruct{
				ID:           int64(id),
				InstID:       host.HostID,
				ObjectID:     common.BKInnerObjIDHost,
				AsstInstID:   host.PlatID,
				AsstObjectID: common.BKInnerObjIDPlat,
				OwnerID:      host.OwnerID,
				CreateTime:   nowTime,
				LastTime:     nowTime,
			}
			err = db.Table(common.BKTableNameInstAsst).Insert(ctx, addAsstInst)
			if nil != err {
				return err
			}
		}
	}
	return nil
}
