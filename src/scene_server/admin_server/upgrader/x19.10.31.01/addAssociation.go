package x19_10_31_01

import (
	"configcenter/src/common"
	"configcenter/src/common/metadata"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
	"context"
)

func addAssociation(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {
	falseVar := false
	Asst := metadata.Association{
		OwnerID:         conf.OwnerID,
		AsstKindID:      "contain",
		ObjectID:        "user_group",
		AsstObjID:       "user",
		AssociationName: "user_group_contain_user",
		Mapping:         metadata.OneToManyMapping,
		OnDelete:        metadata.NoAction,
		IsPre:           &falseVar,
	}

	_, _, err := upgrader.Upsert(ctx, db, common.BKTableNameObjAsst, Asst, "id", []string{"obj_id", "asst_obj_id"}, []string{"id"})
	if err != nil {
		return err
	}

	return nil
}
