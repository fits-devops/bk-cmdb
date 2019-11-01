package x19_09_29_01

import (
	"configcenter/src/common"
	"configcenter/src/common/metadata"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
	"context"
)

func UpdateIaasAssociation(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {
	falseVar := false
	idcrackAssts := []metadata.Association{
		{
			OwnerID:         conf.OwnerID,
			AsstKindID:      "contain",
			ObjectID:        "idc",
			AsstObjID:       "idcrack",
			AssociationName: "idc_contain_idcrack",
			AssociationAliasName: "包含机柜",
			Mapping:         metadata.OneToManyMapping,
			OnDelete:        metadata.NoAction,
			IsPre:           &falseVar,
		},
		{
			OwnerID:         conf.OwnerID,
			AsstKindID:      "connect",
			ObjectID:        "idcrack",
			AsstObjID:       "user",
			AssociationName: "idcrack_connect_user",
			AssociationAliasName: "运维负责人",
			Mapping:         metadata.ManyToManyMapping,
			OnDelete:        metadata.NoAction,
			IsPre:           &falseVar,
		},
	}

	for _, idcrackAsst := range idcrackAssts {
		_, _, err := upgrader.Upsert(ctx, db, common.BKTableNameObjAsst, idcrackAsst, "id", []string{"obj_id", "asst_obj_id"}, []string{"id"})
		if err != nil {
			return err
		}
	}
	return nil
}
