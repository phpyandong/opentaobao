package mos

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mos"
)

/* 
新商场当面付商户扫码付 
alibaba.mos.onsite.trade.pay

收银员使用扫码设备读取用户“付款码”后，将二维码或条码信息通过本接口上送至喵街发起支付。
*/
func AlibabaMosOnsiteTradePay(clt *core.SDKClient, req *mos.AlibabaMosOnsiteTradePayRequest, session string) (*mos.AlibabaMosOnsiteTradePayAPIResponse, error) {
    var resp mos.AlibabaMosOnsiteTradePayAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
