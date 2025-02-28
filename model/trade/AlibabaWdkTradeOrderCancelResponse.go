package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
外部交易订单取消接口 APIResponse
alibaba.wdk.trade.order.cancel

通过该接口可以再盒马取消交易订单，并处理相关业务流程。主要用于和外部商户的订单进行同步和融合业务流程处理
*/
type AlibabaWdkTradeOrderCancelAPIResponse struct {
    model.CommonResponse
    AlibabaWdkTradeOrderCancelResponse
}

type AlibabaWdkTradeOrderCancelResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_trade_order_cancel_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 执行结果
    
    Result   *OrderResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
