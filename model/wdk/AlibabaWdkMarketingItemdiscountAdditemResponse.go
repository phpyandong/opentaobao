package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
报名特价商品 APIResponse
alibaba.wdk.marketing.itemdiscount.additem

在商品特价活动中报名特价商品
*/
type AlibabaWdkMarketingItemdiscountAdditemAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingItemdiscountAdditemResponse
}

type AlibabaWdkMarketingItemdiscountAdditemResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_itemdiscount_additem_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品报名活动的返回结果
    
    Result   *MarketResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
