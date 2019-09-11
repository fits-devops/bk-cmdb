package metadata

import (
	"configcenter/src/common/mapstr"
)

const (
	AttributeFieldID            = "id"
	AttributeFieldOwnerID       = "org_id"
	AttributeFieldObjectID      = "obj_id"
	AttributeFieldPropertyID    = "property_id"
	AttributeFieldPropertyName  = "property_name"
	AttributeFieldPropertyGroup = "property_group"
	AttributeFieldPropertyIndex = "property_index"
	AttributeFieldPropertyType  = "property_type"
	AttributeFieldUnit          = "unit"
	AttributeFieldPlaceHoler    = "placeholder"
	AttributeFieldIsEditable    = "editable"
	AttributeFieldIsPre         = "ispre"
	AttributeFieldIsRequired    = "isrequired"
	AttributeFieldIsReadOnly    = "isreadonly"
	AttributeFieldIsOnly        = "isonly"
	AttributeFieldIsSystem      = "issystem"
	AttributeFieldIsAPI         = "isapi"
	AttributeFieldOption        = "option"
	AttributeFieldDescription   = "description"
	AttributeFieldCreator       = "creator"
	AttributeFieldCreateTime    = "create_time"
	AttributeFieldLastTime      = "last_time"
)

// Attribute attribute metadata definition
type Attribute struct {
	Metadata          `field:"metadata" json:"metadata" bson:"metadata"`
	ID                int64       `field:"id" json:"id" bson:"id"`
	OwnerID           string      `field:"org_id" json:"org_id" bson:"org_id"`
	ObjectID          string      `field:"obj_id" json:"obj_id" bson:"obj_id"`
	PropertyID        string      `field:"property_id" json:"property_id" bson:"property_id"`
	PropertyName      string      `field:"property_name" json:"property_name" bson:"property_name"`
	PropertyGroup     string      `field:"property_group" json:"property_group" bson:"property_group"`
	PropertyGroupName string      `field:"property_group_name,ignoretomap" json:"property_group_name" bson:"-"`
	PropertyIndex     int64       `field:"property_index" json:"property_index" bson:"property_index"`
	PropertyType      string      `field:"property_type" json:"property_type" bson:"property_type"`
	Unit              string      `field:"unit" json:"unit" bson:"unit"`
	Placeholder       string      `field:"placeholder" json:"placeholder" bson:"placeholder"`
	IsEditable        bool        `field:"editable" json:"editable" bson:"editable"`
	IsPre             bool        `field:"ispre" json:"ispre" bson:"ispre"`
	IsRequired        bool        `field:"isrequired" json:"isrequired" bson:"isrequired"`
	IsReadOnly        bool        `field:"isreadonly" json:"isreadonly" bson:"isreadonly"`
	IsOnly            bool        `field:"isonly" json:"isonly" bson:"isonly"`
	IsSystem          bool        `field:"issystem" json:"issystem" bson:"issystem"`
	IsAPI             bool        `field:"isapi" json:"isapi" bson:"isapi"`
	Option            interface{} `field:"option" json:"option" bson:"option"`
	Description       string      `field:"description" json:"description" bson:"description"`
	Creator           string      `field:"creator" json:"creator" bson:"creator"`
	CreateTime        *Time       `json:"create_time" bson:"create_time"`
	LastTime          *Time       `json:"last_time" bson:"last_time"`
}

// AttributeGroup attribute metadata definition
type AttributeGroup struct {
	ID         int64  `field:"id" json:"id" bson:"id"`
	OwnerID    string `field:"org_id" json:"org_id" bson:"org_id"`
	ObjectID   string `field:"obj_id" json:"obj_id" bson:"obj_id"`
	IsDefault  bool   `field:"isdefault" json:"isdefault" bson:"isdefault"`
	IsPre      bool   `field:"ispre" json:"ispre" bson:"ispre"`
	GroupID    string `field:"group_id" json:"group_id" bson:"group_id"`
	GroupName  string `field:"group_name" json:"group_name" bson:"group_name"`
	GroupIndex int64  `field:"group_index" json:"group_index" bson:"group_index"`
}

// Parse load the data from mapstr attribute into attribute instance
func (cli *Attribute) Parse(data mapstr.MapStr) (*Attribute, error) {

	err := mapstr.SetValueToStructByTags(cli, data)
	if nil != err {
		return nil, err
	}

	return cli, err
}

// ToMapStr to mapstr
func (cli *Attribute) ToMapStr() mapstr.MapStr {
	return mapstr.SetValueToMapStrByTags(cli)
}

// ObjAttDes 对象模型属性
type ObjAttDes struct {
	Attribute         `json:",inline" bson:",inline"`
	PropertyGroupName string `json:"property_group_name"`
}
