package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
判断支付宝用户是否绑定淘宝账号 APIResponse
alibaba.aliqin.flow.alipay.isbindingtbaccount

判断支付宝用户是否绑定淘宝账号
*/
type AlibabaAliqinFlowAlipayIsbindingtbaccountAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinFlowAlipayIsbindingtbaccountResponse
}

type AlibabaAliqinFlowAlipayIsbindingtbaccountResponse struct {
    XMLName xml.Name `xml:"alibaba_aliqin_flow_alipay_isbindingtbaccount_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // error
    
    Error   bool `json:"error,omitempty" xml:"error,omitempty"`

    
    // TRUE代表绑定，FALSE代表未绑定
    
    Value   string `json:"value,omitempty" xml:"value,omitempty"`

    
    // errorCode
    
    AlicomFlowErrorCode   string `json:"alicom_flow_error_code,omitempty" xml:"alicom_flow_error_code,omitempty"`

    
    // errorMsg
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`

    
}
