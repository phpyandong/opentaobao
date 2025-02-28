package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
核销放行的查询接口 APIRequest
taobao.vmarket.eticket.auth.beforeconsume

针对O2O电子凭证核销放行业务，为满足码商能够核销淘宝码而开放的核销查询接口
*/
type TaobaoVmarketEticketAuthBeforeconsumeRequest struct {
    model.Params

    // 核销的码，只支持单个码，多个码核销需要多次调用
    verifyCode   string 

    // 核销方的ID，如果是普通码商必须传入机具ID,如果是私有码商家（即原有的信任商家）可默认传入私有码商ID
    operatorid   string 

    // 网点ID,网点授权核销时，必须传入；其他核销方式可不传
    storeid   string 

}

func NewTaobaoVmarketEticketAuthBeforeconsumeRequest() *TaobaoVmarketEticketAuthBeforeconsumeRequest{
    return &TaobaoVmarketEticketAuthBeforeconsumeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoVmarketEticketAuthBeforeconsumeRequest) GetApiMethodName() string {
    return "taobao.vmarket.eticket.auth.beforeconsume"
}

func (r TaobaoVmarketEticketAuthBeforeconsumeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoVmarketEticketAuthBeforeconsumeRequest) SetVerifyCode(verifyCode string) error {
    r.verifyCode = verifyCode
    r.Set("verify_code", verifyCode)
    return nil
}

func (r TaobaoVmarketEticketAuthBeforeconsumeRequest) GetVerifyCode() string {
    return r.verifyCode
}

func (r *TaobaoVmarketEticketAuthBeforeconsumeRequest) SetOperatorid(operatorid string) error {
    r.operatorid = operatorid
    r.Set("operatorid", operatorid)
    return nil
}

func (r TaobaoVmarketEticketAuthBeforeconsumeRequest) GetOperatorid() string {
    return r.operatorid
}

func (r *TaobaoVmarketEticketAuthBeforeconsumeRequest) SetStoreid(storeid string) error {
    r.storeid = storeid
    r.Set("storeid", storeid)
    return nil
}

func (r TaobaoVmarketEticketAuthBeforeconsumeRequest) GetStoreid() string {
    return r.storeid
}

