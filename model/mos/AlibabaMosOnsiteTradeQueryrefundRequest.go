package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退款查询 APIRequest
alibaba.mos.onsite.trade.queryrefund

商户可使用该接口查询退款请求是否执行成功。
*/
type AlibabaMosOnsiteTradeQueryrefundRequest struct {
    model.Params

    // 退款外部流水号
    outRequestNo   string 

    // 订单号。可能为外部订单号，也可能为喵街订单号
    orderNo   string 

}

func NewAlibabaMosOnsiteTradeQueryrefundRequest() *AlibabaMosOnsiteTradeQueryrefundRequest{
    return &AlibabaMosOnsiteTradeQueryrefundRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMosOnsiteTradeQueryrefundRequest) GetApiMethodName() string {
    return "alibaba.mos.onsite.trade.queryrefund"
}

func (r AlibabaMosOnsiteTradeQueryrefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMosOnsiteTradeQueryrefundRequest) SetOutRequestNo(outRequestNo string) error {
    r.outRequestNo = outRequestNo
    r.Set("out_request_no", outRequestNo)
    return nil
}

func (r AlibabaMosOnsiteTradeQueryrefundRequest) GetOutRequestNo() string {
    return r.outRequestNo
}

func (r *AlibabaMosOnsiteTradeQueryrefundRequest) SetOrderNo(orderNo string) error {
    r.orderNo = orderNo
    r.Set("order_no", orderNo)
    return nil
}

func (r AlibabaMosOnsiteTradeQueryrefundRequest) GetOrderNo() string {
    return r.orderNo
}

