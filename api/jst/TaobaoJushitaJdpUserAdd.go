package jst

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jst"
)

/* 
添加数据推送用户 
taobao.jushita.jdp.user.add

提供给接入数据推送的应用添加数据推送服务的用户
*/
func TaobaoJushitaJdpUserAdd(clt *core.SDKClient, req *jst.TaobaoJushitaJdpUserAddRequest, session string) (*jst.TaobaoJushitaJdpUserAddAPIResponse, error) {
    var resp jst.TaobaoJushitaJdpUserAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
