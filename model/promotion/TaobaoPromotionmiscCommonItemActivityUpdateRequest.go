package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改通用单品优惠活动 APIRequest
taobao.promotionmisc.common.item.activity.update

修改通用单品优惠活动。
1、该接口只修改活动基本信息，如需要增加、删除参与该活动的商品及优惠，请调用taobao.promotionmisc.common.item.detail.add和taobao.promotionmisc.common.item.detail.delete接口
2、使用该接口时需要把未做修改的字段值也传入
*/
type TaobaoPromotionmiscCommonItemActivityUpdateRequest struct {
    model.Params

    // 优惠活动ID
    activityId   int64 

    // 活动名称，不能超过32字符
    name   string 

    // 活动描述，不能超过100字符
    description   string 

    // 活动开始时间
    startTime   string 

    // 活动结束时间
    endTime   string 

    // 是否指定人群标签
    isUserTag   bool 

    // 用户标签。当is_user_tag为true时，该值才有意义。
    userTag   string 

}

func NewTaobaoPromotionmiscCommonItemActivityUpdateRequest() *TaobaoPromotionmiscCommonItemActivityUpdateRequest{
    return &TaobaoPromotionmiscCommonItemActivityUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPromotionmiscCommonItemActivityUpdateRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.common.item.activity.update"
}

func (r TaobaoPromotionmiscCommonItemActivityUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPromotionmiscCommonItemActivityUpdateRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

func (r TaobaoPromotionmiscCommonItemActivityUpdateRequest) GetActivityId() int64 {
    return r.activityId
}

func (r *TaobaoPromotionmiscCommonItemActivityUpdateRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoPromotionmiscCommonItemActivityUpdateRequest) GetName() string {
    return r.name
}

func (r *TaobaoPromotionmiscCommonItemActivityUpdateRequest) SetDescription(description string) error {
    r.description = description
    r.Set("description", description)
    return nil
}

func (r TaobaoPromotionmiscCommonItemActivityUpdateRequest) GetDescription() string {
    return r.description
}

func (r *TaobaoPromotionmiscCommonItemActivityUpdateRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoPromotionmiscCommonItemActivityUpdateRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoPromotionmiscCommonItemActivityUpdateRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoPromotionmiscCommonItemActivityUpdateRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoPromotionmiscCommonItemActivityUpdateRequest) SetIsUserTag(isUserTag bool) error {
    r.isUserTag = isUserTag
    r.Set("is_user_tag", isUserTag)
    return nil
}

func (r TaobaoPromotionmiscCommonItemActivityUpdateRequest) GetIsUserTag() bool {
    return r.isUserTag
}

func (r *TaobaoPromotionmiscCommonItemActivityUpdateRequest) SetUserTag(userTag string) error {
    r.userTag = userTag
    r.Set("user_tag", userTag)
    return nil
}

func (r TaobaoPromotionmiscCommonItemActivityUpdateRequest) GetUserTag() string {
    return r.userTag
}

