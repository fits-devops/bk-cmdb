package x18_12_13_01

import (
	"context"
	"strings"

	"configcenter/src/common"
	"configcenter/src/common/metadata"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func addswitchAssociation(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {
	falseVar := false
	switchAsst := metadata.Association{
		OwnerID:         conf.OwnerID,
		AsstKindID:      "connect",
		ObjectID:        "host",
		AsstObjID:       "switch",
		AssociationName: "host_connect_switch",
		Mapping:         metadata.ManyToManyMapping,
		OnDelete:        metadata.NoAction,
		IsPre:           &falseVar,
	}

	_, _, err := upgrader.Upsert(ctx, db, common.BKTableNameObjAsst, switchAsst, "id", []string{"obj_id", "asst_obj_id"}, []string{"id"})
	if err != nil {
		return err
	}

	return nil
}

func changeNetDeviceTableName(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {
	err := db.DropTable("cc_Netcollect_Device")
	if err != nil && !strings.Contains(err.Error(), "ns not found") {
		return err
	}
	err = db.DropTable("cc_Netcollect_Property")
	if err != nil && !strings.Contains(err.Error(), "ns not found") {
		return err
	}

	tablenames := []string{"cc_NetcollectDevice", "cc_NetcollectProperty"}
	for _, tablename := range tablenames {
		exists, err := db.HasTable(tablename)
		if err != nil {
			return err
		}
		if !exists {
			if err = db.CreateTable(tablename); err != nil && !db.IsDuplicatedError(err) {
				return err
			}
		}
	}
	return nil
}
