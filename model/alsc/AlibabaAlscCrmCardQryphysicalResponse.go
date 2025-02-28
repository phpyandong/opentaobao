package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询物理卡 APIResponse
alibaba.alsc.crm.card.qryphysical

查询物理卡
*/
type AlibabaAlscCrmCardQryphysicalAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmCardQryphysicalResponse
}

type AlibabaAlscCrmCardQryphysicalResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_card_qryphysical_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
