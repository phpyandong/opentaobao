package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取呼叫详情 APIResponse
alibaba.aliqin.fc.voice.getdetail

通过呼叫id获取呼叫相关的数据
*/
type AlibabaAliqinFcVoiceGetdetailAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinFcVoiceGetdetailResponse
}

type AlibabaAliqinFcVoiceGetdetailResponse struct {
    XMLName xml.Name `xml:"alibaba_aliqin_fc_voice_getdetail_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误码
    
    AlicomCode   string `json:"alicom_code,omitempty" xml:"alicom_code,omitempty"`

    
    // 错误信息
    
    AlicomMsg   string `json:"alicom_msg,omitempty" xml:"alicom_msg,omitempty"`

    
    // 请求是否成功
    
    AlicomSuccess   bool `json:"alicom_success,omitempty" xml:"alicom_success,omitempty"`

    
    // 返回值，在没有结果时为空。recordFile表示的是录音文件地址
    
    Model   string `json:"model,omitempty" xml:"model,omitempty"`

    
}
