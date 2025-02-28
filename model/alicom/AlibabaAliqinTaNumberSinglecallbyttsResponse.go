package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据号码tts单呼 APIResponse
alibaba.aliqin.ta.number.singlecallbytts

将语音验证码和语音通知发布至聚石塔渠道
*/
type AlibabaAliqinTaNumberSinglecallbyttsAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinTaNumberSinglecallbyttsResponse
}

type AlibabaAliqinTaNumberSinglecallbyttsResponse struct {
    XMLName xml.Name `xml:"alibaba_aliqin_ta_number_singlecallbytts_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *AlibabaAliqinTaNumberSinglecallbyttsResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
