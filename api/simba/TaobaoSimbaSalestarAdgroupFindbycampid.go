package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
(销量明星)批量获取推广计划下的推广组信息 
taobao.simba.salestar.adgroup.findbycampid

批量得到推广计划下的推广组
*/
func TaobaoSimbaSalestarAdgroupFindbycampid(clt *core.SDKClient, req *simba.TaobaoSimbaSalestarAdgroupFindbycampidRequest, session string) (*simba.TaobaoSimbaSalestarAdgroupFindbycampidAPIResponse, error) {
    var resp simba.TaobaoSimbaSalestarAdgroupFindbycampidAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
