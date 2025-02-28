package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
百川用户名密码登录 APIResponse
taobao.baichuan.openaccount.login

百川用户名密码登录
*/
type TaobaoBaichuanOpenaccountLoginAPIResponse struct {
    model.CommonResponse
    TaobaoBaichuanOpenaccountLoginResponse
}

type TaobaoBaichuanOpenaccountLoginResponse struct {
    XMLName xml.Name `xml:"baichuan_openaccount_login_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // name
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`

    
}
