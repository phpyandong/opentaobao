package hotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
哈罗合作方获取酒店库存报价 APIRequest
taobao.xhotel.price.get.for.hello

哈罗合作项目，供哈罗合作方按需查询已开通城市下的标准酒店下指定时间段内的库存报价信息；在用户登录方面，返回结果不涉及用户个人信息，不涉及商家信息；仅根据不同用户，查询对应会员等级后，返回不同报价；
*/
type TaobaoXhotelPriceGetForHelloRequest struct {
    model.Params

    // 参数封装
    hotelPriceParam   *HotelPriceParam 

}

func NewTaobaoXhotelPriceGetForHelloRequest() *TaobaoXhotelPriceGetForHelloRequest{
    return &TaobaoXhotelPriceGetForHelloRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelPriceGetForHelloRequest) GetApiMethodName() string {
    return "taobao.xhotel.price.get.for.hello"
}

func (r TaobaoXhotelPriceGetForHelloRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelPriceGetForHelloRequest) SetHotelPriceParam(hotelPriceParam *HotelPriceParam) error {
    r.hotelPriceParam = hotelPriceParam
    r.Set("hotel_price_param", hotelPriceParam)
    return nil
}

func (r TaobaoXhotelPriceGetForHelloRequest) GetHotelPriceParam() *HotelPriceParam {
    return r.hotelPriceParam
}

