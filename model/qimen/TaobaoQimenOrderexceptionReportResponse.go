package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单异常通知接口 APIResponse
taobao.qimen.orderexception.report

WMS调用奇门的接口,当WMS接收到ERP的发货指令时，由于种种原因（5.1.5说明了各种异常场景）可能无法完成发货。WMS通过调用此接口，通知ERP具体异常情况
*/
type TaobaoQimenOrderexceptionReportAPIResponse struct {
    model.CommonResponse
    TaobaoQimenOrderexceptionReportResponse
}

type TaobaoQimenOrderexceptionReportResponse struct {
    XMLName xml.Name `xml:"qimen_orderexception_report_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 
    
    Response   *Response `json:"response,omitempty" xml:"response,omitempty"`

    
}
