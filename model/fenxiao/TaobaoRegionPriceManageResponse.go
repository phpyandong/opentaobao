package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
编辑区域价格 APIResponse
taobao.region.price.manage

编辑区域价格
*/
type TaobaoRegionPriceManageAPIResponse struct {
    model.CommonResponse
    TaobaoRegionPriceManageResponse
}

type TaobaoRegionPriceManageResponse struct {
    XMLName xml.Name `xml:"region_price_manage_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // success
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
