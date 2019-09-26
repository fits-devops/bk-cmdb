### 创建分组基本信息
- API: POST /api/{version}/objectatt/group/new
- API 名称：create_group
- 功能说明: 
	- 中文: 创建分组
	- English: create a group for a object

- input body

``` json

{
    "group_id":"3jbvwqbhq75",
    "group_name":"未命名",
    "group_index":5,
    "obj_id":"process",
    "org_id":"0"
}

```

**注:以上 JSON 数据中各字段的取值仅为示例数据。**

-  intput字段说明

| 字段|类型|必填|默认值|说明|Description|
|---|---|---|---|---| ---|
|group_id|string| 是|无|分组ID，纯英文字符序列，不允许修改|the group identifier|
|group_name|string|是|无|分组名字，用于展示|the group name|
|group_index| int|是|0|分组排序|the group index|
|obj_id|string|是|无|模型ID，用于指明该分组的所属|the object identifier|
|org_id|string|是|无|开发商账号|supplier account code|


- output

``` json
{
	"result": true,
	"error_code": 0,
	"error_msg": null,
	"data": {
		"id": 1046
	}
}
```

**注:以上 JSON 数据中各字段的取值仅为示例数据。**

- output字段说明

| 名称  | 类型  | 说明 |Description|
|---|---|---|---|
| result | bool | 请求成功与否。true:请求成功；false请求失败 |request result|
| error_code | int | 错误编码。 0表示success，>0表示失败错误 |error code. 0 represent success, >0 represent failure code |
| error_msg | string | 请求失败返回的错误信息 |error message from failed request|
| data | object| 请求返回的数据 |return data|

data 说明

| 名称  | 类型  | 说明 |Description|
|---|---|---|---|
|id|int|新增数据记录的ID|the id of the new data record|

### 查询分组基本信息
- API: POST /api/{version}/objectatt/group/property/owner/{org_id}/object/{obj_id}
- API 名称：search_group
- 功能说明：
	- 中文：查询分组信息
	- English：query the grouping of models

- input body

	无

- input 字段说明

| 名称  | 类型 |必填| 默认值 | 说明 | Description|
| ---  | ---  | --- |---  | --- | ---|
|obj_id|string|是|无|模型ID|the object id|
|org_id|string|是|无|开发商账号|supplier account code|

- output


``` json
{
    "result": true,
    "code": 0,
    "message": null,
    "data": [
        {
            "group_id": "default",
            "group_index": 1,
            "group_name": "基础信息",
            "isdefault": true,
            "obj_id": "host",
            "org_id": "0",
            "id": 5,
            "ispre": false
        }
    ]
}
```

**注:以上 JSON 数据中各字段的取值仅为示例数据。**

- output字段说明

| 名称  | 类型  | 说明 |Description|
|---|---|---|---|
| result | bool | 请求成功与否。true:请求成功；false请求失败 |request result|
| error_code | int | 错误编码。 0表示success，>0表示失败错误 |error code. 0 represent success, >0 represent failure code |
| error_msg | string | 请求失败返回的错误信息 |error message from failed request|
| data | array| 请求返回的数据 |return data|

data 说明

| 名称  | 类型  | 说明 |Description|
|---|---|---|---|
|group_id|string|分组标识|the group identifier|
|group_index|int|分组排序|the group sort index|
|group_name|string|分组名|the group name|
|isdefault|bool|true-默认分组,false-普通分组|true - the defualt group, false - the common group|
|obj_id|string|模型标识|the object identifier|
|org_id|string|开发商账号|supplier account code|
|id|int|数据记录ID|the data record id|
|ispre|bool|true - 内置分组, false - 自定义定义分组|true - is inner, false - is customer|


### 修改分组基本信息
- API: PUT  /api/{version}/objectatt/group/update
- API 名称：update_group
- input body
``` json
{
    "condition":{
        "id":10
    },
    "data":{
        "group_name":"GSEkit 进程管理信息2"
    }
}
```

**注:以上 JSON 数据中各字段的取值仅为示例数据。**

- input参数说明

| 字段|类型|必填|默认值|说明|Description|
|---|---|---|---|---|---|
|condition|object|是|无|更新分组信息的条件|the condition to update|
|data|object|是|无|更新数据至|update the data to|


condition 字段说明：

| 字段|类型|必填|默认值|说明|Description|
|---|---|---|---|---|---|
|id|int|否|无|分组id|the group id|

data 字段说明：

|字段|类型|必填|默认值|说明|Description|
|---|---|---|---|---|---|
|group_name|string|否|无|分组名字，用于展示|the group name|
|group_index| int|否|0|分组排序|the group index|


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

- output字段说明

| 名称  | 类型  | 说明 |Description|
|---|---|---|---|
| result | bool | 请求成功与否。true:请求成功；false请求失败 |request result|
| error_code | int | 错误编码。 0表示success，>0表示失败错误 |error code. 0 represent success, >0 represent failure code |
| error_msg | string | 请求失败返回的错误信息 |error message from failed request|
| data | string| 操作结果 |the result|

### 删除分组基本信息
- API: DELETE /api/{version}/objectatt/group/groupid/{id}
- API 名称：delete_group
- 功能说明：
	- 中文：删除模型分组
	- English：delete the group of the object

- input body 

	无

- input 字段说明

| 字段|类型|必填|默认值|说明|Description|
|---|---|---|---|---|---|
|id|int|是|无|分组记录标识|the group record id|


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

- output字段说明

| 名称  | 类型  | 说明 |Description|
|---|---|---|---|
| result | bool | 请求成功与否。true:请求成功；false请求失败 |request result|
| error_code | int | 错误编码。 0表示success，>0表示失败错误 |error code. 0 represent success, >0 represent failure code |
| error_msg | string | 请求失败返回的错误信息 |error message from failed request|
| data | string| 操作结果 |the result|




### 删除模型属性分组
- API: DELETE  /api/{version}/objectatt/group/owner/{org_id}/object/{obj_id}/propertyids/{property_id}/groupids/{group_id}
- API 名称：delete_object_property_group
- 功能说明：
	- 中文：删除模型属性分组
	- English：delete the group of the object property

- input body

	无

- input 字段说明

| 字段|类型|必填|默认值|说明|Description|
|---|---|---|---|---|---|
|group_id|string|是|无|分组ID|the group record  id|
|property_id|string|是|无|属性ID|the property identifier|
|obj_id|string|是|无|模型ID|the object identifier|
|org_id|string|是|无|开发商账号|supplier account code|


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

- output字段说明

| 名称  | 类型  | 说明 |Description|
|---|---|---|---|
| result | bool | 请求成功与否。true:请求成功；false请求失败 |request result|
| error_code | int | 错误编码。 0表示success，>0表示失败错误 |error code. 0 represent success, >0 represent failure code |
| error_msg | string | 请求失败返回的错误信息 |error message from failed request|
| data | string| 操作结果 |the result|