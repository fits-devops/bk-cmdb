package x18_12_12_06

import (
	"context"
	"fmt"

	"configcenter/src/common"
	"configcenter/src/common/condition"
	"configcenter/src/common/metadata"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func keyfunc(a, b string) string { return a + ":" + b }

func reconcilUnique(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {
	attrs := []metadata.Attribute{}

	attrCond := condition.CreateCondition()
	attrCond.Field(common.BKObjIDField).In([]string{
		common.BKInnerObjIDSwitch,
		common.BKInnerObjIDRouter,
		common.BKInnerObjIDBlance,
		common.BKInnerObjIDFirewall,
		common.BKInnerObjIDWeblogic,
		common.BKInnerObjIDTomcat,
		common.BKInnerObjIDApache})
	err := db.Table(common.BKTableNameObjAttDes).Find(attrCond.ToMapStr()).All(ctx, &attrs)
	if err != nil {
		return err
	}
	var propertyIDToProperty = map[string]metadata.Attribute{}

	for _, attr := range attrs {
		propertyIDToProperty[keyfunc(attr.ObjectID, attr.PropertyID)] = attr
	}

	shouldCheck := []string{
		keyfunc(common.BKInnerObjIDSwitch, common.BKAssetIDField),
		keyfunc(common.BKInnerObjIDRouter, common.BKAssetIDField),
		keyfunc(common.BKInnerObjIDBlance, common.BKAssetIDField),
		keyfunc(common.BKInnerObjIDFirewall, common.BKAssetIDField),
		keyfunc(common.BKInnerObjIDWeblogic, common.BKInstKeyField),
		keyfunc(common.BKInnerObjIDTomcat, common.BKInstKeyField),
		keyfunc(common.BKInnerObjIDApache, common.BKInstKeyField),
	}

	if notExistFields := checkKeysShouldExists(propertyIDToProperty, shouldCheck); len(notExistFields) > 0 {
		return fmt.Errorf("expected field not exists: %v", notExistFields)
	}

	uniques := []metadata.ObjectUnique{
		buildUnique(propertyIDToProperty, common.BKInnerObjIDSwitch, common.BKAssetIDField),
		buildUnique(propertyIDToProperty, common.BKInnerObjIDRouter, common.BKAssetIDField),
		buildUnique(propertyIDToProperty, common.BKInnerObjIDBlance, common.BKAssetIDField),
		buildUnique(propertyIDToProperty, common.BKInnerObjIDFirewall, common.BKAssetIDField),
		buildUnique(propertyIDToProperty, common.BKInnerObjIDWeblogic, common.BKInstKeyField),
		buildUnique(propertyIDToProperty, common.BKInnerObjIDTomcat, common.BKInstKeyField),
		buildUnique(propertyIDToProperty, common.BKInnerObjIDApache, common.BKInstKeyField),
	}

	for _, unique := range uniques {
		exists, err := isUniqueExists(ctx, db, conf, unique)
		if err != nil {
			return err
		}
		if exists {
			continue
		}

		uid, err := db.NextSequence(ctx, common.BKTableNameObjUnique)
		if err != nil {
			return err
		}
		unique.ID = uid
		if err := db.Table(common.BKTableNameObjUnique).Insert(ctx, unique); err != nil {
			return err
		}
	}

	return nil
}

func buildUnique(propertyIDToProperty map[string]metadata.Attribute, model, field string) metadata.ObjectUnique {
	return metadata.ObjectUnique{
		ObjID:     model,
		MustCheck: true,
		Keys: []metadata.UniqueKey{
			{
				Kind: metadata.UniqueKeyKindProperty,
				ID:   uint64(propertyIDToProperty[keyfunc(model, field)].ID),
			},
		},
		Ispre:    false,
		OwnerID:  common.BKDefaultOwnerID,
		LastTime: metadata.Now(),
	}
}

func checkKeysShouldExists(m map[string]metadata.Attribute, shouldExistKeys []string) []string {
	notValidKeys := []string{}
	for _, k := range shouldExistKeys {
		if _, ok := m[k]; !ok {
			notValidKeys = append(notValidKeys, k)
		}
	}
	return notValidKeys
}

func isUniqueExists(ctx context.Context, db dal.RDB, conf *upgrader.Config, unique metadata.ObjectUnique) (bool, error) {
	keyhash := unique.KeysHash()
	uniqueCond := condition.CreateCondition()
	uniqueCond.Field(common.BKObjIDField).Eq(unique.ObjID)
	uniqueCond.Field(common.BKOwnerIDField).Eq(conf.OwnerID)
	existUniques := []metadata.ObjectUnique{}

	err := db.Table(common.BKTableNameObjUnique).Find(uniqueCond.ToMapStr()).All(ctx, &existUniques)
	if err != nil {
		return false, err
	}

	for _, uni := range existUniques {
		if uni.KeysHash() == keyhash {
			return true, nil
		}
	}
	return false, nil

}
