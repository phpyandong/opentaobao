package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
主动发起通知接口 APIRequest
taobao.vmarket.eticket.manage.notify

外部合作商家主动发起通知接口
*/
type TaobaoVmarketEticketManageNotifyRequest struct {
    model.Params

    // 订单编号
    orderId   int64 

    // 码商ID，如果是码商，必须传，如果是信任卖家，不需要传
    codemerchantId   int64 

    // 需要调用的通知方法，目前仅支持是send（发码）或resend（重新发码）
    notifyMethod   string 

}

func NewTaobaoVmarketEticketManageNotifyRequest() *TaobaoVmarketEticketManageNotifyRequest{
    return &TaobaoVmarketEticketManageNotifyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoVmarketEticketManageNotifyRequest) GetApiMethodName() string {
    return "taobao.vmarket.eticket.manage.notify"
}

func (r TaobaoVmarketEticketManageNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoVmarketEticketManageNotifyRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r TaobaoVmarketEticketManageNotifyRequest) GetOrderId() int64 {
    return r.orderId
}

func (r *TaobaoVmarketEticketManageNotifyRequest) SetCodemerchantId(codemerchantId int64) error {
    r.codemerchantId = codemerchantId
    r.Set("codemerchant_id", codemerchantId)
    return nil
}

func (r TaobaoVmarketEticketManageNotifyRequest) GetCodemerchantId() int64 {
    return r.codemerchantId
}

func (r *TaobaoVmarketEticketManageNotifyRequest) SetNotifyMethod(notifyMethod string) error {
    r.notifyMethod = notifyMethod
    r.Set("notify_method", notifyMethod)
    return nil
}

func (r TaobaoVmarketEticketManageNotifyRequest) GetNotifyMethod() string {
    return r.notifyMethod
}

