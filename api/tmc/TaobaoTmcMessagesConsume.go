package tmc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmc"
)

/* 
消费多条消息 
taobao.tmc.messages.consume

消费多条消息。消费时如果没有返回消息，建议做控制，不要一直调api，浪费应用的流量。如对程序做好优化，若没有消息则，sleep 100ms 等。
*/
func TaobaoTmcMessagesConsume(clt *core.SDKClient, req *tmc.TaobaoTmcMessagesConsumeRequest, session string) (*tmc.TaobaoTmcMessagesConsumeAPIResponse, error) {
    var resp tmc.TaobaoTmcMessagesConsumeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
