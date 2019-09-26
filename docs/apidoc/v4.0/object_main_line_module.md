
### 添加模型主关联
- API POST /api/{version}/topo/model/mainline
- API 名称：create_mainline_object
- 功能说明：
	- 中文：添加主线模型
	- English：create the main line model

- input body

``` json
{
	"classification_id": "XXX",
	"obj_id": "cc_test",
	"obj_name": "cc_test",
	"org_id": "0",
	"asst_obj_id": "id-XXX",
	"obj_icon": "icon-XXX"
}
```

**注:以上 JSON 数据中各字段的取值仅为示例数据。**

- input 字段说明

- 输入参数

|字段|类型|必填|默认值|说明|Description|
|---|---|---|---|---|---|
|classification_id|string|是|无|对象模型的分类ID，只能用英文字母序列命名|the classification identifier|
| obj_id |string|是|无|对象模型的ID，只能用英文字母序列命名|the object identifier|
| obj_name |string|是|无|对象模型的名字，用于展示，可以使用人类可以阅读的任何语言|the object name|
|org_id|string|是|无|开发商账号|supplier account code|
| asst_obj_id |string|是|无|主线模型关联的父对象模型的ID（obj_id）|the association object identifier|
| obj_icon|string|是|无|模型的图标|the icon of the object|

- output

``` json
{
	"result": true,
	"error_code": 0,
	"error_msg": null,
	"data": "success"
}
```

**注:以上 JSON 数据中各字段的取值仅为示例数据。**

- output 字段说明

| 名称  | 类型  | 说明 |Description|
|---|---|---|---|
| result | bool | 请求成功与否。true:请求成功；false请求失败 |request result true or false|
| error_code | int | 错误编码。 0表示success，>0表示失败错误 |error code. 0 represent success, >0 represent failure code |
| error_msg | string | 请求失败返回的错误信息 |error message from failed request|
| data | string | 请求返回的数据 |the data response|

### 删除模型主关联

- API: DELETE  /api/{version}/topo/model/mainline/owners/{org_id}/objectids/{obj_id}
- API 名称：delete_mainline_object
- 功能说明：
	- 中文：删除主线模型
	- English：delete the mainline object

- input body

    无


- input 字段说明

| 字段|类型|必填|默认值|说明|Description|
|---|---|---|---|---|---|
|org_id|string|是|无|开发商账号|supplier account code|
|obj_id|string|是|无|对象模型的ID|the object identifier|


- output

``` json
{
	"result": true,
	"error_code": 0,
	"error_msg": null,
	"data": "success"
}
```

**注:以上 JSON 数据中各字段的取值仅为示例数据。**

- output 字段说明

| 名称  | 类型  | 说明 |Description|
|---|---|---|---|
| result | bool | 请求成功与否。true:请求成功；false请求失败 |request result true or false|
| error_code | int | 错误编码。 0表示success，>0表示失败错误 |error code. 0 represent success, >0 represent failure code |
| error_msg | string | 请求失败返回的错误信息 |error message from failed request|
| data | string | 请求返回的数据 |the data response|

### 查询模型拓扑

- API: GET/api/{version}/topo/model/{org_id}  
- API 名称：search_mainline_object
- 功能说明：
	- 中文：搜索主线模型
	- English：search the main line model

- input body

    无

-  input字段说明

|字段|类型|必填|默认值|说明|Description|
|---|---|---|---|---|---|
|org_id|string|是|无|开发商账号|supplier account code|


- output

``` json
{
	"result": true,
	"error_code": 0,
	"error_msg": null,
	"data": [{
		"bk_next_name": "",
		"bk_next_obj": "",
		"obj_id": "biz",
		"obj_name": "业务",
		"bk_pre_obj_id": "",
		"bk_pre_obj_name": "",
		"org_id": "0"
	}]
}
```

**注:以上 JSON 数据中各字段的取值仅为示例数据。**

- output 字段说明

| 名称  | 类型  | 说明 |Description|
|---|---|---|---|
| result | bool | 请求成功与否。true:请求成功；false请求失败 |request result true or false|
| error_code | int | 错误编码。 0表示success，>0表示失败错误 |error code. 0 represent success, >0 represent failure code |
| error_msg | string | 请求失败返回的错误信息 |error message from failed request|
| data | array | 请求返回的数据 |the data response|

data 字段说明：

| 名称  | 类型  | 说明 |Description|
|---|---|---|---|
|bk_next_name|string|下一个模型的名字|the next object name|
|bk_next_obj|string|下一个模型的ID|the next object identifier|
|obj_id|string|当前的模型ID|the current object identifier|
|obj_name|string|当前模型的名字|the current object name|
|bk_pre_obj_id|string|上一个模型的ID|the pre object identifier|
|bk_pre_obj_name|string|上一个模型的名字|the pre object name|
|org_id|string|开发商账号|supplier account code|



### 获取实例拓扑

- API: GET /api/{version}/topo/inst/{org_id}/{biz_id}
- API 名称：get_inst_topo
- 功能说明：
	- 中文：获取实例拓扑
	- English：get the  topo of the inst
	
- input body

    无


- input 输入参数

| 字段|类型|必填|默认值|说明|Description|
|---|---|---|---|---|---|
|biz_id|int|是|无|业务id|the business id|
|org_id|string|是|无|开发商账号|supplier account code|


- output

``` json
{
	"result": true,
	"error_code": 0,
	"error_msg": null,
	"data": [{
		"default": 0,
		"inst_id": 96,
		"inst_name": "cc_biz_test",
		"obj_id": "biz",
		"obj_name": "业务",
		"child": [{
			"default": 0,
			"inst_id": 58,
			"inst_name": "obj_id_name",
			"obj_id": "obj_id",
			"obj_name": "obj_id_name",
			"child": [{
				"default": 0,
				"inst_id": 59,
				"inst_name": "obj_inst_name",
				"obj_id": "obj_inst",
				"obj_name": "obj_inst",
				"child": []
			}]
		}]
	}]
}
```

**注:以上 JSON 数据中各字段的取值仅为示例数据。**

- output 字段说明

| 名称  | 类型  | 说明 |Description|
|---|---|---|---|
| result | bool | 请求成功与否。true:成功；false:失败 |request result true or false|
| error_code | int | 错误编码。 0表示success，>0表示失败错误 |error code. 0 represent success, >0 represent failure code |
| error_msg | string | 请求失败返回的错误信息 |error message from failed request|
| data | array | 请求返回的数据 |the data response|

data 字段说明：

| 名称  | 类型  | 说明 |Description|
|---|---|---|---|
|inst_id|int|实例ID|the inst identifier|
|inst_name|string|实例名字|the inst name|
|obj_id|string|模型的标识|the object identifier|
|obj_name|string|模型名|the object name|
|child|array|实例集合|the inst array|

**注:child节点下包含的字段于data节点包含的字段一致。**

###  获取子节点实例

- API: GET /api/{version}/topo/inst/child/{org_id}/{obj_id}/{biz_id}/{inst_id}
- API名称：search_inst_topo
- 功能说明：
	- 中文：获取子节点实例拓扑
	- English：search inst topo

- input body

    无

- input 输入参数

| 字段|类型|必填|默认值|Description|
|---|---|---|---|---|
|biz_id|int|是|无|业务id|the business id|
|org_id|string|是|无|开发商账号|supplier account code|
|obj_id|string|是|无|对象模型的ID|the object identifier|
|inst_id|string|是|无|实例ID|the inst id|

- output

``` json
{
	"result": true,
	"error_code": 0,
	"error_msg": null,
	"data": [{
		"default": 0,
		"inst_id": 96,
		"inst_name": "cc_biz_test",
		"obj_id": "biz",
		"obj_name": "业务",
		"child": [{
			"default": 0,
			"inst_id": 58,
			"inst_name": "obj_id_name",
			"obj_id": "obj_id",
			"obj_name": "obj_id_name",
			"child": [{
				"default": 0,
				"inst_id": 59,
				"inst_name": "obj_inst_name",
				"obj_id": "obj_inst",
				"obj_name": "obj_inst",
				"child": []
			}]
		}]
	}]
}
```

**注:以上 JSON 数据中各字段的取值仅为示例数据。**

- output 字段说明

| 名称  | 类型  | 说明 |Description|
|---|---|---|---|
| result | bool | 请求成功与否。true:请求成功；false请求失败 |request result true or false|
| error_code | int | 错误编码。 0表示success，>0表示失败错误 |error code. 0 represent success, >0 represent failure code |
| error_msg | string | 请求失败返回的错误信息 |error message from failed request|
| data | array | 请求返回的数据 |the data response|

data 字段说明：

| 名称  | 类型  | 说明 |Description|
|---|---|---|---|
|default|int|1-资源模块（空闲机），2-故障模块（故障机）|1-Resource Module(Idle Machine),2-Fault Module(Fault Machine)|
|inst_id|int|实例ID|the inst identifier|
|inst_name|string|实例名字|the inst name|
|obj_id|string|模型的标识|the object identifier|
|obj_name|string|模型名|the object name|
|child|array|实例集合|the inst array|

**注:child节点下包含的字段于data节点包含的字段一致。**

###  查询内置模块集
- API: GET /api/{version}/topo/internal/{org_id}/{biz_id}
- API名称： get_internal_topo
- 功能说明：
	- 中文：获取业务的空闲机和故障机模块
	- English：get the internal idle-cluster and the fault-cluster


- input body

    无


- input 字段说明

| 字段|类型|必填|默认值|说明Description|
|---|---|---|---|---|
|org_id|string|是|无|开发商账号|supplier account code|
|biz_id|int|是|无|业务ID|the business id|


- output
```
{
    "result":true,
    "error_code":0,
    "error_msg":null,
    "data":{
        “module":[
            {
                “module_id":503,
                “module_name":"空闲机"
            },
            {
                “module_id":504,
                “module_name":"故障机"
            }
        ],
        “set_id":214,
        “set_name":"内置模块集"
    }
}
```

**注:以上 JSON 数据中各字段的取值仅为示例数据。**

- output 字段说明

| 名称  | 类型  | 说明 |Description|
|---|---|---|---|
| result | bool | 请求成功与否。true:请求成功；false请求失败 |request result true or false|
| error_code | int | 错误编码。 0表示success，>0表示失败错误 |error code. 0 represent success, >0 represent failure code |
| error_msg | string | 请求失败返回的错误信息 |error message from failed request|
| data | object | 请求返回的数据 |the data response|

data 字段说明:

| 名称  | 类型  | 说明 |Description| 
|---|---|---|---|
|set_id|int|集群ID|the set id|
|set_name|string|集群名字|the set name|

module 字段说明:

| 名称  | 类型  | 说明 |Description| 
|---|---|---|---|
| module_id|int|模块记录ID|the module data record id|
|module_name|string|模块名|the module name|