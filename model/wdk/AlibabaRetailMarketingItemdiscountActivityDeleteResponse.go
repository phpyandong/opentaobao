package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除单品特价活动【同城零售】 APIResponse
alibaba.retail.marketing.itemdiscount.activity.delete

同城零售单品特价活动删除
*/
type AlibabaRetailMarketingItemdiscountActivityDeleteAPIResponse struct {
    model.CommonResponse
    AlibabaRetailMarketingItemdiscountActivityDeleteResponse
}

type AlibabaRetailMarketingItemdiscountActivityDeleteResponse struct {
    XMLName xml.Name `xml:"alibaba_retail_marketing_itemdiscount_activity_delete_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 操作结果
    
    Result   *OctopusOpenResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
