package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建OC订单 APIResponse
taobao.oc.order.create

创建OC订单接口
*/
type TaobaoOcOrderCreateAPIResponse struct {
    model.CommonResponse
    TaobaoOcOrderCreateResponse
}

type TaobaoOcOrderCreateResponse struct {
    XMLName xml.Name `xml:"oc_order_create_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 表示是否执行成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 执行失败原因
    
    ErrorDescription   string `json:"error_description,omitempty" xml:"error_description,omitempty"`

    
    // OcOrder
    
    OcOrder   *OcOrder `json:"oc_order,omitempty" xml:"oc_order,omitempty"`

    
}
