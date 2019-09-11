package metadata

import (
	"configcenter/src/common/mapstr"
)

const (
	GroupFieldID         = "id"
	GroupFieldGroupID    = "group_id"
	GroupFieldGroupName  = "group_name"
	GroupFieldGroupIndex = "group_index"
	GroupFieldObjectID   = "obj_id"
	GroupFieldOwnerID    = "org_id"
	GroupFieldIsDefault  = "isdefault"
	GroupFieldIsPre      = "ispre"
)

// PropertyGroupObjectAtt uset to update or delete the property group object attribute
type PropertyGroupObjectAtt struct {
	Condition struct {
		OwnerID    string `field:"org_id" json:"org_id"`
		ObjectID   string `field:"obj_id" json:"obj_id"`
		PropertyID string `field:"property_id" json:"property_id"`
	} `json:"condition"`
	Data struct {
		PropertyGroupID string `field:"property_group" json:"property_group"`
		PropertyIndex   int    `field:"property_index" json:"property_index"`
	} `json:"data"`
}

// Group group metadata definition
type Group struct {
	Metadata   `field:"metadata" json:"metadata" bson:"metadata"`
	ID         int64  `field:"id" json:"id" bson:"id"`
	GroupID    string `field:"group_id" json:"group_id" bson:"group_id"`
	GroupName  string `field:"group_name" json:"group_name" bson:"group_name"`
	GroupIndex int64  `field:"group_index" json:"group_index" bson:"group_index"`
	ObjectID   string `field:"obj_id" json:"obj_id" bson:"obj_id"`
	OwnerID    string `field:"org_id" json:"org_id" bson:"org_id"`
	IsDefault  bool   `field:"isdefault" json:"isdefault" bson:"isdefault"`
	IsPre      bool   `field:"ispre" json:"ispre" bson:"ispre"`
}

// Parse load the data from mapstr group into group instance
func (cli *Group) Parse(data mapstr.MapStr) (*Group, error) {

	err := mapstr.SetValueToStructByTags(cli, data)
	if nil != err {
		return nil, err
	}

	return cli, err
}

// ToMapStr to mapstr
func (cli *Group) ToMapStr() mapstr.MapStr {
	return mapstr.SetValueToMapStrByTags(cli)
}
