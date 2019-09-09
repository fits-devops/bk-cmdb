package metadata

import (
	"configcenter/src/common/mapstr"
)

const (
	ClassFieldID                 = "id"
	ClassFieldClassificationID   = "classification_id"
	ClassFieldClassificationName = "classification_name"
	ClassFieldClassificationType = "classification_type"
	ClassFieldClassificationIcon = "classification_icon"
	ClassFieldOwnerID            = "org_id"
)

// Classification the classification metadata definition
type Classification struct {
	Metadata           `field:"metadata" json:"metadata" bson:"metadata"`
	ID                 int64  `field:"id" json:"id" bson:"id"`
	ClassificationID   string `field:"classification_id"  json:"classification_id" bson:"classification_id"`
	ClassificationName string `field:"classification_name" json:"classification_name" bson:"classification_name"`
	ClassificationType string `field:"classification_type" json:"classification_type" bson:"classification_type"`
	ClassificationIcon string `field:"classification_icon" json:"classification_icon" bson:"classification_icon"`
	OwnerID            string `field:"org_id" json:"org_id" bson:"org_id"  `
}

// Parse load the data from mapstr classification into classification instance
func (cli *Classification) Parse(data mapstr.MapStr) (*Classification, error) {

	err := mapstr.SetValueToStructByTags(cli, data)
	if nil != err {
		return nil, err
	}

	return cli, err
}

// ToMapStr to mapstr
func (cli *Classification) ToMapStr() mapstr.MapStr {
	return mapstr.SetValueToMapStrByTags(cli)
}
