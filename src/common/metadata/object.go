package metadata

import (
	"configcenter/src/common"
	"configcenter/src/common/mapstr"
)

const (
	ModelFieldID          = "id"
	ModelFieldObjCls      = "classification_id"
	ModelFieldObjIcon     = "obj_icon"
	ModelFieldObjectID    = "obj_id"
	ModelFieldObjectName  = "obj_name"
	ModelFieldIsPre       = "ispre"
	ModelFieldIsPaused    = "ispaused"
	ModelFieldPosition    = "position"
	ModelFieldOwnerID     = "org_id"
	ModelFieldDescription = "description"
	ModelFieldCreator     = "creator"
	ModelFieldModifier    = "modifier"
	ModelFieldCreateTime  = "create_time"
	ModelFieldLastTime    = "last_time"
)

// Object object metadata definition
type Object struct {
	Metadata    `field:"metadata" json:"metadata" bson:"metadata"`
	ID          int64  `field:"id" json:"id" bson:"id"`
	ObjCls      string `field:"classification_id" json:"classification_id" bson:"classification_id"`
	ObjIcon     string `field:"obj_icon" json:"obj_icon" bson:"obj_icon"`
	ObjectID    string `field:"obj_id" json:"obj_id" bson:"obj_id"`
	ObjectName  string `field:"obj_name" json:"obj_name" bson:"obj_name"`
	IsPre       bool   `field:"ispre" json:"ispre" bson:"ispre"`
	IsPaused    bool   `field:"ispaused" json:"ispaused" bson:"ispaused"`
	Position    string `field:"position" json:"position" bson:"position"`
	OwnerID     string `field:"org_id" json:"org_id" bson:"org_id"`
	Description string `field:"description" json:"description" bson:"description"`
	Creator     string `field:"creator" json:"creator" bson:"creator"`
	Modifier    string `field:"modifier" json:"modifier" bson:"modifier"`
	CreateTime  *Time  `field:"create_time" json:"create_time" bson:"create_time"`
	LastTime    *Time  `field:"last_time" json:"last_time" bson:"last_time"`
}

// GetDefaultInstPropertyName get default inst
func (o *Object) GetDefaultInstPropertyName() string {
	return common.DefaultInstName
}

// GetInstIDFieldName get instid filed
func (o *Object) GetInstIDFieldName() string {
	return GetInstIDFieldByObjID(o.ObjectID)

}

func GetInstIDFieldByObjID(objID string) string {
	switch objID {
	case common.BKInnerObjIDApp:
		return common.BKAppIDField
	case common.BKInnerObjIDSet:
		return common.BKSetIDField
	case common.BKInnerObjIDModule:
		return common.BKModuleIDField
	case common.BKInnerObjIDObject:
		return common.BKInstIDField
	case common.BKInnerObjIDHost:
		return common.BKHostIDField
	case common.BKInnerObjIDProc:
		return common.BKProcIDField
	case common.BKInnerObjIDPlat:
		return common.BKCloudIDField
	default:
		return common.BKInstIDField
	}

}

// GetInstNameFieldName get the inst name
func (o *Object) GetInstNameFieldName() string {
	switch o.ObjectID {
	case common.BKInnerObjIDApp:
		return common.BKAppNameField
	case common.BKInnerObjIDSet:
		return common.BKSetNameField
	case common.BKInnerObjIDModule:
		return common.BKModuleNameField
	case common.BKInnerObjIDHost:
		return common.BKInstNameField
	case common.BKInnerObjIDProc:
		return common.BKProcNameField
	case common.BKInnerObjIDPlat:
		return common.BKCloudNameField
	default:
		return common.BKInstNameField
	}
}

// GetObjectType get the object type
func (o *Object) GetObjectType() string {
	switch o.ObjectID {
	case common.BKInnerObjIDApp:
		return o.ObjectID
	case common.BKInnerObjIDSet:
		return o.ObjectID
	case common.BKInnerObjIDModule:
		return o.ObjectID
	case common.BKInnerObjIDHost:
		return o.ObjectID
	case common.BKInnerObjIDProc:
		return o.ObjectID
	case common.BKInnerObjIDPlat:
		return o.ObjectID
	default:
		return common.BKInnerObjIDObject
	}
}

// GetObjectID get the object type
func (o *Object) GetObjectID() string {
	return o.ObjectID
}

// IsCommon is common object
func (o *Object) IsCommon() bool {
	switch o.ObjectID {
	case common.BKInnerObjIDApp:
		return false
	case common.BKInnerObjIDSet:
		return false
	case common.BKInnerObjIDModule:
		return false
	case common.BKInnerObjIDHost:
		return false
	case common.BKInnerObjIDProc:
		return false
	case common.BKInnerObjIDPlat:
		return false
	default:
		return true
	}
}

// Parse load the data from mapstr object into object instance
func (o *Object) Parse(data mapstr.MapStr) (*Object, error) {

	err := mapstr.SetValueToStructByTags(o, data)
	if nil != err {
		return nil, err
	}

	return o, err
}

// ToMapStr to mapstr
func (o *Object) ToMapStr() mapstr.MapStr {
	return mapstr.SetValueToMapStrByTags(o)
}

// MainLineObject main line object definition
type MainLineObject struct {
	Object        `json:",inline"`
	AssociationID string `json:"asst_obj_id"`
}

type ObjectClsDes struct {
	ID      int    `json:"id" bson:"id"`
	ClsID   string `json:"classification_id" bson:"classification_id"`
	ClsName string `json:"classification_name" bson:"classification_name"`
	ClsType string `json:"classification_type" bson:"classification_type" `
	ClsIcon string `json:"classification_icon" bson:"classification_icon"`
}

type InnerModule struct {
	ModuleID   int64  `json:"module_id"`
	ModuleName string `json:"module_name"`
}
type InnterAppTopo struct {
	SetID   int64         `json:"set_id"`
	SetName string        `json:"set_name"`
	Module  []InnerModule `json:"module"`
}

// TopoItem define topo item
type TopoItem struct {
	ClassificationID string `json:"classification_id"`
	Position         string `json:"position"`
	ObjID            string `json:"obj_id"`
	OwnerID          string `json:"org_id"`
	ObjName          string `json:"obj_name"`
}

// ObjectTopo define the common object topo
type ObjectTopo struct {
	LabelType string   `json:"label_type"`
	LabelName string   `json:"label_name"`
	Label     string   `json:"label"`
	From      TopoItem `json:"from"`
	To        TopoItem `json:"to"`
	Arrows    string   `json:"arrows"`
}
