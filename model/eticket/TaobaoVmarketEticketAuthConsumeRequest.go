package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
核销放行的核销接口 APIRequest
taobao.vmarket.eticket.auth.consume

针对O2O电子凭证核销放行业务，为满足码商能够核销淘宝码而开放的核销接口
*/
type TaobaoVmarketEticketAuthConsumeRequest struct {
    model.Params

    // 核销的码，只支持单个码，多个码核销需要多次调用
    verifyCode   string 

    // 核销份数
    consumeNum   int64 

    // 核销方的ID，如果是普通码商必须传入机具ID,如果是私有码商家（即原有的信任商家）可默认传入私有码商ID
    operatorid   string 

    // 自定义核销流水号，需要小于等于100个字符(a-zA-Z0-9_)
    serialNum   string 

    // 网点ID,网点授权核销时，必须传入；其他核销方式可不传
    storeid   string 

}

func NewTaobaoVmarketEticketAuthConsumeRequest() *TaobaoVmarketEticketAuthConsumeRequest{
    return &TaobaoVmarketEticketAuthConsumeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoVmarketEticketAuthConsumeRequest) GetApiMethodName() string {
    return "taobao.vmarket.eticket.auth.consume"
}

func (r TaobaoVmarketEticketAuthConsumeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoVmarketEticketAuthConsumeRequest) SetVerifyCode(verifyCode string) error {
    r.verifyCode = verifyCode
    r.Set("verify_code", verifyCode)
    return nil
}

func (r TaobaoVmarketEticketAuthConsumeRequest) GetVerifyCode() string {
    return r.verifyCode
}

func (r *TaobaoVmarketEticketAuthConsumeRequest) SetConsumeNum(consumeNum int64) error {
    r.consumeNum = consumeNum
    r.Set("consume_num", consumeNum)
    return nil
}

func (r TaobaoVmarketEticketAuthConsumeRequest) GetConsumeNum() int64 {
    return r.consumeNum
}

func (r *TaobaoVmarketEticketAuthConsumeRequest) SetOperatorid(operatorid string) error {
    r.operatorid = operatorid
    r.Set("operatorid", operatorid)
    return nil
}

func (r TaobaoVmarketEticketAuthConsumeRequest) GetOperatorid() string {
    return r.operatorid
}

func (r *TaobaoVmarketEticketAuthConsumeRequest) SetSerialNum(serialNum string) error {
    r.serialNum = serialNum
    r.Set("serial_num", serialNum)
    return nil
}

func (r TaobaoVmarketEticketAuthConsumeRequest) GetSerialNum() string {
    return r.serialNum
}

func (r *TaobaoVmarketEticketAuthConsumeRequest) SetStoreid(storeid string) error {
    r.storeid = storeid
    r.Set("storeid", storeid)
    return nil
}

func (r TaobaoVmarketEticketAuthConsumeRequest) GetStoreid() string {
    return r.storeid
}

