package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发奖接口 APIResponse
alibaba.benefit.send

发奖接口
*/
type AlibabaBenefitSendAPIResponse struct {
    model.CommonResponse
    AlibabaBenefitSendResponse
}

type AlibabaBenefitSendResponse struct {
    XMLName xml.Name `xml:"alibaba_benefit_send_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回消息
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
    // 接口返回代码
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 是否处理成功
    
    ResultSuccess   bool `json:"result_success,omitempty" xml:"result_success,omitempty"`

    
    // 权益id
    
    RightId   int64 `json:"right_id,omitempty" xml:"right_id,omitempty"`

    
    // 奖品名称
    
    PrizeName   string `json:"prize_name,omitempty" xml:"prize_name,omitempty"`

    
}
