package qianniu

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qianniu"
)

/* 
创建轻任务 
taobao.qianniu.task.create

发起一个轻任务，分配给多个执行者，并发送消息提醒，由任务发起者调用
*/
func TaobaoQianniuTaskCreate(clt *core.SDKClient, req *qianniu.TaobaoQianniuTaskCreateRequest, session string) (*qianniu.TaobaoQianniuTaskCreateAPIResponse, error) {
    var resp qianniu.TaobaoQianniuTaskCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
