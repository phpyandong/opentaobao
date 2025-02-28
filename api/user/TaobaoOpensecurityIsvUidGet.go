package user

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/user"
)

/* 
获取open security uid for isv 
taobao.opensecurity.isv.uid.get

根据 open_uid 获取 open_uid_isv 用于同一个 isv的多个app间数据关联
*/
func TaobaoOpensecurityIsvUidGet(clt *core.SDKClient, req *user.TaobaoOpensecurityIsvUidGetRequest, session string) (*user.TaobaoOpensecurityIsvUidGetAPIResponse, error) {
    var resp user.TaobaoOpensecurityIsvUidGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
