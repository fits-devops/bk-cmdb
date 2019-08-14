package student

import (
	"configcenter/src/common"
	"configcenter/src/common/blog"
	"configcenter/src/common/mapstr"
	"configcenter/src/common/metadata"
	"configcenter/src/common/universalsql"
	"configcenter/src/source_controller/testservice/core"
)

//统计班级信息
func (m *classManager) count(ctx core.ContextParams, cond universalsql.Condition) (cnt uint64, err error) {

	cnt, err = m.dbProxy.Table(common.BKTableNameStudentClass).Find(cond.ToMapStr()).Count(ctx)
	if nil != err {
		blog.Errorf("request(%s): it is failed to execute a database count operation on the table(%s) by the condition(%#v), error info is %s", ctx.ReqID, common.BKTableNameStudentClass, cond.ToMapStr(), err.Error())
		return 0, err
	}
	return cnt, err
}

//插入班级信息
func (m *classManager) save(ctx core.ContextParams, class metadata.Class) (id uint64, err error) {

	id, err = m.dbProxy.NextSequence(ctx, common.BKTableNameStudentClass)
	if nil != err {
		blog.Errorf("request(%s): it is failed to create a new sequence id on the table(%s) of the database, error info is %s", ctx.ReqID, common.BKTableNameStudentClass, err.Error())
		return id, ctx.Error.New(common.CCErrObjectDBOpErrno, err.Error())
	}

	class.ID = int64(id)
	class.OwnerID = ctx.SupplierAccount

	err = m.dbProxy.Table(common.BKTableNameStudentClass).Insert(ctx, class)
	return id, err
}

//更新班级信息
func (m *classManager) update(ctx core.ContextParams, data mapstr.MapStr, cond universalsql.Condition) (cnt uint64, err error) {

	cnt, err = m.count(ctx, cond)
	if nil != err {
		return cnt, err
	}

	if 0 == cnt {
		return cnt, nil
	}

	data.Remove(metadata.ClassFieldClassID)
	err = m.dbProxy.Table(common.BKTableNameStudentClass).Update(ctx, cond.ToMapStr(), data)
	if nil != err {
		blog.Errorf("request(%s): it is failed to execute a database update operation on the table(%s) by the condition(%#v) , error info is %s", ctx.ReqID, common.BKTableNameStudentClass, cond.ToMapStr(), err.Error())
		return 0, err
	}
	return cnt, err
}

//删除班级信息
func (m *classManager) delete(ctx core.ContextParams, cond universalsql.Condition) (cnt uint64, err error) {

	cnt, err = m.count(ctx, cond)
	if nil != err {
		return cnt, err
	}

	if 0 == cnt {
		return 0, err
	}

	err = m.dbProxy.Table(common.BKTableNameStudentClass).Delete(ctx, cond.ToMapStr())
	if nil != err {
		blog.Errorf("request(%s): it is failed to execute a database deletion operation on the table(%s) by the condition(%#v), error info is %s", ctx.ReqID, common.BKTableNameStudentClass, cond.ToMapStr(), err.Error())
		return 0, err
	}

	return cnt, err
}

//查询班级信息
func (m *classManager) search(ctx core.ContextParams, cond universalsql.Condition) ([]metadata.Class, error) {

	results := []metadata.Class{}
	err := m.dbProxy.Table(common.BKTableNameStudentClass).Find(cond.ToMapStr()).All(ctx, &results)
	return results, err
}

//查询班级信息并返回mapstr数据结构
func (m *classManager) searchReturnMapStr(ctx core.ContextParams, cond universalsql.Condition) ([]mapstr.MapStr, error) {

	results := []mapstr.MapStr{}
	err := m.dbProxy.Table(common.BKTableNameStudentClass).Find(cond.ToMapStr()).All(ctx, &results)
	return results, err
}
