package student

/**
class操作的私有方法
 */

import (
	"configcenter/src/common"
	"configcenter/src/common/blog"
	"configcenter/src/common/metadata"
	"configcenter/src/common/universalsql"
	"configcenter/src/common/universalsql/mongo"
	"configcenter/src/source_controller/testservice/core"
)

//通过班级ID校验是否可用
func (m *classManager) isValid(ctx core.ContextParams, classID string) (bool, error) {

	cond := mongo.NewCondition()
	cond.Element(&mongo.Eq{Key: metadata.ClassFieldClassID, Val: classID})
	cond.Element(&mongo.Eq{Key: metadata.ClassFieldClassOwnerID, Val: ctx.SupplierAccount})

	cnt, err := m.count(ctx, cond)
	return 0 != cnt, err
}

//通过班级ID校验数据库是否存在班级信息
func (m *classManager) isExists(ctx core.ContextParams, classID string, meta metadata.Metadata) (origin *metadata.Class, exists bool, err error) {

	origin = &metadata.Class{}
	cond := mongo.NewCondition()
	cond.Element(&mongo.Eq{Key: metadata.ClassFieldClassOwnerID, Val: ctx.SupplierAccount})
	cond.Element(&mongo.Eq{Key: metadata.ClassFieldClassID, Val: classID})

	err = m.dbProxy.Table(common.BKTableNameStudentClass).Find(cond.ToMapStr()).One(ctx, origin)
	if nil != err && !m.dbProxy.IsNotFoundError(err) {
		return origin, false, err
	}
	return origin, !m.dbProxy.IsNotFoundError(err), nil
}

//查询是否存在该学生信息
func (m *classManager) hasStudent(ctx core.ContextParams, cond universalsql.Condition) (cnt uint64, exists bool, err error) {

	clsItems, err := m.search(ctx, cond)
	if nil != err {
		return 0, false, err
	}

	clsIDS := []string{}
	for _, item := range clsItems {
		clsIDS = append(clsIDS, item.ClassID)
	}

	checkModelCond := mongo.NewCondition()
	checkModelCond.Element(mongo.Field(metadata.StudentFieldClassID).In(clsIDS))
	checkModelCond.Element(mongo.Field(metadata.StudentFieldOwnerID).Eq(ctx.SupplierAccount))

	cnt, err = m.dbProxy.Table(common.BKTableNameStudent).Find(checkModelCond.ToMapStr()).Count(ctx)
	if nil != err {
		blog.Errorf("request(%s): it is failed to execute database count operation on the table(%s) by the condition(%#v), error info is %s", ctx.ReqID, common.BKTableNameStudent, cond.ToMapStr(), err.Error())
		return 0, false, err
	}
	exists = 0 != cnt
	return cnt, exists, err
}
