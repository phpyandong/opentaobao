package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
外贸直通车删除关键词 APIResponse
alibaba.scbp.ad.keyword.delete

外贸直通车删除关键词
*/
type AlibabaScbpAdKeywordDeleteAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdKeywordDeleteResponse
}

type AlibabaScbpAdKeywordDeleteResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_delete_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 删除关键词是否成功
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
