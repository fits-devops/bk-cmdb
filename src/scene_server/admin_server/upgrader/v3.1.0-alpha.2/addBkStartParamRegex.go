package v3v0v1alpha2

import (
	"context"
	"time"

	"configcenter/src/common"
	"configcenter/src/common/blog"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func addBkStartParamRegex(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	tablename := common.BKTableNameObjAttDes
	now := time.Now()

	type Attribute struct {
		ID                int64       `field:"id" json:"id" bson:"id"`
		OwnerID           string      `field:"org_id" json:"org_id" bson:"org_id"`
		ObjectID          string      `field:"obj_id" json:"obj_id" bson:"obj_id"`
		PropertyID        string      `field:"property_id" json:"property_id" bson:"property_id"`
		PropertyName      string      `field:"property_name" json:"property_name" bson:"property_name"`
		PropertyGroup     string      `field:"property_group" json:"property_group" bson:"property_group"`
		PropertyGroupName string      `field:"property_group_name,ignoretomap" json:"property_group_name" bson:"-"`
		PropertyIndex     int64       `field:"property_index" json:"property_index" bson:"property_index"`
		Unit              string      `field:"unit" json:"unit" bson:"unit"`
		Placeholder       string      `field:"placeholder" json:"placeholder" bson:"placeholder"`
		IsEditable        bool        `field:"editable" json:"editable" bson:"editable"`
		IsPre             bool        `field:"ispre" json:"ispre" bson:"ispre"`
		IsRequired        bool        `field:"isrequired" json:"isrequired" bson:"isrequired"`
		IsReadOnly        bool        `field:"isreadonly" json:"isreadonly" bson:"isreadonly"`
		IsOnly            bool        `field:"isonly" json:"isonly" bson:"isonly"`
		IsSystem          bool        `field:"issystem" json:"issystem" bson:"issystem"`
		IsAPI             bool        `field:"isapi" json:"isapi" bson:"isapi"`
		PropertyType      string      `field:"property_type" json:"property_type" bson:"property_type"`
		Option            interface{} `field:"option" json:"option" bson:"option"`
		Description       string      `field:"description" json:"description" bson:"description"`
		Creator           string      `field:"creator" json:"creator" bson:"creator"`
		CreateTime        *time.Time  `json:"create_time" bson:"create_time"`
		LastTime          *time.Time  `json:"last_time" bson:"last_time"`
	}

	row := &Attribute{
		ObjectID:      common.BKInnerObjIDProc,
		PropertyID:    "start_param_regex",
		PropertyName:  "启动参数匹配规则",
		IsRequired:    false,
		IsOnly:        false,
		IsEditable:    true,
		PropertyGroup: "default",
		PropertyType:  common.FieldTypeLongChar,
		Option:        "",
		OwnerID:       conf.OwnerID,
		IsPre:         true,
		IsReadOnly:    false,
		CreateTime:    &now,
		Creator:       common.CCSystemOperatorUserName,
		LastTime:      &now,
		Description:   "通过进程启动参数唯一识别进程，比如kafka和zookeeper的二进制名称为java，通过启动参数包含kafka或zookeeper来区分",
	}
	_, _, err = upgrader.Upsert(ctx, db, tablename, row, "id", []string{common.BKObjIDField, common.BKPropertyIDField, common.BKOwnerIDField}, []string{})
	if nil != err {
		blog.Errorf("[upgrade v3.1.0-alpha.2] addBkStartParamRegex  %s", err)
		return err
	}

	return nil
}

func updateLanguageField(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	condition := map[string]interface{}{
		common.BKObjIDField:      common.BKInnerObjIDApp,
		common.BKPropertyIDField: "language",
	}
	data := map[string]interface{}{
		"isreadonly": false,
		"editable":   true,
	}
	err = db.Table(common.BKTableNameObjAttDes).Update(ctx, condition, data)
	if nil != err {
		blog.Errorf("[upgrade v3.1.0-alpha.2] updateLanguageField error  %s", err.Error())
		return err
	}
	return nil
}
