package tmallhk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
上传小票数据 APIResponse
tmall.traceplatform.ticket.order.upload

upsertOrderBySeller
*/
type TmallTraceplatformTicketOrderUploadAPIResponse struct {
    model.CommonResponse
    TmallTraceplatformTicketOrderUploadResponse
}

type TmallTraceplatformTicketOrderUploadResponse struct {
    XMLName xml.Name `xml:"tmall_traceplatform_ticket_order_upload_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回值
    
    Result   *DataResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
