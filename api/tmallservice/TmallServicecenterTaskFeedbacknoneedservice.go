package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
服务商反馈无需安装工单接口 
tmall.servicecenter.task.feedbacknoneedservice

服务商反馈无需安装工单接口
*/
func TmallServicecenterTaskFeedbacknoneedservice(clt *core.SDKClient, req *tmallservice.TmallServicecenterTaskFeedbacknoneedserviceRequest, session string) (*tmallservice.TmallServicecenterTaskFeedbacknoneedserviceAPIResponse, error) {
    var resp tmallservice.TmallServicecenterTaskFeedbacknoneedserviceAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
