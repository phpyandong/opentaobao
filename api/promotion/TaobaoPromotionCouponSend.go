package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
店铺优惠券发放接口 
taobao.promotion.coupon.send

通过接口批量发放店铺优惠券（每次只能发送100张，只能发给当前授权卖家店铺的会员），发送成功则返回为空，发送失败则返回失败的买家列表和发送成功的买家和优惠券的number。注：如果所有买家都发放失败的话，is_success也为true，建议调用者根据返回的集合判断是否送入的买家都发放成功了
*/
func TaobaoPromotionCouponSend(clt *core.SDKClient, req *promotion.TaobaoPromotionCouponSendRequest, session string) (*promotion.TaobaoPromotionCouponSendAPIResponse, error) {
    var resp promotion.TaobaoPromotionCouponSendAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
