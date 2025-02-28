package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建结算调整单 APIResponse
tmall.service.settleadjustment.request

提供给服务商在对结算有异议时，发起结算调整单。
通过说明工单ID，调整费用值，调整原因进行新建结算调整单。
*/
type TmallServiceSettleadjustmentRequestAPIResponse struct {
    model.CommonResponse
    TmallServiceSettleadjustmentRequestResponse
}

type TmallServiceSettleadjustmentRequestResponse struct {
    XMLName xml.Name `xml:"tmall_service_settleadjustment_request_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TmallServiceSettleadjustmentRequestResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
