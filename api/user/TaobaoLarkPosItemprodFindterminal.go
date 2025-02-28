package user

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/user"
)

/* 
终端配置支持 
taobao.lark.pos.itemprod.findterminal

终端配置支持,读取如果不存在则创建和远程的连接配置并返回
*/
func TaobaoLarkPosItemprodFindterminal(clt *core.SDKClient, req *user.TaobaoLarkPosItemprodFindterminalRequest, session string) (*user.TaobaoLarkPosItemprodFindterminalAPIResponse, error) {
    var resp user.TaobaoLarkPosItemprodFindterminalAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
