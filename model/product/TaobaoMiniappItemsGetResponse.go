package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量获取商品信息 APIResponse
taobao.miniapp.items.get

获取商品公开属性，只允许在商家应用环境中使用
*/
type TaobaoMiniappItemsGetAPIResponse struct {
    model.CommonResponse
    TaobaoMiniappItemsGetResponse
}

type TaobaoMiniappItemsGetResponse struct {
    XMLName xml.Name `xml:"miniapp_items_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // Item(商品)结构
    
    Items   []Item `json:"items,omitempty" xml:"items>item,omitempty"`
    
    
}
