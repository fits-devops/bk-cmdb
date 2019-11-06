/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package service

import (
	"strconv"
	"strings"

	"configcenter/src/common"
	"configcenter/src/common/blog"
	"configcenter/src/common/condition"
	"configcenter/src/common/mapstr"
	"configcenter/src/common/metadata"
	paraparse "configcenter/src/common/paraparse"
	"configcenter/src/scene_server/topo_server/core/operation"
	"configcenter/src/scene_server/topo_server/core/types"
)

// CreateInst create a new inst
func (s *Service) CreateInst(params types.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	objID := pathParams("obj_id")
	obj, err := s.Core.ObjectOperation().FindSingleObject(params, objID)
	if nil != err {
		blog.Errorf("failed to search the inst, %s", err.Error())
		return nil, err
	}
	// 处理提交过来需要关联的数据
	asstList := make(map[string]interface {})
	for index,val := range data {
		if index == "asst_id"{
			asstList["asst_data"] = val
			break
		}
	}
	if data.Exists("BatchInfo") {
		/*
		   BatchInfo data format:
		    {
		      "BatchInfo": {
		        "4": { // excel line number
		          "inst_id": 1,
		          "inst_key": "a22",
		          "inst_name": "a11",
		          "version": "121",
		          "import_from": "1"
		        },
		      "input_type": "excel"
		    }
		*/
		batchInfo := new(operation.InstBatchInfo)
		if err := data.MarshalJSONInto(batchInfo); err != nil {
			blog.Errorf("create instance failed, import object[%s] instance batch, but got invalid BatchInfo:[%v], err: %+v", objID, batchInfo, err)
			return nil, params.Err.Error(common.CCErrCommParamsIsInvalid)
		}

		setInst, err := s.Core.InstOperation().CreateInstBatch(params, obj, batchInfo)
		if nil != err {
			blog.Errorf("failed to create new object %s, %s", objID, err.Error())
			return nil, err
		}

		// auth register new created
		if len(setInst.SuccessCreated) != 0 {
			if err := s.AuthManager.RegisterInstancesByID(params.Context, params.Header, objID, setInst.SuccessCreated...); err != nil {
				blog.Errorf("create instance suceess, but register instances to iam failed, instances: %+v, err: %+v", setInst.SuccessCreated, err)
				return nil, params.Err.Error(common.CCErrCommRegistResourceToIAMFailed)
			}
		}

		// auth update registered instances
		if len(setInst.SuccessUpdated) == 0 {
			if err := s.AuthManager.UpdateRegisteredInstanceByID(params.Context, params.Header, objID, setInst.SuccessUpdated...); err != nil {
				blog.Errorf("update registered instances to iam failed, err: %+v", err)
				return nil, params.Err.Error(common.CCErrCommUnRegistResourceToIAMFailed)
			}
		}

		return setInst, nil
	}
	setInst, err := s.Core.InstOperation().CreateInst(params, obj, data)
	if nil != err {
		blog.Errorf("failed to create a new %s, %s", objID, err.Error())
		return nil, err
	}

	instanceID, err := setInst.GetInstID()
	if err != nil {
		blog.Errorf("create instance failed, unexpected error, create instance success, but get id failed, instance: %+v, err: %+v", setInst, err)
		return nil, err
	}
	// auth: register instances to iam
	if err := s.AuthManager.RegisterInstancesByID(params.Context, params.Header, objID, instanceID); err != nil {
		blog.Errorf("create instance success, but register instance to iam failed, instance: %d, err: %+v", instanceID, err)
		return nil, params.Err.Error(common.CCErrCommRegistResourceToIAMFailed)
	}
	// 创建关联关系 数据处理
	if  _, ok := asstList["asst_data"]; ok {
		associationIds := make([]int64,1)
		request := &metadata.CreateAssociationInstRequest{}
		for objAsstIdKey, value1 := range asstList["asst_data"].(map[string]interface{}) {
			for key2, idsValue := range value1.(map[string]interface{}) {
				if key2 == "id" {
					// id 可能存在多个
					for _,idString := range idsValue.([]interface{}) {
						asstData :=  mapstr.MapStr{}   //数据类型为map
						paramId := mapstr.MapStr{}
						paramId.Set("id", idString)
						id, err := paramId.Int64("id")
						if nil != err {
							blog.Errorf("[api-att] failed to parse the path params id(%s), error info is %s ", pathParams("id"), err.Error())
							return nil, err
						}


						asstData["inst_id"] = id
						asstData["asst_inst_id"] = instanceID
						asstData["obj_asst_id"] = objAsstIdKey

						if err := asstData.MarshalJSONInto(request); err != nil {
							return nil, params.Err.New(common.CCErrCommParamsInvalid, err.Error())
						}
						// 创建实例关联 交给 关联操作的方法去处理 ret.Data.ID =4
						ret, err := s.Core.AssociationOperation().CreateInst(params, request)
						if err != nil {
							// 删除关联的id
							for _,associationId:=range associationIds {
								if associationId > 0 {
									// 删除已关联的id
									s.Core.AssociationOperation().DeleteInst(params, associationId)
								}

							}
							// 删除实例ID
							s.Core.InstOperation().DeleteInstByInstID(params, obj, []int64{instanceID}, true)

							blog.Errorf("create instance association failed, do coreservice create failed, err: %+v, rid: %s", err, params.ReqID)
							return nil, err
						}
						associationIds = append(associationIds, ret.Data.ID)
					}
				}
			}
		}
	}

	return setInst.ToMapStr(), nil
}

func (s *Service) DeleteInsts(params types.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	obj, err := s.Core.ObjectOperation().FindSingleObject(params, pathParams("obj_id"))
	if nil != err {
		blog.Errorf("[api-inst] failed to find the objects(%s), error info is %s", pathParams("obj_id"), err.Error())
		return nil, err
	}

	deleteCondition := &operation.OpCondition{}
	if err := data.MarshalJSONInto(deleteCondition); nil != err {
		return nil, err
	}

	// auth: deregister resources
	if err := s.AuthManager.DeregisterInstanceByRawID(params.Context, params.Header, obj.GetObjectID(), deleteCondition.Delete.InstID...); err != nil {
		blog.Errorf("batch delete instance failed, deregister instance failed, instID: %d, err: %s", deleteCondition.Delete.InstID, err)
		return nil, params.Err.Error(common.CCErrCommUnRegistResourceToIAMFailed)
	}

	return nil, s.Core.InstOperation().DeleteInstByInstID(params, obj, deleteCondition.Delete.InstID, true)
}

// DeleteInst delete the inst
func (s *Service) DeleteInst(params types.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	if "batch" == pathParams("inst_id") {
		return s.DeleteInsts(params, pathParams, queryParams, data)
	}

	instID, err := strconv.ParseInt(pathParams("inst_id"), 10, 64)
	if nil != err {
		blog.Errorf("[api-inst]failed to parse the inst id, error info is %s", err.Error())
		return nil, params.Err.Errorf(common.CCErrCommParamsNeedInt, "inst id")
	}

	obj, err := s.Core.ObjectOperation().FindSingleObject(params, pathParams("obj_id"))
	if nil != err {
		blog.Errorf("[api-inst] failed to find the objects(%s), error info is %s", pathParams("obj_id"), err.Error())
		return nil, err
	}

	// auth: deregister resources
	if err := s.AuthManager.DeregisterInstanceByRawID(params.Context, params.Header, obj.GetObjectID(), instID); err != nil {
		blog.Errorf("delete instance failed, deregister instance failed, instID: %d, err: %s", instID, err)
		return nil, params.Err.Error(common.CCErrCommUnRegistResourceToIAMFailed)
	}

	err = s.Core.InstOperation().DeleteInstByInstID(params, obj, []int64{instID}, true)
	return nil, err
}

func (s *Service) UpdateInsts(params types.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {
	objID := pathParams("obj_id")

	updateCondition := &operation.OpCondition{}
	if err := data.MarshalJSONInto(updateCondition); nil != err {
		blog.Errorf("[api-inst] failed to parse the input data(%v), error info is %s", data, err.Error())
		return nil, err
	}

	// check inst_id field to be not empty, is dangerous for empty inst_id field, which will update or delete all instance
	for idx, item := range updateCondition.Update {
		if item.InstID == 0 {
			blog.Errorf("update instance failed, %d's update item's field `inst_id` emtpy", idx)
			return nil, params.Err.Error(common.CCErrCommParamsInvalid)
		}
	}
	for idx, instID := range updateCondition.Delete.InstID {
		if instID == 0 {
			blog.Errorf("update instance failed, %d's delete item's field `inst_id` emtpy", idx)
			return nil, params.Err.Error(common.CCErrCommParamsInvalid)
		}
	}

	obj, err := s.Core.ObjectOperation().FindSingleObject(params, objID)
	if nil != err {
		blog.Errorf("[api-inst] failed to find the objects(%s), error info is %s", pathParams("obj_id"), err.Error())
		return nil, err
	}

	instanceIDs := make([]int64, 0)
	for _, item := range updateCondition.Update {
		instanceIDs = append(instanceIDs, item.InstID)
		cond := condition.CreateCondition()
		cond.Field(obj.GetInstIDFieldName()).Eq(item.InstID)
		err = s.Core.InstOperation().UpdateInst(params, item.InstInfo, obj, cond, item.InstID)
		if nil != err {
			blog.Errorf("[api-inst] failed to update the object(%s) inst (%d),the data (%#v), error info is %s", obj.Object().ObjectID, item.InstID, data, err.Error())
			return nil, err
		}
	}

	// auth: deregister resources
	if err := s.AuthManager.UpdateRegisteredInstanceByID(params.Context, params.Header, objID, instanceIDs...); err != nil {
		blog.Errorf("update inst success, but update register to iam failed, instanceIDs: %+v, err: %+v", instanceIDs, err)
		return nil, params.Err.Error(common.CCErrCommRegistResourceToIAMFailed)
	}

	return nil, nil
}

// UpdateInst update the inst
func (s *Service) UpdateInst(params types.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	if "batch" == pathParams("inst_id") {
		return s.UpdateInsts(params, pathParams, queryParams, data)
	}

	objID := pathParams("obj_id")
	instID, err := strconv.ParseInt(pathParams("inst_id"), 10, 64)
	if nil != err {
		blog.Errorf("[api-inst]failed to parse the inst id, error info is %s", err.Error())
		return nil, params.Err.Errorf(common.CCErrCommParamsNeedInt, "inst id")
	}

	obj, err := s.Core.ObjectOperation().FindSingleObject(params, objID)
	if nil != err {
		blog.Errorf("[api-inst] failed to find the objects(%s), error info is %s", pathParams("obj_id"), err.Error())
		return nil, err
	}

	// this is a special logic for mainline object instance.
	// for auth reason, the front's request add metadata for mainline model's instance update.
	// but actually, it's should not add metadata field in the request.
	// so, we need remove it from the data if it's a mainline model instance.
	yes, err := s.Core.AssociationOperation().IsMainlineObject(params, objID)
	if err != nil {
		return nil, err
	}
	if yes {
		data.Remove("metadata")
	}

	cond := condition.CreateCondition()
	cond.Field(obj.GetInstIDFieldName()).Eq(instID)
	err = s.Core.InstOperation().UpdateInst(params, data, obj, cond, instID)
	if nil != err {
		blog.Errorf("[api-inst] failed to update the object(%s) inst (%s),the data (%#v), error info is %s", obj.Object().ObjectID, pathParams("inst_id"), data, err.Error())
		return nil, err
	}

	// auth: deregister resources
	if err := s.AuthManager.UpdateRegisteredInstanceByID(params.Context, params.Header, objID, instID); err != nil {
		blog.Error("update inst failed, authorization failed, instID: %d, err: %+v", instID, err)
		return nil, params.Err.Error(common.CCErrCommRegistResourceToIAMFailed)
	}

	return nil, err
}

// SearchInst search the inst
func (s *Service) SearchInsts(params types.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {
	objID := pathParams("obj_id")

	obj, err := s.Core.ObjectOperation().FindSingleObject(params, objID)
	if nil != err {
		blog.Errorf("[api-inst] failed to find the objects(%s), error info is %s", pathParams("obj_id"), err.Error())
		return nil, err
	}

	//	if nil != params.MetaData {
	//		data.Set(metadata.BKMetadata, *params.MetaData)
	//	}
	// construct the query inst condition
	queryCond := &paraparse.SearchParams{
		Condition: mapstr.New(),
	}
	if err := data.MarshalJSONInto(queryCond); nil != err {
		blog.Errorf("[api-inst] failed to parse the data and the condition, the input (%#v), error info is %s", data, err.Error())
		return nil, err
	}
	page := metadata.ParsePage(queryCond.Page)
	query := &metadata.QueryInput{}
	query.Condition = queryCond.Condition
	query.Fields = strings.Join(queryCond.Fields, ",")
	query.Limit = page.Limit
	query.Sort = page.Sort
	query.Start = page.Start

	cnt, instItems, err := s.Core.InstOperation().FindInst(params, obj, query, false)
	if nil != err {
		blog.Errorf("[api-inst] failed to find the objects(%s), error info is %s", pathParams("obj_id"), err.Error())
		return nil, err
	}

	result := mapstr.MapStr{}
	result.Set("count", cnt)
	result.Set("info", instItems)
	return result, nil
}

// SearchInstAndAssociationDetail search the inst with association details
func (s *Service) SearchInstAndAssociationDetail(params types.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {
	objID := pathParams("obj_id")
	obj, err := s.Core.ObjectOperation().FindSingleObject(params, objID)
	if nil != err {
		blog.Errorf("[api-inst] failed to find the objects(%s), error info is %s", pathParams("obj_id"), err.Error())
		return nil, err
	}

	// construct the query inst condition
	queryCond := &paraparse.SearchParams{
		Condition: mapstr.New(),
	}
	if err := data.MarshalJSONInto(queryCond); nil != err {
		blog.Errorf("[api-inst] failed to parse the data and the condition, the input (%#v), error info is %s", data, err.Error())
		return nil, err
	}
	page := metadata.ParsePage(queryCond.Page)
	query := &metadata.QueryInput{}
	query.Condition = queryCond.Condition
	query.Fields = strings.Join(queryCond.Fields, ",")
	query.Limit = page.Limit
	query.Sort = page.Sort
	query.Start = page.Start

	cnt, instItems, err := s.Core.InstOperation().FindInst(params, obj, query, true)
	if nil != err {
		blog.Errorf("[api-inst] failed to find the objects(%s), error info is %s", pathParams("obj_id"), err.Error())
		return nil, err
	}

	result := mapstr.MapStr{}
	result.Set("count", cnt)
	result.Set("info", instItems)
	return result, nil
}

// SearchInstByObject search the inst of the object
func (s *Service) SearchInstByObject(params types.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	objID := pathParams("obj_id")
	obj, err := s.Core.ObjectOperation().FindSingleObject(params, objID)
	if nil != err {
		blog.Errorf("[api-inst] failed to find the objects(%s), error info is %s", pathParams("obj_id"), err.Error())
		return nil, err
	}

	queryCond := &paraparse.SearchParams{
		Condition: mapstr.New(),
	}
	if err := data.MarshalJSONInto(queryCond); nil != err {
		blog.Errorf("[api-inst] failed to parse the data and the condition, the input (%#v), error info is %s", data, err.Error())
		return nil, err
	}
	page := metadata.ParsePage(queryCond.Page)
	query := &metadata.QueryInput{}
	query.Condition = queryCond.Condition
	query.Fields = strings.Join(queryCond.Fields, ",")
	query.Limit = page.Limit
	query.Sort = page.Sort
	query.Start = page.Start
	cnt, instItems, err := s.Core.InstOperation().FindInst(params, obj, query, false)
	if nil != err {
		blog.Errorf("[api-inst] failed to find the objects(%s), error info is %s", pathParams("obj_id"), err.Error())
		return nil, err
	}

	result := mapstr.MapStr{}
	result.Set("count", cnt)
	result.Set("info", instItems)
	return result, nil
}

// SearchInstByAssociation search inst by the association inst
func (s *Service) SearchInstByAssociation(params types.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	objID := pathParams("obj_id")

	obj, err := s.Core.ObjectOperation().FindSingleObject(params, objID)
	if nil != err {
		blog.Errorf("[api-inst] failed to find the objects(%s), error info is %s", pathParams("obj_id"), err.Error())
		return nil, err
	}

	cnt, instItems, err := s.Core.InstOperation().FindInstByAssociationInst(params, obj, data)
	if nil != err {
		blog.Errorf("[api-inst] failed to find the objects(%s), error info is %s", pathParams("obj_id"), err.Error())
		return nil, err
	}

	result := mapstr.MapStr{}
	result.Set("count", cnt)
	result.Set("info", instItems)
	return result, nil
}

// SearchInstByInstID search the inst by inst ID
func (s *Service) SearchInstByInstID(params types.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {
	objID := pathParams("obj_id")

	instID, err := strconv.ParseInt(pathParams("inst_id"), 10, 64)
	if nil != err {
		return nil, params.Err.New(common.CCErrTopoInstSelectFailed, err.Error())
	}

	obj, err := s.Core.ObjectOperation().FindSingleObject(params, objID)
	if nil != err {
		blog.Errorf("[api-inst] failed to find the objects(%s), error info is %s", pathParams("obj_id"), err.Error())
		return nil, err
	}

	cond := condition.CreateCondition()
	cond.Field(obj.GetInstIDFieldName()).Eq(instID)
	queryCond := &metadata.QueryInput{}
	queryCond.Condition = cond.ToMapStr()

	cnt, instItems, err := s.Core.InstOperation().FindInst(params, obj, queryCond, false)
	if nil != err {
		blog.Errorf("[api-inst] failed to find the objects(%s), error info is %s", pathParams("obj_id"), err.Error())
		return nil, err
	}

	result := mapstr.MapStr{}
	result.Set("count", cnt)
	result.Set("info", instItems)

	return result, nil
}

// SearchInstChildTopo search the child inst topo for a inst
func (s *Service) SearchInstChildTopo(params types.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {
	objID := pathParams("bk_object_id")

	instID, err := strconv.ParseInt(pathParams("inst_id"), 10, 64)
	if nil != err {
		return nil, err
	}

	obj, err := s.Core.ObjectOperation().FindSingleObject(params, objID)
	if nil != err {
		blog.Errorf("[api-inst] failed to find the objects(%s), error info is %s", pathParams("obj_id"), err.Error())
		return nil, err
	}

	query := &metadata.QueryInput{}
	cond := condition.CreateCondition()
	cond.Field(obj.GetInstIDFieldName()).Eq(instID)

	query.Condition = cond.ToMapStr()
	query.Limit = common.BKNoLimit

	_, instItems, err := s.Core.InstOperation().FindInstChildTopo(params, obj, instID, query)
	return instItems, err

}

// SearchInstTopo search the inst topo
func (s *Service) SearchInstTopo(params types.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	objID := pathParams("obj_id")
	instID, err := strconv.ParseInt(pathParams("inst_id"), 10, 64)
	if nil != err {
		blog.Errorf("search inst topo failed, path parameter inst_id invalid, inst_id: %s, err: %+v", pathParams("inst_id"), err)
		return nil, params.Err.Error(common.CCErrCommParamsIsInvalid)
	}

	obj, err := s.Core.ObjectOperation().FindSingleObject(params, objID)
	if nil != err {
		blog.Errorf("[api-inst] failed to find the objects(%s), error info is %s", pathParams("obj_id"), err.Error())
		return nil, err
	}

	query := &metadata.QueryInput{}
	cond := condition.CreateCondition()
	cond.Field(obj.GetInstIDFieldName()).Eq(instID)

	query.Condition = cond.ToMapStr()
	query.Limit = common.BKNoLimit

	_, instItems, err := s.Core.InstOperation().FindInstTopo(params, obj, instID, query)

	return instItems, err
}
