package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
百川新登录二次验证 APIResponse
taobao.baichuan.openaccount.newlogindoublecheck

百川新登录二次验证
*/
type TaobaoBaichuanOpenaccountNewlogindoublecheckAPIResponse struct {
    model.CommonResponse
    TaobaoBaichuanOpenaccountNewlogindoublecheckResponse
}

type TaobaoBaichuanOpenaccountNewlogindoublecheckResponse struct {
    XMLName xml.Name `xml:"baichuan_openaccount_newlogindoublecheck_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // name
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`

    
}
