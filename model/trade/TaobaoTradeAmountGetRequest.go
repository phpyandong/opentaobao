package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
交易帐务查询 APIRequest
taobao.trade.amount.get

卖家查询该笔交易的资金帐务相关的数据；
1. 只供卖家使用，买家不可使用
2. 可查询所有的状态的交易，但不同状态时交易的相关数据可能会有不同
*/
type TaobaoTradeAmountGetRequest struct {
    model.Params

    // 交易编号
    tid   int64 

    // 订单帐务详情需要返回的字段信息，可选值如下：1. TradeAmount中可指定的fields：tid,alipay_no,created,pay_time,end_time,total_fee,payment,post_fee,cod_fee,commission_fee,buyer_obtain_point_fee2. OrderAmount中可指定的fields：order_amounts.oid,order_amounts.title,order_amounts.num_iid,order_amounts.sku_properties_name,order_amounts.sku_id,order_amounts.num,order_amounts.price,order_amounts.discount_fee,order_amounts.adjust_fee,order_amounts.payment,order_amounts.promotion_name3. order_amounts(返回OrderAmount的所有内容)4. promotion_details(指定该值会返回主订单的promotion_details中除id之外的所有字段)
    fields   string 

}

func NewTaobaoTradeAmountGetRequest() *TaobaoTradeAmountGetRequest{
    return &TaobaoTradeAmountGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTradeAmountGetRequest) GetApiMethodName() string {
    return "taobao.trade.amount.get"
}

func (r TaobaoTradeAmountGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTradeAmountGetRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoTradeAmountGetRequest) GetTid() int64 {
    return r.tid
}

func (r *TaobaoTradeAmountGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoTradeAmountGetRequest) GetFields() string {
    return r.fields
}

