package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口查询同步订单 APIResponse
alibaba.wdk.syncedorder.query

外部商户查询同步到五道口的订单
*/
type AlibabaWdkSyncedorderQueryAPIResponse struct {
    model.CommonResponse
    AlibabaWdkSyncedorderQueryResponse
}

type AlibabaWdkSyncedorderQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_syncedorder_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 返回码
    
    ReturnCode   int64 `json:"return_code,omitempty" xml:"return_code,omitempty"`

    
    // 描述
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 订单号
    
    BizOrderId   string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`

    
}
