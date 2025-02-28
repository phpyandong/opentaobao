package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
快递改约api 
taobao.logistics.express.modify.appoint

商家通过此api操作修改物流单，交易单的收货人地址、收货人联系方式、预约配送日期
*/
func TaobaoLogisticsExpressModifyAppoint(clt *core.SDKClient, req *logistic.TaobaoLogisticsExpressModifyAppointRequest, session string) (*logistic.TaobaoLogisticsExpressModifyAppointAPIResponse, error) {
    var resp logistic.TaobaoLogisticsExpressModifyAppointAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
