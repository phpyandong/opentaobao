package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
喵师傅核销接口 APIRequest
tmall.msf.verify

msf服务核销的top接口
*/
type TmallMsfVerifyRequest struct {
    model.Params

    // 111
    shopId   string 

    // 111
    bizType   string 

    // 111
    code   string 

}

func NewTmallMsfVerifyRequest() *TmallMsfVerifyRequest{
    return &TmallMsfVerifyRequest{
        Params: model.NewParams(),
    }
}

func (r TmallMsfVerifyRequest) GetApiMethodName() string {
    return "tmall.msf.verify"
}

func (r TmallMsfVerifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallMsfVerifyRequest) SetShopId(shopId string) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

func (r TmallMsfVerifyRequest) GetShopId() string {
    return r.shopId
}

func (r *TmallMsfVerifyRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

func (r TmallMsfVerifyRequest) GetBizType() string {
    return r.bizType
}

func (r *TmallMsfVerifyRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r TmallMsfVerifyRequest) GetCode() string {
    return r.code
}

