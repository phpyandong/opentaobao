package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
单个关键词效果报表 APIResponse
alibaba.scbp.effect.keyword.single.get

单个关键词效果报表
*/
type AlibabaScbpEffectKeywordSingleGetAPIResponse struct {
    model.CommonResponse
    AlibabaScbpEffectKeywordSingleGetResponse
}

type AlibabaScbpEffectKeywordSingleGetResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_effect_keyword_single_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 单个关键词报表
    
    KeywordEffectList   []SingleAdKeywordEffectDto `json:"keyword_effect_list,omitempty" xml:"keyword_effect_list>single_ad_keyword_effect_dto,omitempty"`
    
    
    // 总个数
    
    TotalNum   int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`

    
    // 总页数
    
    TotalPage   int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`

    
}
