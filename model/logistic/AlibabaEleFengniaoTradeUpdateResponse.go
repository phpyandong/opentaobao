package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新蜂鸟扣费状态 APIResponse
alibaba.ele.fengniao.trade.update

汇金扣费成功后，回调该接口更新扣费状态
*/
type AlibabaEleFengniaoTradeUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaEleFengniaoTradeUpdateResponse
}

type AlibabaEleFengniaoTradeUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_ele_fengniao_trade_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 1:成功 0：失败
    
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`

    
    // 无此交易记录
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`

    
}
