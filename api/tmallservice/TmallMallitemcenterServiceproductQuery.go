package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
天猫服务商服务产品查询 
tmall.mallitemcenter.serviceproduct.query

查询天猫服务的服务商发布的服务产品
*/
func TmallMallitemcenterServiceproductQuery(clt *core.SDKClient, req *tmallservice.TmallMallitemcenterServiceproductQueryRequest, session string) (*tmallservice.TmallMallitemcenterServiceproductQueryAPIResponse, error) {
    var resp tmallservice.TmallMallitemcenterServiceproductQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
