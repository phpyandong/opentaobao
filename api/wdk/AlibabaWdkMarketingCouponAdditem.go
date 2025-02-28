package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
在优惠券活动下添加商品。【注意，此接口暂不支持并发！】 
alibaba.wdk.marketing.coupon.additem

在优惠券活动下添加商品。【注意，此接口暂不支持并发！】
如果是商品券，则添加的商品为券适用的商品；
如果是品类券，则添加的商品为券排除的商品；
*/
func AlibabaWdkMarketingCouponAdditem(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingCouponAdditemRequest, session string) (*wdk.AlibabaWdkMarketingCouponAdditemAPIResponse, error) {
    var resp wdk.AlibabaWdkMarketingCouponAdditemAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
