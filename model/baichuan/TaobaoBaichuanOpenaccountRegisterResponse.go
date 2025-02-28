package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
百川账号注册 APIResponse
taobao.baichuan.openaccount.register

百川账号注册
*/
type TaobaoBaichuanOpenaccountRegisterAPIResponse struct {
    model.CommonResponse
    TaobaoBaichuanOpenaccountRegisterResponse
}

type TaobaoBaichuanOpenaccountRegisterResponse struct {
    XMLName xml.Name `xml:"baichuan_openaccount_register_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // name
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`

    
}
