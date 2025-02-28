package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
关键词批量改价 APIResponse
alibaba.scbp.ad.keyword.price.batchupdate

关键词批量改价
*/
type AlibabaScbpAdKeywordPriceBatchupdateAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdKeywordPriceBatchupdateResponse
}

type AlibabaScbpAdKeywordPriceBatchupdateResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_price_batchupdate_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 修改失败关键词列表
    
    KeywordErrorResultList   []KeywordErrorResultDto `json:"keyword_error_result_list,omitempty" xml:"keyword_error_result_list>keyword_error_result_dto,omitempty"`
    
    
}
