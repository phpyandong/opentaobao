package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
服务工单拉取 
tmall.servicecenter.task.get

接口供服务供应商通过交易主订单查询其未拉取的任务类工单
*/
func TmallServicecenterTaskGet(clt *core.SDKClient, req *tmallservice.TmallServicecenterTaskGetRequest, session string) (*tmallservice.TmallServicecenterTaskGetAPIResponse, error) {
    var resp tmallservice.TmallServicecenterTaskGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
