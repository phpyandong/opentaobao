package xhotel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhotel"
)

/* 
获取直连酒店（客栈）签约上线进度信息 
taobao.xhotel.order.hotelsign.query

获取直连酒店（客栈）签约上线进度信息
*/
func TaobaoXhotelOrderHotelsignQuery(clt *core.SDKClient, req *xhotel.TaobaoXhotelOrderHotelsignQueryRequest, session string) (*xhotel.TaobaoXhotelOrderHotelsignQueryAPIResponse, error) {
    var resp xhotel.TaobaoXhotelOrderHotelsignQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
