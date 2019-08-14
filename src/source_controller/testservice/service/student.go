package service

import (
	"configcenter/src/common/mapstr"
	"configcenter/src/common/metadata"
	"configcenter/src/source_controller/testservice/core"
)

// Class班级信息表的Handler方法
func (s *testService) CreateManyClass(params core.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	inputDatas := metadata.CreateManyClass{}
	if err := data.MarshalJSONInto(&inputDatas); nil != err {
		return nil, err
	}
	return s.core.StudentOperation().CreateManyClass(params, inputDatas)
}

func (s *testService) CreateOneClass(params core.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	inputData := metadata.CreateOneClass{}
	if err := data.MarshalJSONInto(&inputData.Data); nil != err {
		return nil, err
	}
	return s.core.StudentOperation().CreateOneClass(params, inputData)
}

func (s *testService) SetOneClass(params core.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	inputData := metadata.SetOneClass{}
	if err := data.MarshalJSONInto(&inputData); nil != err {
		return nil, err
	}

	return s.core.StudentOperation().SetOneClass(params, inputData)
}

func (s *testService) SetManyClass(params core.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	inputDatas := metadata.SetManyClass{}
	if err := data.MarshalJSONInto(&inputDatas); nil != err {
		return nil, err
	}
	return s.core.StudentOperation().SetManyClass(params, inputDatas)
}

func (s *testService) UpdateClass(params core.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	inputData := metadata.UpdateOption{}
	if err := data.MarshalJSONInto(&inputData); nil != err {
		return nil, err
	}
	return s.core.StudentOperation().UpdateClass(params, inputData)
}

func (s *testService) DeleteClass(params core.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	inputData := metadata.DeleteOption{}
	if err := data.MarshalJSONInto(&inputData); nil != err {
		return nil, err
	}
	return s.core.StudentOperation().DeleteClass(params, inputData)
}

func (s *testService) SearchClass(params core.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	inputData := metadata.QueryCondition{}
	if err := data.MarshalJSONInto(&inputData); nil != err {
		return nil, err
	}

	dataResult, err := s.core.StudentOperation().SearchClass(params, inputData)
	if nil != err {
		return dataResult, err
	}

	// translate language
	for index := range dataResult.Info {
		dataResult.Info[index].ClassName = s.TranslateClassName(params.Lang, &dataResult.Info[index])
	}

	return dataResult, err
}

// Student学生信息表的Handler方法
func (s *testService) CreateStudent(params core.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	inputData := metadata.CreateStudent{}
	if err := data.MarshalJSONInto(&inputData); nil != err {
		return nil, err
	}
	return s.core.StudentOperation().CreateStudent(params, inputData)
}