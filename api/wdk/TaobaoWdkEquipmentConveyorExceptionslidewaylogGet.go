package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
异常通道日志查询 
taobao.wdk.equipment.conveyor.exceptionslidewaylog.get

五道口悬挂链异常通道事件查询
*/
func TaobaoWdkEquipmentConveyorExceptionslidewaylogGet(clt *core.SDKClient, req *wdk.TaobaoWdkEquipmentConveyorExceptionslidewaylogGetRequest, session string) (*wdk.TaobaoWdkEquipmentConveyorExceptionslidewaylogGetAPIResponse, error) {
    var resp wdk.TaobaoWdkEquipmentConveyorExceptionslidewaylogGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
