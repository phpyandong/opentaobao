package mos

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mos"
)

/* 
退款查询 
alibaba.mos.onsite.trade.queryrefund

商户可使用该接口查询退款请求是否执行成功。
*/
func AlibabaMosOnsiteTradeQueryrefund(clt *core.SDKClient, req *mos.AlibabaMosOnsiteTradeQueryrefundRequest, session string) (*mos.AlibabaMosOnsiteTradeQueryrefundAPIResponse, error) {
    var resp mos.AlibabaMosOnsiteTradeQueryrefundAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
