package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家类目删除接口 APIResponse
alibaba.wdk.sku.category.delete

商家类目删除接口
*/
type AlibabaWdkSkuCategoryDeleteAPIResponse struct {
    model.CommonResponse
    AlibabaWdkSkuCategoryDeleteResponse
}

type AlibabaWdkSkuCategoryDeleteResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_sku_category_delete_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用结果
    
    Result   *AlibabaWdkSkuCategoryDeleteApiResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
