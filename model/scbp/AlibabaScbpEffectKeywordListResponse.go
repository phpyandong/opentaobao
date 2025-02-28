package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
关键词报表 APIResponse
alibaba.scbp.effect.keyword.list

关键词报表
*/
type AlibabaScbpEffectKeywordListAPIResponse struct {
    model.CommonResponse
    AlibabaScbpEffectKeywordListResponse
}

type AlibabaScbpEffectKeywordListResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_effect_keyword_list_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 关键词效果数据列表
    
    KeywordReportList   []AdKeywordEffectDto `json:"keyword_report_list,omitempty" xml:"keyword_report_list>ad_keyword_effect_dto,omitempty"`
    
    
    // 总个数
    
    TotalNum   int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`

    
    // 总页数
    
    TotalPage   int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`

    
}
