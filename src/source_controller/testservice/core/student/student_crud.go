package student

import (
	"time"

	"configcenter/src/common"
	"configcenter/src/common/blog"
	"configcenter/src/common/mapstr"
	"configcenter/src/common/metadata"
	"configcenter/src/common/universalsql"
	"configcenter/src/source_controller/testservice/core"
)

func (m *studentManager) count(ctx core.ContextParams, cond universalsql.Condition) (uint64, error) {

	cnt, err := m.dbProxy.Table(common.BKTableNameStudent).Find(cond.ToMapStr()).Count(ctx)
	if nil != err {
		blog.Errorf("request(%s): it is failed to execute database count operation by the condition (%#v), error info is %s", ctx.ReqID, cond.ToMapStr(), err.Error())
		return 0, ctx.Error.Errorf(common.CCErrObjectDBOpErrno, err.Error())
	}

	return cnt, err
}

func (m *studentManager) save(ctx core.ContextParams, student *metadata.Student) (id uint64, err error) {

	id, err = m.dbProxy.NextSequence(ctx, common.BKTableNameStudent)
	if err != nil {
		blog.Errorf("request(%s): it is failed to make sequence id on the table (%s), error info is %s", ctx.ReqID, common.BKTableNameStudent, err.Error())
		return id, ctx.Error.New(common.CCErrObjectDBOpErrno, err.Error())
	}
	student.ID = int64(id)
	student.OwnerID = ctx.SupplierAccount

	if nil == student.LastTime {
		student.LastTime = &metadata.Time{}
		student.LastTime.Time = time.Now()
	}
	if nil == student.CreateTime {
		student.CreateTime = &metadata.Time{}
		student.CreateTime.Time = time.Now()
	}

	err = m.dbProxy.Table(common.BKTableNameStudent).Insert(ctx, student)
	return id, err
}

func (m *studentManager) update(ctx core.ContextParams, data mapstr.MapStr, cond universalsql.Condition) (cnt uint64, err error) {

	cnt, err = m.count(ctx, cond)
	if nil != err {
		return 0, err
	}

	if 0 == cnt {
		return 0, nil
	}

	data.Set(metadata.StudentFieldLastTime, time.Now())

	err = m.dbProxy.Table(common.BKTableNameStudent).Update(ctx, cond.ToMapStr(), data)
	if nil != err {
		blog.Errorf("request(%s): it is failed to execute database update operation on the table (%s), error info is %s", ctx.ReqID, common.BKTableNameStudent, err.Error())
		return 0, ctx.Error.New(common.CCErrObjectDBOpErrno, err.Error())
	}

	return cnt, err
}

func (m *studentManager) search(ctx core.ContextParams, cond universalsql.Condition) ([]metadata.Student, error) {

	dataResult := []metadata.Student{}
	if err := m.dbProxy.Table(common.BKTableNameStudent).Find(cond.ToMapStr()).All(ctx, &dataResult); nil != err {
		blog.Errorf("request(%s): it is failed to find all models by the condition (%#v), error info is %s", ctx.ReqID, cond.ToMapStr(), err.Error())
		return dataResult, ctx.Error.New(common.CCErrObjectDBOpErrno, err.Error())
	}

	return dataResult, nil
}

func (m *studentManager) searchReturnMapStr(ctx core.ContextParams, cond universalsql.Condition) ([]mapstr.MapStr, error) {

	dataResult := []mapstr.MapStr{}
	if err := m.dbProxy.Table(common.BKTableNameStudent).Find(cond.ToMapStr()).All(ctx, &dataResult); nil != err {
		blog.Errorf("request(%s): it is failed to find all models by the condition (%#v), error info is %s", ctx.ReqID, cond.ToMapStr(), err.Error())
		return dataResult, ctx.Error.New(common.CCErrObjectDBOpErrno, err.Error())
	}

	return dataResult, nil
}

func (m *studentManager) delete(ctx core.ContextParams, cond universalsql.Condition) (uint64, error) {

	cnt, err := m.count(ctx, cond)
	if nil != err {
		return 0, err
	}

	if 0 == cnt {
		return 0, nil
	}

	if err = m.dbProxy.Table(common.BKTableNameStudent).Delete(ctx, cond.ToMapStr()); nil != err {
		blog.Errorf("request(%s): it is failed to execute a deletion operation on the table (%s), error info is %s", ctx.ReqID, common.BKTableNameStudent, err.Error())
		return 0, ctx.Error.New(common.CCErrObjectDBOpErrno, err.Error())
	}

	return cnt, nil
}
