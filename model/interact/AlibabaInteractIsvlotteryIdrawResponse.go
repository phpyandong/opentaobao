package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
互动到店抽奖 APIResponse
alibaba.interact.isvlottery.idraw

互动到店抽奖
*/
type AlibabaInteractIsvlotteryIdrawAPIResponse struct {
    model.CommonResponse
    AlibabaInteractIsvlotteryIdrawResponse
}

type AlibabaInteractIsvlotteryIdrawResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_isvlottery_idraw_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 抽奖中奖信息
    
    Data   *LotteryProxyResult `json:"data,omitempty" xml:"data,omitempty"`

    
    // 是否调用成功
    
    Succ   bool `json:"succ,omitempty" xml:"succ,omitempty"`

    
    // 错误信息
    
    BizCode   string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`

    
    // 错误信息描述
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`

    
}
