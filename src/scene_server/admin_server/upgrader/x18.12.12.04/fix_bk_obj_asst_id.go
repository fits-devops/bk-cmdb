package x18_12_12_04

import (
	"context"
	"fmt"
	"strings"

	"configcenter/src/common"
	"configcenter/src/common/condition"
	"configcenter/src/common/mapstr"
	"configcenter/src/common/metadata"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func fixBKObjAsstID(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {

	findCond := condition.CreateCondition()
	findCond.Field(common.AssociationKindIDField).NotEq(common.AssociationKindMainline)
	objassts := []metadata.Association{}
	err := db.Table(common.BKTableNameObjAsst).Find(findCond.ToMapStr()).All(ctx, &objassts)
	if err != nil {
		return err
	}

	for _, objasst := range objassts {
		name := objasst.AssociationName
		if common.GetObjByType(objasst.AsstObjID) == common.BKInnerObjIDObject {
			name = strings.TrimPrefix(name, "bk_")
		}
		if countCharacter(name, '_') > 1 {
			continue
		}

		cond := condition.CreateCondition()
		cond.Field(common.BKFieldID).Eq(objasst.ID)

		newObjAsstID := buildObjAsstID(objasst)
		data := mapstr.MapStr{
			common.AssociationObjAsstIDField: newObjAsstID,
		}
		err := db.Table(common.BKTableNameObjAsst).Update(ctx, mapstr.MapStr{common.BKFieldID: objasst.ID}, data)
		if err != nil {
			return err
		}

		err = db.Table(common.BKTableNameInstAsst).Update(
			ctx,
			mapstr.MapStr{
				common.BKObjIDField:     objasst.ObjectID,
				common.BKAsstObjIDField: objasst.AsstObjID,
			},
			data)
		if err != nil {
			return err
		}

	}
	return nil
}

func buildObjAsstID(asst metadata.Association) string {
	return fmt.Sprintf("%s_%s_%s", asst.ObjectID, asst.AsstKindID, asst.AsstObjID)
}

func countCharacter(src string, sub rune) int {
	count := 0
	for _, s := range src {
		if s == sub {
			count++
		}
	}
	return count
}
