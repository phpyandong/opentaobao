package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
供应商同步接口 
taobao.qimen.supplier.synchronize

这个接口用来同步供应商信息
*/
func TaobaoQimenSupplierSynchronize(clt *core.SDKClient, req *qimen.TaobaoQimenSupplierSynchronizeRequest, session string) (*qimen.TaobaoQimenSupplierSynchronizeAPIResponse, error) {
    var resp qimen.TaobaoQimenSupplierSynchronizeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
