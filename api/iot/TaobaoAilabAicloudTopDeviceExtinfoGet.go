package iot

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/iot"
)

/* 
获取设备扩展信息 
taobao.ailab.aicloud.top.device.extinfo.get

获取设备扩展信息
*/
func TaobaoAilabAicloudTopDeviceExtinfoGet(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceExtinfoGetRequest, session string) (*iot.TaobaoAilabAicloudTopDeviceExtinfoGetAPIResponse, error) {
    var resp iot.TaobaoAilabAicloudTopDeviceExtinfoGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
