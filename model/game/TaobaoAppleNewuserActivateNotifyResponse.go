package game

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新用户激活通知接口 APIResponse
taobao.apple.newuser.activate.notify

资和信主动通知激活结果
*/
type TaobaoAppleNewuserActivateNotifyAPIResponse struct {
    model.CommonResponse
    TaobaoAppleNewuserActivateNotifyResponse
}

type TaobaoAppleNewuserActivateNotifyResponse struct {
    XMLName xml.Name `xml:"apple_newuser_activate_notify_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 处理结果说明
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
    // 处理结果码
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
}
