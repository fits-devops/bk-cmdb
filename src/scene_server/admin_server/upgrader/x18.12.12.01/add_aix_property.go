package x18_12_12_01

import (
	"context"
	"time"

	"configcenter/src/common"
	"configcenter/src/common/condition"
	"configcenter/src/common/mapstr"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/scene_server/validator"
	"configcenter/src/storage/dal"
)

type Attribute struct {
	ID                int64       `json:"id" bson:"id"`
	OwnerID           string      `json:"org_id" bson:"org_id"`
	ObjectID          string      `json:"obj_id" bson:"obj_id"`
	PropertyID        string      `json:"property_id" bson:"property_id"`
	PropertyName      string      `json:"property_name" bson:"property_name"`
	PropertyGroup     string      `json:"property_group" bson:"property_group"`
	PropertyGroupName string      `json:"property_group_name" bson:"-"`
	PropertyIndex     int64       `json:"property_index" bson:"property_index"`
	Unit              string      `json:"unit" bson:"unit"`
	Placeholder       string      `json:"placeholder" bson:"placeholder"`
	IsEditable        bool        `json:"editable" bson:"editable"`
	IsPre             bool        `json:"ispre" bson:"ispre"`
	IsRequired        bool        `json:"isrequired" bson:"isrequired"`
	IsReadOnly        bool        `json:"isreadonly" bson:"isreadonly"`
	IsOnly            bool        `json:"isonly" bson:"isonly"`
	IsSystem          bool        `json:"issystem" bson:"issystem"`
	IsAPI             bool        `json:"isapi" bson:"isapi"`
	PropertyType      string      `json:"property_type" bson:"property_type"`
	Option            interface{} `json:"option" bson:"option"`
	Description       string      `json:"description" bson:"description"`
	Creator           string      `json:"creator" bson:"creator"`
	CreateTime        *time.Time  `json:"create_time" bson:"create_time"`
	LastTime          *time.Time  `json:"last_time" bson:"last_time"`
}

func addAIXProperty(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {

	cond := condition.CreateCondition()
	cond.Field(common.BKOwnerIDField).Eq(common.BKDefaultOwnerID)
	cond.Field(common.BKObjIDField).Eq(common.BKInnerObjIDHost)
	cond.Field(common.BKPropertyIDField).Eq(common.BKOSTypeField)

	ostypeProperty := Attribute{}
	err := db.Table(common.BKTableNameObjAttDes).Find(cond.ToMapStr()).One(ctx, &ostypeProperty)
	if err != nil {
		return err
	}

	enumOpts := validator.ParseEnumOption(ostypeProperty.Option)
	for _, enum := range enumOpts {
		if enum.ID == "3" {
			return nil
		}
	}

	aixEnum := validator.EnumVal{
		ID:   "3",
		Name: "AIX",
		Type: "text",
	}
	enumOpts = append(enumOpts, aixEnum)

	data := mapstr.MapStr{
		common.BKOptionField: enumOpts,
	}

	err = db.Table(common.BKTableNameObjAttDes).Update(ctx, cond.ToMapStr(), data)
	if err != nil {
		return err
	}
	return nil
}
