package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发票回调接口 APIResponse
taobao.bus.invoice.return

汽车票发票回调接口
*/
type TaobaoBusInvoiceReturnAPIResponse struct {
    model.CommonResponse
    TaobaoBusInvoiceReturnResponse
}

type TaobaoBusInvoiceReturnResponse struct {
    XMLName xml.Name `xml:"bus_invoice_return_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误码
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 错误信息
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
