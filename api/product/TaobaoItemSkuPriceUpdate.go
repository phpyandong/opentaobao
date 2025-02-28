package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
更新商品SKU的价格 
taobao.item.sku.price.update

更新商品SKU的价格
*/
func TaobaoItemSkuPriceUpdate(clt *core.SDKClient, req *product.TaobaoItemSkuPriceUpdateRequest, session string) (*product.TaobaoItemSkuPriceUpdateAPIResponse, error) {
    var resp product.TaobaoItemSkuPriceUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
