package student

import (
	"configcenter/src/common"
	"configcenter/src/common/blog"
	"configcenter/src/common/metadata"
	"configcenter/src/common/universalsql/mongo"
	"configcenter/src/source_controller/testservice/core"
	"configcenter/src/storage/dal"
)

var _ core.StudentOperation = (*studentManager)(nil)

//student学生信息模块的总入口
type studentManager struct {
	*classManager  //继承class班级管理的接口方法
	dbProxy   dal.RDB
}

// 初始化一个学生信息管理的实例
func New(dbProxy dal.RDB) core.StudentOperation {

	coreMgr := &studentManager{dbProxy: dbProxy}

	coreMgr.classManager = &classManager{dbProxy: dbProxy, std: coreMgr}

	return coreMgr
}

func (m *studentManager) CreateStudent(ctx core.ContextParams, inputParam metadata.CreateStudent) (*metadata.CreateOneDataResult, error) {

	dataResult := &metadata.CreateOneDataResult{}

	// check the model attributes value
	if 0 == len(inputParam.Std.StudentID) {
		blog.Errorf("request(%s): it is failed to create a new model, because of the modelID (%s) is not set", ctx.ReqID, inputParam.Std.StudentID)
		return dataResult, ctx.Error.Errorf(common.CCErrCommParamsNeedSet, metadata.StudentFieldStudentID)
	}

	// check the input class ID
	isValid, err := m.classManager.isValid(ctx, inputParam.Std.ClassID)
	if nil != err {
		blog.Errorf("request(%s): it is failed to check whether the classificationID(%s) is invalid, error info is %s", ctx.ReqID, inputParam.Std.ClassID, err.Error())
		return dataResult, err
	}

	if !isValid {
		blog.Warnf("request(%s): it is failed to create a new model, because of the classificationID (%s) is invalid", ctx.ReqID, inputParam.Std.ClassID)
		return dataResult, ctx.Error.Errorf(common.CCErrCommParamsIsInvalid, metadata.ClassFieldID)
	}

	// check the model if it is exists
	condCheckModel := mongo.NewCondition()
	condCheckModel.Element(&mongo.Eq{Key: metadata.StudentFieldStudentID, Val: inputParam.Std.StudentID})
	condCheckModel.Element(&mongo.Eq{Key: metadata.StudentFieldOwnerID, Val: ctx.SupplierAccount})

	_, exists, err := m.isExists(ctx, condCheckModel)
	if nil != err {
		blog.Errorf("request(%s): it is failed to check whether the model (%s) is exists, error info is %s ", ctx.ReqID, inputParam.Std.StudentID, err.Error())
		return dataResult, err
	}

	if exists {
		blog.Warnf("request(%s): it is failed to  create a new model , because of the model (%s) is already exists ", ctx.ReqID, inputParam.Std.StudentID)
		return dataResult, ctx.Error.Errorf(common.CCErrCommDuplicateItem, "")
	}
	inputParam.Std.OwnerID = ctx.SupplierAccount
	id, err := m.save(ctx, &inputParam.Std)
	if nil != err {
		blog.Errorf("request(%s): it is failed to save the model (%#v), error info is %s", ctx.ReqID, inputParam.Std, err.Error())
		return dataResult, err
	}
	dataResult.Created.ID = id
	return dataResult, nil
}