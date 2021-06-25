package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证查询状态接口 APIResponse
alibaba.security.jaq.rp.status

聚安全实人认证查询状态接口
*/
type AlibabaSecurityJaqRpStatusAPIResponse struct {
    model.CommonResponse
    Response *AlibabaSecurityJaqRpStatusResponse `json:"alibaba_security_jaq_rp_status_response,omitempty"`
}

type AlibabaSecurityJaqRpStatusResponse struct {

    // 状态信息
    Data   *RpStatusResultBo `json:"data,omitempty"`

}
