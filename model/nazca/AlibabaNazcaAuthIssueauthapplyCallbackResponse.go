package nazca

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
出证申请回调 APIResponse
alibaba.nazca.auth.issueauthapply.callback

出证申请回调
*/
type AlibabaNazcaAuthIssueauthapplyCallbackAPIResponse struct {
    model.CommonResponse
    AlibabaNazcaAuthIssueauthapplyCallbackResponse
}

type AlibabaNazcaAuthIssueauthapplyCallbackResponse struct {
    XMLName xml.Name `xml:"alibaba_nazca_auth_issueauthapply_callback_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *ActionResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
