package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家类目新增接口 APIResponse
alibaba.wdk.sku.category.add

商家类目新增接口
*/
type AlibabaWdkSkuCategoryAddAPIResponse struct {
    model.CommonResponse
    AlibabaWdkSkuCategoryAddResponse
}

type AlibabaWdkSkuCategoryAddResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_sku_category_add_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用结果
    
    Result   *AlibabaWdkSkuCategoryAddApiResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
