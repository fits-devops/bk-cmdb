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
	Assts := []metadata.Association{
		{
			OwnerID:              conf.OwnerID,
			AsstKindID:           "contain",
			AssociationAliasName: "关联用户",
			ObjectID:             "user_group",
			AsstObjID:            "user",
			AssociationName:      "user_group_contain_user",
			Mapping:              metadata.OneToManyMapping,
			OnDelete:             metadata.NoAction,
			IsPre:                &falseVar,
		},
		{
			OwnerID:              conf.OwnerID,
			AsstKindID:           "default",
			AssociationAliasName: "所属机柜",
			ObjectID:             "storage",
			AsstObjID:            "idcrack",
			AssociationName:      "storage_default_idcrack",
			Mapping:              metadata.ManyToManyMapping,
			OnDelete:             metadata.NoAction,
			IsPre:                &falseVar,
		},
	}

	for _, Asst := range Assts {
		_, _, err := upgrader.Upsert(ctx, db, common.BKTableNameObjAsst, Asst, "id", []string{"obj_id", "asst_obj_id"}, []string{"id"})
		if err != nil {
			return err
		}
	}

	return nil
}
