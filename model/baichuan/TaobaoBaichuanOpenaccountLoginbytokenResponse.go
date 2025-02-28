package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
百川TOKEN 登录 APIResponse
taobao.baichuan.openaccount.loginbytoken

百川TOKEN 登录
*/
type TaobaoBaichuanOpenaccountLoginbytokenAPIResponse struct {
    model.CommonResponse
    TaobaoBaichuanOpenaccountLoginbytokenResponse
}

type TaobaoBaichuanOpenaccountLoginbytokenResponse struct {
    XMLName xml.Name `xml:"baichuan_openaccount_loginbytoken_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // name
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`

    
}
