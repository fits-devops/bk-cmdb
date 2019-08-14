package student

import (
	"configcenter/src/common"
	"configcenter/src/common/blog"
	"configcenter/src/common/errors"
	"configcenter/src/common/mapstr"
	"configcenter/src/common/metadata"
	"configcenter/src/common/universalsql/mongo"
	"configcenter/src/common/util"
	"configcenter/src/source_controller/testservice/core"
	"configcenter/src/storage/dal"
)

//定义学生信息关联的班级信息操作数据结构
type classManager struct {
	std   *studentManager
	dbProxy dal.RDB
}

//实现创建一个班级信息的接口方法
func (m *classManager) CreateOneClass(ctx core.ContextParams, inputParam metadata.CreateOneClass) (*metadata.CreateOneDataResult, error) {

	//检查ClassID是否必传
	if 0 == len(inputParam.Data.ClassID) {
		blog.Errorf("request(%s): it is failed to create the class, because of the classID (%#v) is not set", ctx.ReqID, inputParam.Data)
		return &metadata.CreateOneDataResult{}, ctx.Error.Errorf(common.CCErrCommParamsNeedSet, metadata.ClassFieldClassID)
	}

	//检查传入的ClassID是否存在，在class_private.go中定义私有方法
	_, exists, err := m.isExists(ctx, inputParam.Data.ClassID, inputParam.Data.Metadata)
	if nil != err {
		blog.Errorf("request(%s): it is failed to check if the class ID (%s)is exists, error info is %s", ctx.ReqID, inputParam.Data.ClassID, err.Error())
		return nil, err
	}
	if exists {
		blog.Errorf("class (%#v)is duplicated", inputParam.Data)
		return nil, ctx.Error.Errorf(common.CCErrCommDuplicateItem, "")
	}

	inputParam.Data.OwnerID = ctx.SupplierAccount

	//开始插入班级信息数据, 在class_crud.go中定义CRUD(增删改查)方法
	id, err := m.save(ctx, inputParam.Data)
	if nil != err {
		blog.Errorf("request(%s): it is failed to save the class(%#v), error info is %s", ctx.ReqID, inputParam.Data, err.Error())
		return &metadata.CreateOneDataResult{}, err
	}
	return &metadata.CreateOneDataResult{Created: metadata.CreatedDataResult{ID: id}}, err
}

//实现批量创建班级信息的接口方法
func (m *classManager) CreateManyClass(ctx core.ContextParams, inputParam metadata.CreateManyClass) (*metadata.CreateManyDataResult, error) {

	dataResult := &metadata.CreateManyDataResult{
		CreateManyInfoResult: metadata.CreateManyInfoResult{
			Created:    []metadata.CreatedDataResult{},  //新增多少条
			Repeated:   []metadata.RepeatedDataResult{}, //存在多少条
			Exceptions: []metadata.ExceptionResult{},	//异常多少条
		},
	}

	addExceptionFunc := func(idx int64, err errors.CCErrorCoder, class *metadata.Class) {
		dataResult.CreateManyInfoResult.Exceptions = append(dataResult.CreateManyInfoResult.Exceptions, metadata.ExceptionResult{
			OriginIndex: idx,
			Message:     err.Error(),
			Code:        int64(err.GetCode()),
			Data:        class,
		})
	}

	//循环插入
	for itemIdx, item := range inputParam.Data {

		if 0 == len(item.ClassID) {
			blog.Errorf("request(%s): it is failed to create the class, because of the classID (%#v) is not set", ctx.ReqID, item.ClassID)
			addExceptionFunc(int64(itemIdx), ctx.Error.Errorf(common.CCErrCommParamsNeedSet, metadata.ClassFieldClassID).(errors.CCErrorCoder), &item)
			continue
		}

		_, exists, err := m.isExists(ctx, item.ClassID, item.Metadata)
		if nil != err {
			blog.Errorf("request(%s): it is failed to check the class ID (%s) is exists, error info is %s", ctx.ReqID, item.ClassID, err.Error())
			addExceptionFunc(int64(itemIdx), err.(errors.CCErrorCoder), &item)
			continue
		}

		if exists {
			dataResult.Repeated = append(dataResult.Repeated, metadata.RepeatedDataResult{OriginIndex: int64(itemIdx), Data: mapstr.NewFromStruct(item, "field")})
			continue
		}

		item.OwnerID = ctx.SupplierAccount
		id, err := m.save(ctx, item)
		if nil != err {
			blog.Errorf("request(%s): it is failed to save the clasisfication(%#v), error info is %s", ctx.ReqID, item, err.Error())
			addExceptionFunc(int64(itemIdx), err.(errors.CCErrorCoder), &item)
			continue
		}

		dataResult.Created = append(dataResult.Created, metadata.CreatedDataResult{
			ID: id,
		})

	}

	return dataResult, nil
}

//实现批量设置班级信息的接口方法
func (m *classManager) SetManyClass(ctx core.ContextParams, inputParam metadata.SetManyClass) (*metadata.SetDataResult, error) {

	dataResult := &metadata.SetDataResult{
		Created:    []metadata.CreatedDataResult{},
		Updated:    []metadata.UpdatedDataResult{},
		Exceptions: []metadata.ExceptionResult{},
	}

	addExceptionFunc := func(idx int64, err errors.CCErrorCoder, class *metadata.Class) {
		dataResult.Exceptions = append(dataResult.Exceptions, metadata.ExceptionResult{
			OriginIndex: idx,
			Message:     err.Error(),
			Code:        int64(err.GetCode()),
			Data:        class,
		})
	}

	for itemIdx, item := range inputParam.Data {

		if 0 == len(item.ClassID) {
			blog.Errorf("request(%s): it is failed to create the class, because of the classID (%#v) is not set", ctx.ReqID, item.ClassID)
			addExceptionFunc(int64(itemIdx), ctx.Error.Errorf(common.CCErrCommParamsNeedSet, metadata.ClassFieldClassID).(errors.CCErrorCoder), &item)
			continue
		}

		origin, exists, err := m.isExists(ctx, item.ClassID, item.Metadata)
		if nil != err {
			blog.Errorf("request(%s): it is failed to check the class ID (%s) is exists, error info is %s", ctx.ReqID, item.ClassID, err.Error())
			addExceptionFunc(int64(itemIdx), err.(errors.CCErrorCoder), &item)
			continue
		}

		if exists {

			cond := mongo.NewCondition()
			cond.Element(&mongo.Eq{Key: metadata.ClassFieldID, Val: origin.ID})
			if _, err := m.update(ctx, mapstr.NewFromStruct(item, "field"), cond); nil != err {
				blog.Errorf("request(%s): it is failed to update some fields(%#v) of the class by the condition(%#v), error info is %s", ctx.ReqID, item, cond.ToMapStr(), err.Error())
				addExceptionFunc(int64(itemIdx), err.(errors.CCErrorCoder), &item)
				continue
			}

			dataResult.UpdatedCount.Count++
			dataResult.Updated = append(dataResult.Updated, metadata.UpdatedDataResult{
				OriginIndex: int64(itemIdx),
				ID:          uint64(origin.ID),
			})
			continue
		}

		item.OwnerID = ctx.SupplierAccount

		id, err := m.save(ctx, item)
		if nil != err {
			blog.Errorf("request(%s): it is failed to save the class(%#v), error info is %s", ctx.ReqID, item, err.Error())
			addExceptionFunc(int64(itemIdx), err.(errors.CCErrorCoder), &item)
			continue
		}

		dataResult.CreatedCount.Count++
		dataResult.Created = append(dataResult.Created, metadata.CreatedDataResult{
			ID: id,
		})

	}

	return dataResult, nil
}

//实现单个设置班级信息的接口方法
func (m *classManager) SetOneClass(ctx core.ContextParams, inputParam metadata.SetOneClass) (*metadata.SetDataResult, error) {

	dataResult := &metadata.SetDataResult{
		Created:    []metadata.CreatedDataResult{},
		Updated:    []metadata.UpdatedDataResult{},
		Exceptions: []metadata.ExceptionResult{},
	}

	if 0 == len(inputParam.Data.ClassID) {
		blog.Errorf("request(%s): it is failed to set the class, because of the classID (%#v) is not set", ctx.ReqID, inputParam.Data)
		return dataResult, ctx.Error.Errorf(common.CCErrCommParamsNeedSet, metadata.ClassFieldClassID)
	}

	origin, exists, err := m.isExists(ctx, inputParam.Data.ClassID, inputParam.Data.Metadata)
	if nil != err {
		blog.Errorf("request(%s): it is failed to check the class ID (%s) is exists, error info is %s", ctx.ReqID, inputParam.Data.ClassID, err.Error())
		return dataResult, err
	}

	addExceptionFunc := func(idx int64, err errors.CCErrorCoder, class *metadata.Class) {
		dataResult.Exceptions = append(dataResult.Exceptions, metadata.ExceptionResult{
			OriginIndex: idx,
			Message:     err.Error(),
			Code:        int64(err.GetCode()),
			Data:        class,
		})
	}
	if exists {

		cond := mongo.NewCondition()
		cond.Element(&mongo.Eq{Key: metadata.ClassFieldID, Val: origin.ID})
		if _, err := m.update(ctx, mapstr.NewFromStruct(inputParam.Data, "field"), cond); nil != err {
			blog.Errorf("request(%s): it is failed to update some fields(%#v) for a class by the condition(%#v), error info is %s", ctx.ReqID, inputParam.Data, cond.ToMapStr(), err.Error())
			addExceptionFunc(0, err.(errors.CCErrorCoder), &inputParam.Data)
			return dataResult, nil
		}
		dataResult.UpdatedCount.Count++
		dataResult.Updated = append(dataResult.Updated, metadata.UpdatedDataResult{ID: uint64(origin.ID)})
		return dataResult, err
	}

	inputParam.Data.OwnerID = ctx.SupplierAccount
	id, err := m.save(ctx, inputParam.Data)
	if nil != err {
		blog.Errorf("request(%s): it is failed to save the class(%#v), error info is %s", ctx.ReqID, inputParam.Data, err.Error())
		addExceptionFunc(0, err.(errors.CCErrorCoder), origin)
		return dataResult, err
	}
	dataResult.CreatedCount.Count++
	dataResult.Created = append(dataResult.Created, metadata.CreatedDataResult{ID: id})
	return dataResult, err
}

//实现更新班级信息的接口方法
func (m *classManager) UpdateClass(ctx core.ContextParams, inputParam metadata.UpdateOption) (*metadata.UpdatedCount, error) {

	cond, err := mongo.NewConditionFromMapStr(util.SetModOwner(inputParam.Condition.ToMapInterface(), ctx.SupplierAccount))
	if nil != err {
		blog.Errorf("request(%s): it is failed to convert the condition(%#v) from mapstr into condition object, error info is %s", ctx.ReqID, inputParam.Condition, err.Error())
		return &metadata.UpdatedCount{}, err
	}

	cnt, err := m.update(ctx, inputParam.Data, cond)
	if nil != err {
		blog.Errorf("request(%s): it is failed to update some fields(%#v) for some class by the condition(%#v), error info is %s", ctx.ReqID, inputParam.Data, inputParam.Condition, err.Error())
		return &metadata.UpdatedCount{}, err
	}
	return &metadata.UpdatedCount{Count: cnt}, nil
}

//实现删除班级信息的方法
func (m *classManager) DeleteClass(ctx core.ContextParams, inputParam metadata.DeleteOption) (*metadata.DeletedCount, error) {

	deleteCond, err := mongo.NewConditionFromMapStr(util.SetModOwner(inputParam.Condition.ToMapInterface(), ctx.SupplierAccount))
	if nil != err {
		blog.Errorf("request(%s): it is failed to convert the condition (%#v) from mapstr into condition object, error info is %s", ctx.ReqID, inputParam.Condition, err.Error())
		return &metadata.DeletedCount{}, ctx.Error.New(common.CCErrCommHTTPInputInvalid, err.Error())
	}

	deleteCond.Element(&mongo.Eq{Key: metadata.ClassFieldClassOwnerID, Val: ctx.SupplierAccount})
	cnt, exists, err := m.hasStudent(ctx, deleteCond)
	if nil != err {
		blog.Errorf("request(%s): it is failed to check whether the class which are marked by the condition (%#v) have some models, error info is %s", ctx.ReqID, deleteCond.ToMapStr(), err.Error())
		return &metadata.DeletedCount{}, err
	}
	if exists {
		return &metadata.DeletedCount{}, ctx.Error.Error(common.CCErrTopoObjectClassificationHasObject)
	}

	cnt, err = m.delete(ctx, deleteCond)
	if nil != err {
		blog.Errorf("request(%s): it is failed to delete the class whci are marked by the condition(%#v), error info is %s", ctx.ReqID, deleteCond.ToMapStr(), err.Error())
		return &metadata.DeletedCount{}, err
	}

	return &metadata.DeletedCount{Count: cnt}, nil
}

//实现查询班级信息的接口方法
func (m *classManager) SearchClass(ctx core.ContextParams, inputParam metadata.QueryCondition) (*metadata.QueryClassDataResult, error) {

	dataResult := &metadata.QueryClassDataResult{
		Info: []metadata.Class{},
	}
	searchCond, err := mongo.NewConditionFromMapStr(util.SetQueryOwner(inputParam.Condition.ToMapInterface(), ctx.SupplierAccount))
	if nil != err {
		blog.Errorf("request(%s): it is failed to convert the condition (%#v) from mapstr into condition object, error info is %s", ctx.ReqID, inputParam.Condition, err.Error())
		return dataResult, err
	}

	totalCount, err := m.count(ctx, searchCond)
	if nil != err {
		blog.Errorf("request(%s): it is failed to get the count by the condition (%#v), error info is %s", ctx.ReqID, searchCond.ToMapStr(), err.Error())
		return dataResult, err
	}

	classItems, err := m.search(ctx, searchCond)
	if nil != err {
		blog.Errorf("request(%s): it is failed to search some class by the condition (%#v), error info is %s", ctx.ReqID, searchCond.ToMapStr(), err.Error())
		return dataResult, err
	}

	dataResult.Count = int64(totalCount)
	dataResult.Info = classItems
	return dataResult, nil
}
