package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
查询卖家优惠券 
taobao.promotion.coupons.get

查询卖家已经创建的优惠券，接口返回信息：优惠券ID，面值，创建时间，有效期，使用条件，使用渠道，创建渠道，优惠券总数量
*/
func TaobaoPromotionCouponsGet(clt *core.SDKClient, req *promotion.TaobaoPromotionCouponsGetRequest, session string) (*promotion.TaobaoPromotionCouponsGetAPIResponse, error) {
    var resp promotion.TaobaoPromotionCouponsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
