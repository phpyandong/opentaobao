package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
TMC授权token APIResponse
taobao.tmc.auth.get

TMC连接授权Token
*/
type TaobaoTmcAuthGetAPIResponse struct {
    model.CommonResponse
    TaobaoTmcAuthGetResponse
}

type TaobaoTmcAuthGetResponse struct {
    XMLName xml.Name `xml:"tmc_auth_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
