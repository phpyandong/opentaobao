package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
校验支付密码 APIResponse
alibaba.alsc.crm.customer.checkppw

校验支付密码
*/
type AlibabaAlscCrmCustomerCheckppwAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmCustomerCheckppwResponse
}

type AlibabaAlscCrmCustomerCheckppwResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_customer_checkppw_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
