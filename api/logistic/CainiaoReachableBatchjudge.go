package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
是否派送可达判定批量查询接口 
cainiao.reachable.batchjudge

提供给商家在发货之前做截单处理，输入物流商编码和收发货地址进行可达判定，目前支持国内主流的物流服务商, 支持快运和快递两种类型
*/
func CainiaoReachableBatchjudge(clt *core.SDKClient, req *logistic.CainiaoReachableBatchjudgeRequest, session string) (*logistic.CainiaoReachableBatchjudgeAPIResponse, error) {
    var resp logistic.CainiaoReachableBatchjudgeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
