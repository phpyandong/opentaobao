package util

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/util"
)

/* 
网关一次性token获取 
taobao.top.once.token.get

网关一次性token获取
*/
func TaobaoTopOnceTokenGet(clt *core.SDKClient, req *util.TaobaoTopOnceTokenGetRequest, session string) (*util.TaobaoTopOnceTokenGetAPIResponse, error) {
    var resp util.TaobaoTopOnceTokenGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
