package wlbimports

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
一般进口取消物流订单 APIResponse
taobao.wlb.imports.order.cancel

商家在发货后，需要对订单进行取消，如果仓库已经收货则无法取消。
*/
type TaobaoWlbImportsOrderCancelAPIResponse struct {
    model.CommonResponse
    TaobaoWlbImportsOrderCancelResponse
}

type TaobaoWlbImportsOrderCancelResponse struct {
    XMLName xml.Name `xml:"wlb_imports_order_cancel_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 业务错误描述
    
    ResultErrorMsg   string `json:"result_error_msg,omitempty" xml:"result_error_msg,omitempty"`

    
    // 是否取消订单成功，true：成功，false：失败
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 业务错误编码
    
    ResultErrorCode   string `json:"result_error_code,omitempty" xml:"result_error_code,omitempty"`

    
}
