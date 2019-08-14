package core

import (
	"configcenter/src/common/metadata"
)

// 定义班级操作接口
type ClassOperation interface {
	CreateOneClass(ctx ContextParams, inputParam metadata.CreateOneClass) (*metadata.CreateOneDataResult, error)
	CreateManyClass(ctx ContextParams, inputParam metadata.CreateManyClass) (*metadata.CreateManyDataResult, error)
	SetManyClass(ctx ContextParams, inputParam metadata.SetManyClass) (*metadata.SetDataResult, error)
	SetOneClass(ctx ContextParams, inputParam metadata.SetOneClass) (*metadata.SetDataResult, error)
	UpdateClass(ctx ContextParams, inputParam metadata.UpdateOption) (*metadata.UpdatedCount, error)
	DeleteClass(ctx ContextParams, inputParam metadata.DeleteOption) (*metadata.DeletedCount, error)
	SearchClass(ctx ContextParams, inputParam metadata.QueryCondition) (*metadata.QueryClassDataResult, error)
}

// 定义学生操作接口
type StudentOperation interface {
	ClassOperation  //继承班级操作接口

	CreateStudent(ctx ContextParams, inputParam metadata.CreateStudent) (*metadata.CreateOneDataResult, error)
	//SetStudent(ctx ContextParams, inputParam metadata.SetStudent) (*metadata.SetDataResult, error)
	//UpdateStudent(ctx ContextParams, inputParam metadata.UpdateOption) (*metadata.UpdatedCount, error)
	//DeleteStudent(ctx ContextParams, inputParam metadata.DeleteOption) (*metadata.DeletedCount, error)
	//CascadeDeleteStudent(ctx ContextParams, inputParam metadata.DeleteOption) (*metadata.DeletedCount, error)
	//SearchStudent(ctx ContextParams, inputParam metadata.QueryCondition) (*metadata.QueryStudentDataResult, error)
	//SearchStudentWithClass(ctx ContextParams, inputParam metadata.QueryCondition) (*metadata.QueryStudentWithClassDataResult, error)
}

// 定义core核心入口方法
type Core interface {
	StudentOperation() StudentOperation
}

type core struct {
	student           StudentOperation
}

func (c *core) StudentOperation() StudentOperation {
	return c.student
}

// New create core
func New(student StudentOperation) Core {
	return &core{
		student:           student,
	}
}



