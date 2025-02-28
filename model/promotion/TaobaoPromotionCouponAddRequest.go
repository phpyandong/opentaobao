package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建店铺优惠券接口 APIRequest
taobao.promotion.coupon.add

创建店铺优惠券。有效期内的店铺优惠券总数量不超过50张
*/
type TaobaoPromotionCouponAddRequest struct {
    model.Params

    // 优惠券的面额，必须是3，5，10，20，50，100
    denominations   int64 

    // 优惠券的截止日期
    endTime   string 

    // 订单满多少元才能用这个优惠券，500就是满500元才能使用
    condition   int64 

    // 优惠券的生效时间
    startTime   string 

}

func NewTaobaoPromotionCouponAddRequest() *TaobaoPromotionCouponAddRequest{
    return &TaobaoPromotionCouponAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPromotionCouponAddRequest) GetApiMethodName() string {
    return "taobao.promotion.coupon.add"
}

func (r TaobaoPromotionCouponAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPromotionCouponAddRequest) SetDenominations(denominations int64) error {
    r.denominations = denominations
    r.Set("denominations", denominations)
    return nil
}

func (r TaobaoPromotionCouponAddRequest) GetDenominations() int64 {
    return r.denominations
}

func (r *TaobaoPromotionCouponAddRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoPromotionCouponAddRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoPromotionCouponAddRequest) SetCondition(condition int64) error {
    r.condition = condition
    r.Set("condition", condition)
    return nil
}

func (r TaobaoPromotionCouponAddRequest) GetCondition() int64 {
    return r.condition
}

func (r *TaobaoPromotionCouponAddRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoPromotionCouponAddRequest) GetStartTime() string {
    return r.startTime
}

