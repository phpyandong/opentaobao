package wangwang

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wangwang"
)

/* 
模糊聊天记录查询 
taobao.wangwang.abstract.logquery

模糊聊天记录查询
*/
func TaobaoWangwangAbstractLogquery(clt *core.SDKClient, req *wangwang.TaobaoWangwangAbstractLogqueryRequest, session string) (*wangwang.TaobaoWangwangAbstractLogqueryAPIResponse, error) {
    var resp wangwang.TaobaoWangwangAbstractLogqueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
