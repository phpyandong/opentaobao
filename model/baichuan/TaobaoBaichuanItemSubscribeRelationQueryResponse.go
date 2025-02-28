package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询单个订阅关系 APIResponse
taobao.baichuan.item.subscribe.relation.query

查询单个订阅关系
*/
type TaobaoBaichuanItemSubscribeRelationQueryAPIResponse struct {
    model.CommonResponse
    TaobaoBaichuanItemSubscribeRelationQueryResponse
}

type TaobaoBaichuanItemSubscribeRelationQueryResponse struct {
    XMLName xml.Name `xml:"baichuan_item_subscribe_relation_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *TaobaoBaichuanItemSubscribeRelationQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
