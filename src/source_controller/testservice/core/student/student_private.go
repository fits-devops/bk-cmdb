package student

import (
	"configcenter/src/common"
	"configcenter/src/common/blog"
	"configcenter/src/common/metadata"
	"configcenter/src/common/universalsql"
	"configcenter/src/common/universalsql/mongo"
	"configcenter/src/source_controller/testservice/core"
)

func (m *studentManager) isExists(ctx core.ContextParams, cond universalsql.Condition) (oneStudent *metadata.Student, exists bool, err error) {

	oneStudent = &metadata.Student{}
	err = m.dbProxy.Table(common.BKTableNameStudent).Find(cond.ToMapStr()).One(ctx, oneStudent)
	if nil != err && !m.dbProxy.IsNotFoundError(err) {
		blog.Errorf("request(%s): it is failed to execute database findone operation on the table (%#v) by the condition (%#v), error info is %s", ctx.ReqID, common.BKTableNameStudent, cond.ToMapStr(), err.Error())
		return oneStudent, exists, ctx.Error.New(common.CCErrObjectDBOpErrno, err.Error())
	}
	exists = !m.dbProxy.IsNotFoundError(err)
	return oneStudent, exists, nil
}

func (m *studentManager) isValid(ctx core.ContextParams, stuID string) error {

	checkCond := mongo.NewCondition()
	checkCond.Element(&mongo.Eq{Key: metadata.StudentFieldStudentID, Val: stuID})

	cnt, err := m.dbProxy.Table(common.BKTableNameStudent).Find(checkCond.ToMapStr()).Count(ctx)
	isValid := (0 != cnt)
	if nil != err {
		blog.Errorf("request(%s): it is failed to execute database cout operation on the table (%s) by the condition (%#v), error info is %s", ctx.ReqID, common.BKTableNameStudent, checkCond.ToMapStr(), err.Error())
		return ctx.Error.Error(common.CCErrObjectDBOpErrno)
	}

	if !isValid {
		return ctx.Error.Errorf(common.CCErrCommParamsIsInvalid, stuID)
	}

	return err
}
