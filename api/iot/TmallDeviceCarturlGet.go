package iot

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/iot"
)

/* 
添加商品到购物车 
tmall.device.carturl.get

获取二维码，支持添加商品到购物车
*/
func TmallDeviceCarturlGet(clt *core.SDKClient, req *iot.TmallDeviceCarturlGetRequest, session string) (*iot.TmallDeviceCarturlGetAPIResponse, error) {
    var resp iot.TmallDeviceCarturlGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
