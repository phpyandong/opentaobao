package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
刷新Access Token APIResponse
taobao.top.auth.token.refresh

根据refresh_token重新生成token，目前只有服务市场订购类应用可以刷新token，其他类型应用（如商家后台）使用固定时长token，不提供刷新功能。
*/
type TaobaoTopAuthTokenRefreshAPIResponse struct {
    model.CommonResponse
    TaobaoTopAuthTokenRefreshResponse
}

type TaobaoTopAuthTokenRefreshResponse struct {
    XMLName xml.Name `xml:"top_auth_token_refresh_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回的是json信息
    
    TokenResult   string `json:"token_result,omitempty" xml:"token_result,omitempty"`

    
}
