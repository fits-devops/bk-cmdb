package metadata

/**
Student学生信息
 */

import (
	"configcenter/src/common/mapstr"
)

const (
	StudentFieldID                     = "id"
	StudentFieldStudentID              = "student_id"			//学生ID
	StudentFieldStudentName            = "student_name"		//学生姓名
	StudentFieldStudentSex			   = "student_sex"		//学生性别
	StudentFieldStudentAge			   = "student_age"		//学生年龄
	StudentFieldClassID				   = "class_id"			//班级ID
	StudentFieldOwnerID				   = "owner_id"			//组织ID（拥有者）
	StudentFieldCreateTime			   = "create_time"      //创建时间
	StudentFieldLastTime			   = "last_time"		//最后更新时间
)

// 定义学生的数据结构
type Student struct {
	Metadata           `field:"metadata" json:"metadata" bson:"metadata"`
	ID                 int64  `field:"id" json:"id" bson:"id"`
	StudentID   string `field:"student_id"  json:"student_id" bson:"student_id"`
	StudentName	string `field:"student_name" json:"student_name" bson:"student_name"`
	StudentSex  string	`field:"student_sex" json:"student_sex" bson:"student_sex"`
	StudentAge	int64	`field:"student_age" json:"student_age" bson:"student_age"`
	ClassID		string	`field:"class_id" json:"class_id" bson:"class_id"`
	OwnerID     string `field:"owner_id" json:"owner_id" bson:"owner_id"  `
	CreateTime  *Time  `field:"create_time" json:"create_time" bson:"create_time"`
	LastTime    *Time  `field:"last_time" json:"last_time" bson:"last_time"`
}

// 将API传入的mapstr数据，解析成Student结构
func (s *Student) Parse(data mapstr.MapStr) (*Student, error) {

	err := mapstr.SetValueToStructByTags(s, data)
	if nil != err {
		return nil, err
	}

	return s, err
}

// 将Student 转换成mapstr
func (s *Student) ToMapStr() mapstr.MapStr {
	return mapstr.SetValueToMapStrByTags(s)
}
