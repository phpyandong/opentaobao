package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询某个卖家的店铺优惠券领取活动 APIRequest
taobao.promotion.activity.get

查询某个卖家的店铺优惠券领取活动<br/>返回，优惠券领取活动ID，优惠券ID，总领用量，每人限领量，已领取数量<br/>领取活动状态，优惠券领取链接<br/>最多50个优惠券
*/
type TaobaoPromotionActivityGetRequest struct {
    model.Params

    // 活动的id
    activityId   int64 

}

func NewTaobaoPromotionActivityGetRequest() *TaobaoPromotionActivityGetRequest{
    return &TaobaoPromotionActivityGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPromotionActivityGetRequest) GetApiMethodName() string {
    return "taobao.promotion.activity.get"
}

func (r TaobaoPromotionActivityGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPromotionActivityGetRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

func (r TaobaoPromotionActivityGetRequest) GetActivityId() int64 {
    return r.activityId
}

