package tmallhk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallhk"
)

/* 
上传小票数据 
tmall.traceplatform.ticket.order.upload

upsertOrderBySeller
*/
func TmallTraceplatformTicketOrderUpload(clt *core.SDKClient, req *tmallhk.TmallTraceplatformTicketOrderUploadRequest, session string) (*tmallhk.TmallTraceplatformTicketOrderUploadAPIResponse, error) {
    var resp tmallhk.TmallTraceplatformTicketOrderUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
