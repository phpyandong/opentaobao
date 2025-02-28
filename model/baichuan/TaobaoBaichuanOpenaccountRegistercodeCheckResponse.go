package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
百川检查注册验证码 APIResponse
taobao.baichuan.openaccount.registercode.check

百川检查注册验证码
*/
type TaobaoBaichuanOpenaccountRegistercodeCheckAPIResponse struct {
    model.CommonResponse
    TaobaoBaichuanOpenaccountRegistercodeCheckResponse
}

type TaobaoBaichuanOpenaccountRegistercodeCheckResponse struct {
    XMLName xml.Name `xml:"baichuan_openaccount_registercode_check_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // name
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`

    
}
