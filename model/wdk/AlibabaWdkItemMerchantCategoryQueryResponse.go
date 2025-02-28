package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商品的商家叶子类目 APIResponse
alibaba.wdk.item.merchant.category.query

查询商品的商家叶子类目
*/
type AlibabaWdkItemMerchantCategoryQueryAPIResponse struct {
    model.CommonResponse
    AlibabaWdkItemMerchantCategoryQueryResponse
}

type AlibabaWdkItemMerchantCategoryQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_item_merchant_category_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   *WdkOpenSkuMerchantCatServiceQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
