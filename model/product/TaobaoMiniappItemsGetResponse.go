package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
批量获取商品信息 APIResponse
taobao.miniapp.items.get

获取商品公开属性，只允许在商家应用环境中使用
*/
type TaobaoMiniappItemsGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoMiniappItemsGetResponse `json:"taobao_miniapp_items_get_response,omitempty"`
}

type TaobaoMiniappItemsGetResponse struct {

    // Item(商品)结构
    Items   []Item `json:"items,omitempty"`

}
