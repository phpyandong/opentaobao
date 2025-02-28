package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新增档期商品 APIResponse
alibaba.price.promotion.item.add

批量新增档期活动商品
*/
type AlibabaPricePromotionItemAddAPIResponse struct {
    model.CommonResponse
    AlibabaPricePromotionItemAddResponse
}

type AlibabaPricePromotionItemAddResponse struct {
    XMLName xml.Name `xml:"alibaba_price_promotion_item_add_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Result   *AlibabaPricePromotionItemAddResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
