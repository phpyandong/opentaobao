package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台奖池查询接口 APIResponse
alibaba.marketing.lottery.activity.query

抽奖平台奖池查询接口
*/
type AlibabaMarketingLotteryActivityQueryAPIResponse struct {
    model.CommonResponse
    AlibabaMarketingLotteryActivityQueryResponse
}

type AlibabaMarketingLotteryActivityQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_marketing_lottery_activity_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 分页结果
    
    PagingDto   *PagingDto `json:"paging_dto,omitempty" xml:"paging_dto,omitempty"`

    
    // code
    
    MsgCode   int64 `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
    // success
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // msg
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
}
