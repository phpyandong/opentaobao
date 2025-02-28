package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
账户-报表 APIResponse
alibaba.scbp.effect.account.list

账户-报表,支持最近7天，最近30天，以及180天内时间区间。
*/
type AlibabaScbpEffectAccountListAPIResponse struct {
    model.CommonResponse
    AlibabaScbpEffectAccountListResponse
}

type AlibabaScbpEffectAccountListResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_effect_account_list_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 账户效果数据列表
    
    AccountReportList   []AccountEffectDto `json:"account_report_list,omitempty" xml:"account_report_list>account_effect_dto,omitempty"`
    
    
    // 总个数
    
    TotalNum   int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`

    
    // 总页数
    
    TotalPage   int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`

    
}
