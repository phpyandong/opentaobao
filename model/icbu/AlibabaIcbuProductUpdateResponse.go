package icbu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改商品 APIResponse
alibaba.icbu.product.update

修改国际站商品，支持询盘商品和在线批发商品，支持英文商品和多语言商品
*/
type AlibabaIcbuProductUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaIcbuProductUpdateResponse
}

type AlibabaIcbuProductUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_icbu_product_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 加密后的产品ID
    
    ProductId   string `json:"product_id,omitempty" xml:"product_id,omitempty"`

    
}
