package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
优惠券商品查询 
alibaba.wdk.coupon.sku.query

优惠券商品查询
*/
func AlibabaWdkCouponSkuQuery(clt *core.SDKClient, req *promotion.AlibabaWdkCouponSkuQueryRequest, session string) (*promotion.AlibabaWdkCouponSkuQueryAPIResponse, error) {
    var resp promotion.AlibabaWdkCouponSkuQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
