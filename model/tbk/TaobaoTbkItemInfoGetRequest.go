package tbk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-公用-淘宝客商品详情查询(简版) APIRequest
taobao.tbk.item.info.get

淘宝客商品详情查询（简版）
*/
type TaobaoTbkItemInfoGetRequest struct {
    model.Params

    // 商品ID串，用,分割，最大40个
    numIids   string 

    // 链接形式：1：PC，2：无线，默认：１
    platform   int64 

    // ip地址，影响邮费获取，如果不传或者传入不准确，邮费无法精准提供
    ip   string 

}

func NewTaobaoTbkItemInfoGetRequest() *TaobaoTbkItemInfoGetRequest{
    return &TaobaoTbkItemInfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTbkItemInfoGetRequest) GetApiMethodName() string {
    return "taobao.tbk.item.info.get"
}

func (r TaobaoTbkItemInfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTbkItemInfoGetRequest) SetNumIids(numIids string) error {
    r.numIids = numIids
    r.Set("num_iids", numIids)
    return nil
}

func (r TaobaoTbkItemInfoGetRequest) GetNumIids() string {
    return r.numIids
}

func (r *TaobaoTbkItemInfoGetRequest) SetPlatform(platform int64) error {
    r.platform = platform
    r.Set("platform", platform)
    return nil
}

func (r TaobaoTbkItemInfoGetRequest) GetPlatform() int64 {
    return r.platform
}

func (r *TaobaoTbkItemInfoGetRequest) SetIp(ip string) error {
    r.ip = ip
    r.Set("ip", ip)
    return nil
}

func (r TaobaoTbkItemInfoGetRequest) GetIp() string {
    return r.ip
}

