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

package middleware

import (
	"context"
	"fmt"
	"net/http"

	bkc "configcenter/src/common"
	"configcenter/src/common/backbone"

	"gopkg.in/redis.v5"
)

type Discover struct {
	ctx     context.Context
	pheader http.Header

	redisCli *redis.Client
	*backbone.Engine
}

var msgHandlerCnt = int64(0)

func NewDiscover(ctx context.Context, redisCli *redis.Client, backbone *backbone.Engine) *Discover {
	pheader := http.Header{}
	pheader.Add(bkc.BKHTTPOwnerID, bkc.BKDefaultOwnerID)
	pheader.Add(bkc.BKHTTPHeaderUser, bkc.CCSystemCollectorUserName)

	discover := &Discover{
		redisCli: redisCli,
		ctx:      ctx,
		pheader:  pheader,
	}
	discover.Engine = backbone
	return discover
}

func (d *Discover) Analyze(msg string) error {
	err := d.TryCreateModel(msg)
	if err != nil {
		return fmt.Errorf("create model err: %v, raw: %s", err, msg)
	}

	err = d.UpdateOrAppendAttrs(msg)
	if err != nil {
		return fmt.Errorf("create property err: %v, raw: %s", err, msg)
	}

	err = d.UpdateOrCreateInst(msg)
	if err != nil {
		return fmt.Errorf("create inst err: %v, raw: %s", err, msg)
	}
	return nil
}

var MockMessage = `{
    "host": {
        "host_id": 1,
        "org_id": "0"
    },
    "meta": {
        "model": {
            "classification_id": "middelware",
            "obj_id": "bk_apache",
            "obj_name": "apache",
            "org_id": "0"
        },
        "fields": {
            "inst_name":{
                "property_name": "实例名",
                "property_type":"longchar"
            },
            "bk_ip":{
                "property_name":"IP",
                "property_type": "longchar"
            }
        }
    },
    "data": {
        "inst_name": "apache",
        "bk_ip": "192.168.0.1"
    }
}`
