package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加活动范围 APIRequest
taobao.ump.range.add

指定某项活动中，某个商家的某些类型物品（指定商品或者别的）参加或者不参加活动。当活动详情的参与类型为部分或者部分不参加的时候，需要指定具体哪部分参加或者不参加，使用本接口完成操作。比如部分商品满就送，这里的range用来指定具体哪些商品参加满就送活动。
*/
type TaobaoUmpRangeAddRequest struct {
    model.Params

    // 活动id
    actId   int64 

    // 范围的类型，比如是全店，商品，见：MarketingConstants.PARTICIPATE_TYPE_*
    type   int64 

    // id列表，当范围类型为商品时，该id为商品id.多个id用逗号隔开，一次不超过50个
    ids   string 

}

func NewTaobaoUmpRangeAddRequest() *TaobaoUmpRangeAddRequest{
    return &TaobaoUmpRangeAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUmpRangeAddRequest) GetApiMethodName() string {
    return "taobao.ump.range.add"
}

func (r TaobaoUmpRangeAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUmpRangeAddRequest) SetActId(actId int64) error {
    r.actId = actId
    r.Set("act_id", actId)
    return nil
}

func (r TaobaoUmpRangeAddRequest) GetActId() int64 {
    return r.actId
}

func (r *TaobaoUmpRangeAddRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TaobaoUmpRangeAddRequest) GetType() int64 {
    return r.type
}

func (r *TaobaoUmpRangeAddRequest) SetIds(ids string) error {
    r.ids = ids
    r.Set("ids", ids)
    return nil
}

func (r TaobaoUmpRangeAddRequest) GetIds() string {
    return r.ids
}

