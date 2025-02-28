package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】度假线路商品发布时基础信息获取接口：邮轮扩展信息获取 APIRequest
taobao.alitrip.travel.baseinfo.cruise.get

旅行度假新商品发布时可用的扩展接口，用于获取邮轮类目相关扩展信息。
*/
type TaobaoAlitripTravelBaseinfoCruiseGetRequest struct {
    model.Params

    // true-获取国际邮轮类目扩展信息；false-获取国内邮轮类目扩展信息
    isOverseas   bool 

}

func NewTaobaoAlitripTravelBaseinfoCruiseGetRequest() *TaobaoAlitripTravelBaseinfoCruiseGetRequest{
    return &TaobaoAlitripTravelBaseinfoCruiseGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripTravelBaseinfoCruiseGetRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.baseinfo.cruise.get"
}

func (r TaobaoAlitripTravelBaseinfoCruiseGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripTravelBaseinfoCruiseGetRequest) SetIsOverseas(isOverseas bool) error {
    r.isOverseas = isOverseas
    r.Set("is_overseas", isOverseas)
    return nil
}

func (r TaobaoAlitripTravelBaseinfoCruiseGetRequest) GetIsOverseas() bool {
    return r.isOverseas
}

