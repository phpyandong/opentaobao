package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建通用单品优惠活动 APIRequest
taobao.promotionmisc.common.item.activity.add

创建通用单品优惠活动。
1、该接口只创建活动的基本信息，如需要增加、删除参与该活动的商品及优惠，请调用taobao.promotionmisc.common.item.detail.add和taobao.promotionmisc.common.item.detail.delete接口
2、同一卖家下的活动数量限制为30个，超过限制需先调用taobao.promotionmisc.common.item.activity.delete接口删除无用的活动后才可再创建新的活动
*/
type TaobaoPromotionmiscCommonItemActivityAddRequest struct {
    model.Params

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

func NewTaobaoPromotionmiscCommonItemActivityAddRequest() *TaobaoPromotionmiscCommonItemActivityAddRequest{
    return &TaobaoPromotionmiscCommonItemActivityAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPromotionmiscCommonItemActivityAddRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.common.item.activity.add"
}

func (r TaobaoPromotionmiscCommonItemActivityAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPromotionmiscCommonItemActivityAddRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoPromotionmiscCommonItemActivityAddRequest) GetName() string {
    return r.name
}

func (r *TaobaoPromotionmiscCommonItemActivityAddRequest) SetDescription(description string) error {
    r.description = description
    r.Set("description", description)
    return nil
}

func (r TaobaoPromotionmiscCommonItemActivityAddRequest) GetDescription() string {
    return r.description
}

func (r *TaobaoPromotionmiscCommonItemActivityAddRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoPromotionmiscCommonItemActivityAddRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoPromotionmiscCommonItemActivityAddRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoPromotionmiscCommonItemActivityAddRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoPromotionmiscCommonItemActivityAddRequest) SetIsUserTag(isUserTag bool) error {
    r.isUserTag = isUserTag
    r.Set("is_user_tag", isUserTag)
    return nil
}

func (r TaobaoPromotionmiscCommonItemActivityAddRequest) GetIsUserTag() bool {
    return r.isUserTag
}

func (r *TaobaoPromotionmiscCommonItemActivityAddRequest) SetUserTag(userTag string) error {
    r.userTag = userTag
    r.Set("user_tag", userTag)
    return nil
}

func (r TaobaoPromotionmiscCommonItemActivityAddRequest) GetUserTag() string {
    return r.userTag
}

