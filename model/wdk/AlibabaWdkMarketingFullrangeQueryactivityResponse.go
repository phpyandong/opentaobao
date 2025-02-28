package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
全场活动查询活动 APIResponse
alibaba.wdk.marketing.fullrange.queryactivity

全场活动查询活动
*/
type AlibabaWdkMarketingFullrangeQueryactivityAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingFullrangeQueryactivityResponse
}

type AlibabaWdkMarketingFullrangeQueryactivityResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_fullrange_queryactivity_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 查询返回结果
    
    Result   *MarketResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
