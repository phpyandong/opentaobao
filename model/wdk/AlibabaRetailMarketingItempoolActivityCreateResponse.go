package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建商品池活动【同城零售】 APIResponse
alibaba.retail.marketing.itempool.activity.create

同城零售商品池活动创建
*/
type AlibabaRetailMarketingItempoolActivityCreateAPIResponse struct {
    model.CommonResponse
    AlibabaRetailMarketingItempoolActivityCreateResponse
}

type AlibabaRetailMarketingItempoolActivityCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_retail_marketing_itempool_activity_create_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 操作结果
    
    Result   *OctopusOpenResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
