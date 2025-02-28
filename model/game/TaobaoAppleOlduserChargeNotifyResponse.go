package game

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
老用户激活并兑换通知接口 APIResponse
taobao.apple.olduser.charge.notify

老用户激活并兑换通知接口
*/
type TaobaoAppleOlduserChargeNotifyAPIResponse struct {
    model.CommonResponse
    TaobaoAppleOlduserChargeNotifyResponse
}

type TaobaoAppleOlduserChargeNotifyResponse struct {
    XMLName xml.Name `xml:"apple_olduser_charge_notify_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 处理结果说明
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
    // 处理结果码
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
}
