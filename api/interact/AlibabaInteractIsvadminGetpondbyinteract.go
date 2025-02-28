package interact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/interact"
)

/* 
根据互动实例ID查询奖池信息 
alibaba.interact.isvadmin.getpondbyinteract

根据互动实例ID查询奖池信息
*/
func AlibabaInteractIsvadminGetpondbyinteract(clt *core.SDKClient, req *interact.AlibabaInteractIsvadminGetpondbyinteractRequest, session string) (*interact.AlibabaInteractIsvadminGetpondbyinteractAPIResponse, error) {
    var resp interact.AlibabaInteractIsvadminGetpondbyinteractAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
