package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品池活动新增商品 APIResponse
alibaba.retail.marketing.itempool.activity.sku.add

新增或更新商品池活动商品信息【同城零售】
*/
type AlibabaRetailMarketingItempoolActivitySkuAddAPIResponse struct {
    model.CommonResponse
    AlibabaRetailMarketingItempoolActivitySkuAddResponse
}

type AlibabaRetailMarketingItempoolActivitySkuAddResponse struct {
    XMLName xml.Name `xml:"alibaba_retail_marketing_itempool_activity_sku_add_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 出参
    
    Result   *OctopusOpenResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
