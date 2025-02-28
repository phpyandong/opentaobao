package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
考拉货品新增接口 APIResponse
taobao.kaola.scitem.add

考拉货品新增接口
*/
type TaobaoKaolaScitemAddAPIResponse struct {
    model.CommonResponse
    TaobaoKaolaScitemAddResponse
}

type TaobaoKaolaScitemAddResponse struct {
    XMLName xml.Name `xml:"kaola_scitem_add_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 异常信息
    
    ErrorMessages   []string `json:"error_messages,omitempty" xml:"error_messages>string,omitempty"`
    
    
    // 货品id
    
    Data   int64 `json:"data,omitempty" xml:"data,omitempty"`

    
    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 是否系统异常
    
    IsSystemFailed   bool `json:"is_system_failed,omitempty" xml:"is_system_failed,omitempty"`

    
    // 异常信息Code
    
    SysErrorCode   string `json:"sys_error_code,omitempty" xml:"sys_error_code,omitempty"`

    
}
