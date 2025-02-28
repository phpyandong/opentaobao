package nazca

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
认证的统一回调接口 APIResponse
alibaba.nazca.auth.authapply.callback

认证的统一回调接口
*/
type AlibabaNazcaAuthAuthapplyCallbackAPIResponse struct {
    model.CommonResponse
    AlibabaNazcaAuthAuthapplyCallbackResponse
}

type AlibabaNazcaAuthAuthapplyCallbackResponse struct {
    XMLName xml.Name `xml:"alibaba_nazca_auth_authapply_callback_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *ActionResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
