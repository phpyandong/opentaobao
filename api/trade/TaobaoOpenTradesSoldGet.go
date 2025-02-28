package trade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/trade"
)

/* 
查询卖家已卖出的交易数据（商家应用使用） 
taobao.open.trades.sold.get

搜索当前会话用户作为卖家已卖出的交易数据（只能获取到三个月以内的交易信息）<br/>
1. 返回的数据结果是以订单的创建时间倒序排列的。<br/>
注意：type字段的说明，如果该字段不传，接口默认只查4种类型订单，非默认查询的订单是不返回。遇到订单查不到的情况的，通常都是这个原因造成。解决办法就是type加上订单类型就可正常返回了。<br/>
2.入参fields中传入buyer_nick ，才能返回buyer_open_id
*/
func TaobaoOpenTradesSoldGet(clt *core.SDKClient, req *trade.TaobaoOpenTradesSoldGetRequest, session string) (*trade.TaobaoOpenTradesSoldGetAPIResponse, error) {
    var resp trade.TaobaoOpenTradesSoldGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
