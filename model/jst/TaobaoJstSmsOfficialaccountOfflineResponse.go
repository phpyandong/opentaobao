package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔公众号下线 APIResponse
taobao.jst.sms.officialaccount.offline

聚石塔公众号下线
*/
type TaobaoJstSmsOfficialaccountOfflineAPIResponse struct {
    model.CommonResponse
    TaobaoJstSmsOfficialaccountOfflineResponse
}

type TaobaoJstSmsOfficialaccountOfflineResponse struct {
    XMLName xml.Name `xml:"jst_sms_officialaccount_offline_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 系统异常
    
    ResponseCode   string `json:"response_code,omitempty" xml:"response_code,omitempty"`

    
    // 成功
    
    ResponseSuccess   bool `json:"response_success,omitempty" xml:"response_success,omitempty"`

    
    // 请求id
    
    ResponseId   string `json:"response_id,omitempty" xml:"response_id,omitempty"`

    
    // 成功
    
    Module   bool `json:"module,omitempty" xml:"module,omitempty"`

    
    // 系统异常
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
}
