package interact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/interact"
)

/* 
客户端打开新页面 
alibaba.interact.sensor.openwindow

客户端打开新页面
*/
func AlibabaInteractSensorOpenwindow(clt *core.SDKClient, req *interact.AlibabaInteractSensorOpenwindowRequest, session string) (*interact.AlibabaInteractSensorOpenwindowAPIResponse, error) {
    var resp interact.AlibabaInteractSensorOpenwindowAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
