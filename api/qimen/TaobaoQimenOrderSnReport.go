package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
订单SN通知接口 
taobao.qimen.order.sn.report

WMS调用奇门的接口,在出库、发货、入库等场景下，ERP和WMS之间同步操作的SN列表
*/
func TaobaoQimenOrderSnReport(clt *core.SDKClient, req *qimen.TaobaoQimenOrderSnReportRequest, session string) (*qimen.TaobaoQimenOrderSnReportAPIResponse, error) {
    var resp qimen.TaobaoQimenOrderSnReportAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
