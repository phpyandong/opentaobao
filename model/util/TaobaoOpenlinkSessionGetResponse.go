package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取授权session信息 APIResponse
taobao.openlink.session.get

帮助第三方isv生成三方session
*/
type TaobaoOpenlinkSessionGetAPIResponse struct {
    model.CommonResponse
    TaobaoOpenlinkSessionGetResponse
}

type TaobaoOpenlinkSessionGetResponse struct {
    XMLName xml.Name `xml:"openlink_session_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TaobaoOpenlinkSessionGetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
