package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
退款 APIResponse
alibaba.mos.onsite.trade.refund

当交易发生之后一段时间内，由于消费者或者商户的原因需退款，商户可通过退款接口将支付款退还给消费者，喵街将在收到退款请求并验证成功后，按退款规则将支付款按原路退到消费者账号上。

1. 交易超过可退款时间（签约时设置的可退款时间）的订单无法进行退款。
2. 只支持全额退款。
*/
type AlibabaMosOnsiteTradeRefundAPIResponse struct {
    model.CommonResponse
    AlibabaMosOnsiteTradeRefundResponse
}

type AlibabaMosOnsiteTradeRefundResponse struct {
    XMLName xml.Name `xml:"alibaba_mos_onsite_trade_refund_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 交易退款响应
    
    Result   *AlibabaMosOnsiteTradeRefundResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
