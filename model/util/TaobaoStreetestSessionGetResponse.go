package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据获取压测用户的sessionKey APIResponse
taobao.streetest.session.get

根据正常用户sessionKey获取对应压测用户的sessionKey，该sessionKey只能用户服务商全链路压测
*/
type TaobaoStreetestSessionGetAPIResponse struct {
    model.CommonResponse
    TaobaoStreetestSessionGetResponse
}

type TaobaoStreetestSessionGetResponse struct {
    XMLName xml.Name `xml:"streetest_session_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 压测账号对应的sessionKey
    
    StreeTestSessionKey   string `json:"stree_test_session_key,omitempty" xml:"stree_test_session_key,omitempty"`

    
}
