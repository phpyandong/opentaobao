package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
外部渠道查询退货订单详情接口 APIRequest
alibaba.wdk.trade.refund.query

该接口提供给外部渠道商家，比如欧尚外卖等查询退货订单详情，里面包含退货进度等信息。
*/
type AlibabaWdkTradeRefundQueryRequest struct {
    model.Params

    // 查询请求
    refundGoodsQueryRequest   *RefundGoodsQueryRequest 

}

func NewAlibabaWdkTradeRefundQueryRequest() *AlibabaWdkTradeRefundQueryRequest{
    return &AlibabaWdkTradeRefundQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkTradeRefundQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.trade.refund.query"
}

func (r AlibabaWdkTradeRefundQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkTradeRefundQueryRequest) SetRefundGoodsQueryRequest(refundGoodsQueryRequest *RefundGoodsQueryRequest) error {
    r.refundGoodsQueryRequest = refundGoodsQueryRequest
    r.Set("refund_goods_query_request", refundGoodsQueryRequest)
    return nil
}

func (r AlibabaWdkTradeRefundQueryRequest) GetRefundGoodsQueryRequest() *RefundGoodsQueryRequest {
    return r.refundGoodsQueryRequest
}

