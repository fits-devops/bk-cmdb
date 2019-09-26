package x19_01_18_01

import (
	"context"
	"strings"

	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func dropIndex(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	if err = db.Table("cc_TopoGraphics").
		DropIndex(ctx, "scope_id_1_node_type_1_obj_id_1_inst_id_1_scope_type_1"); err != nil &&
		!strings.Contains(err.Error(), "not found") {
		return err
	}
	return nil
}
