package trade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/trade"
)

/* 
更改交易的收货地址 
taobao.trade.shippingaddress.update

只能更新一笔交易里面的买家收货地址 <br/>只能更新发货前（即买家已付款，等待卖家发货状态）的交易的买家收货地址 <br/>更新后的发货地址可以通过taobao.trade.fullinfo.get查到 <br/>参数中所说的字节为GBK编码的（英文和数字占1字节，中文占2字节）
*/
func TaobaoTradeShippingaddressUpdate(clt *core.SDKClient, req *trade.TaobaoTradeShippingaddressUpdateRequest, session string) (*trade.TaobaoTradeShippingaddressUpdateAPIResponse, error) {
    var resp trade.TaobaoTradeShippingaddressUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
