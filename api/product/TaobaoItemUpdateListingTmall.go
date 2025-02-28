package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
taobao.item.update.listing天猫分流 
taobao.item.update.listing.tmall

* 单个商品上架<br/>* 输入的num_iid必须属于当前会话用户
*/
func TaobaoItemUpdateListingTmall(clt *core.SDKClient, req *product.TaobaoItemUpdateListingTmallRequest, session string) (*product.TaobaoItemUpdateListingTmallAPIResponse, error) {
    var resp product.TaobaoItemUpdateListingTmallAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
