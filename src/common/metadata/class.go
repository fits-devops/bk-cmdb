package metadata

/**
class班级信息
 */

import (
	"configcenter/src/common/mapstr"
)

const (
	ClassFieldID                   = "id"
	ClassFieldClassID              = "class_id"			//班级ID
	ClassFieldClassName            = "class_name"		//班级名称
	ClassFieldClassType            = "class_type"		//班级类型
	ClassFieldClassIcon            = "class_icon"		//班级图标（班徽）
	ClassFieldClassOwnerID 		   = "owner_id"			//拥有者
)

// 定义班级的数据结构
type Class struct {
	Metadata           `field:"metadata" json:"metadata" bson:"metadata"`
	ID                 int64  `field:"id" json:"id" bson:"id"`
	ClassID   string `field:"class_id"  json:"class_id" bson:"class_id"`
	ClassName string `field:"class_name" json:"class_name" bson:"class_name"`
	ClassType string `field:"class_type" json:"class_type" bson:"class_type"`
	ClassIcon string `field:"class_icon" json:"class_icon" bson:"class_icon"`
	OwnerID            string `field:"owner_id" json:"owner_id" bson:"owner_id"  `
}

// 将API传入的mapstr数据，解析成Class结构体
func (c *Class) Parse(data mapstr.MapStr) (*Class, error) {

	err := mapstr.SetValueToStructByTags(c, data)
	if nil != err {
		return nil, err
	}

	return c, err
}

// 将Class 转换成mapstr
func (c *Class) ToMapStr() mapstr.MapStr {
	return mapstr.SetValueToMapStrByTags(c)
}
