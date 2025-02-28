package trade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/trade"
)

/* 
查询外部交易订单接口 
alibaba.wdk.trade.order.query

通过该接口可以在盒马查询交易订单，并处理相关业务流程。主要用于和外部商户的订单进行同步和融合业务流程处理
*/
func AlibabaWdkTradeOrderQuery(clt *core.SDKClient, req *trade.AlibabaWdkTradeOrderQueryRequest, session string) (*trade.AlibabaWdkTradeOrderQueryAPIResponse, error) {
    var resp trade.AlibabaWdkTradeOrderQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
