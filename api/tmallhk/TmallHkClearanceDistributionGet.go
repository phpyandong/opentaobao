package tmallhk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallhk"
)

/* 
分销供应商获取清关材料 
tmall.hk.clearance.distribution.get

供销体系下，提供供应商可以直接获取其订单身份证信息的接口，以使其完成清关。
*/
func TmallHkClearanceDistributionGet(clt *core.SDKClient, req *tmallhk.TmallHkClearanceDistributionGetRequest, session string) (*tmallhk.TmallHkClearanceDistributionGetAPIResponse, error) {
    var resp tmallhk.TmallHkClearanceDistributionGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
