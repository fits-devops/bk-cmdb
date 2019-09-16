# 实例关联

### 查询实例关联

* API：POST / api / {version} / find / instassociation
* API名称：search_inst_association
* 功能说明：
  * 中文：查询实例之间的关联信息
  * 英语：搜索inst之间的关联
* 输入体

```
{
    "condition": {
        "obj_asst_id": "bk_switch_belong_bk_host",
        "asst_id": "",
        "bk_object_id": "",
        "asst_obj_id": ""
    },
    "metadata":{
        "label":{
            "biz_id":"1"
        }
    }
}
```
* 输入字段说明

|字段名|类型|必填|说明|
| ---  | ---  | --- |---  |
|obj_asst_id|串|否|关联唯一标识|
|asst_id|串|否|关联类型|
|bk_object_id|串|否|源模型ID|
|asst_obj_id|串|否|目标实例ID|

* 产量
```
{
    "result": true,
    "error_code": 0,
    "error_msg": null,
    "data": [{
        "obj_asst_id": "",
        "obj_id":"",
        "asst_obj_id":"",
        "inst_id":0,
        "asst_inst_id":0,
        "org_id":""
    }]
}
```
注：以上JSON数据中各字段的取值仅为示例数据。

* 输出字段说明

|字段|类型|说明|描述|
| ---  | ---  | --- |---  |
|结果|布尔|ture：成功，假：失败|true：成功，错误：失败|
|error_code|INT|错误编码.0表示成功，> 0表示失败错误|错误代码。0表示成功，> 0表示失败代码|
|error_msg|串|请求失败返回的错误信息|失败请求的错误消息|
|数据|宾语|结果数据|结果|

* data说明（结构待定）

|字段|类型|说明|描述|
| ---  | ---  | --- |---  |
|obj_asst_id|串|模型关联唯一标识|对象关联唯一标识|
|obj_id|串|源模型ID，冗余字段|源对象ID|
|asst_obj_id|串|目标模型ID|目标对象ID|
|inst_id|INT|源实例ID|source inst id|
|asst_inst_id|INT|目标实例ID|target inst id|
|org_id|串|开发商账号|供应商帐户代码|

### 添加实例关联

* API：POST / api / {version} / create / instassociation
* API名称：create_inst_association
* 功能说明：
  * 中文：添加实例关联
  * 英语：在inst之间创建一个关联
* 输入体
```
{
    "obj_asst_id": "bk_switch_belong_bk_host",
    "inst_id": 1,
    "asst_inst_id": 2,
    "metadata":{
        "label":{
            "biz_id":"1"
        }
    }
}
```
* 输入字段说明

|字段名|类型|必填|说明|
| ---  | ---  | --- |---  |
|obj_asst_id|串|是|唯一标识|
|inst_id|INT|是|源实例ID|
|asst_inst_id|INT|是|目标实例ID|


* 产量

```
{
    "result": true,
    "error_code": 0,
    "error_msg": null,
    "data": {
        "id": 1038
    }
}
```
* 输入字段说明

|字段|类型|说明|描述|
| ---  | ---  | --- |---  |
|结果|布尔|ture：成功，假：失败|true：成功，错误：失败|
|error_code|INT|错误编码.0表示成功，> 0表示失败错误|错误代码。0表示成功，> 0表示失败代码|
|error_msg|串|请求失败返回的错误信息|失败请求的错误消息|
|数据|宾语|操作结果|结果|

* 数据字段说明

|字段|类型|说明|描述|
| ---  | ---  | --- |---  |
|ID|INT|自增ID|自动递增ID|

### 删除实例关联

* API：DELETE / api / {version} / delete / instassociation / {id}
* API名称：delete_inst_association
* 功能说明：
  * 中文：删除实例关联
  * 英语：删除inst之间的关联
* association_id：实联关联关系的自增id值。
* 产量
```
{
    "result": true,
    "error_code": 0,
    "error_msg": null,
    "data": "success"
}
```
* 输出字段说明

|字段|类型|说明|描述|
| ---  | ---  | --- |---  |
|结果|布尔|ture：成功，假：失败|true：成功，错误：失败|
|error_code|INT|错误编码.0表示成功，> 0表示失败错误|错误代码。0表示成功，> 0表示失败代码|
|error_msg|串|请求失败返回的错误信息|失败请求的错误消息|
|数据|宾语|操作结果|结果|
