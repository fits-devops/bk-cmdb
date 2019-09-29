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
			AsstKindID:      "belong",
			ObjectID:        "idcrack",
			AsstObjID:       "idc",
			AssociationName: "idcrack_belong_idc",
			Mapping:         metadata.OneToOneMapping,
			OnDelete:        metadata.NoAction,
			IsPre:           &falseVar,
		},
		{
			OwnerID:         conf.OwnerID,
			AsstKindID:      "belong",
			ObjectID:        "host",
			AsstObjID:       "idcrack",
			AssociationName: "host_belong_idcrack",
			Mapping:         metadata.OneToOneMapping,
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
