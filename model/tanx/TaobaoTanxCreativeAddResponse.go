package tanx

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创意预审新增接口 APIResponse
taobao.tanx.creative.add

创意预审新增接口
*/
type TaobaoTanxCreativeAddAPIResponse struct {
    model.CommonResponse
    TaobaoTanxCreativeAddResponse
}

type TaobaoTanxCreativeAddResponse struct {
    XMLName xml.Name `xml:"tanx_creative_add_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用的成功信息或失败信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 调用返回码
    
    Code   int64 `json:"code,omitempty" xml:"code,omitempty"`

    
    // 是否成功
    
    IsOk   bool `json:"is_ok,omitempty" xml:"is_ok,omitempty"`

    
}
