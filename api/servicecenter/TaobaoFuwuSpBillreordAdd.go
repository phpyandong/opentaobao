package servicecenter

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/servicecenter"
)

/* 
内购服务确认单明细上传接口 
taobao.fuwu.sp.billreord.add

isv能通过该接口上传确认单明细数据
*/
func TaobaoFuwuSpBillreordAdd(clt *core.SDKClient, req *servicecenter.TaobaoFuwuSpBillreordAddRequest, session string) (*servicecenter.TaobaoFuwuSpBillreordAddAPIResponse, error) {
    var resp servicecenter.TaobaoFuwuSpBillreordAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
