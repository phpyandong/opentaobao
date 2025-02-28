package tanx

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tanx"
)

/* 
创意修改接口 
taobao.tanx.audit.creative.modify

创意修改接口
*/
func TaobaoTanxAuditCreativeModify(clt *core.SDKClient, req *tanx.TaobaoTanxAuditCreativeModifyRequest, session string) (*tanx.TaobaoTanxAuditCreativeModifyAPIResponse, error) {
    var resp tanx.TaobaoTanxAuditCreativeModifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
