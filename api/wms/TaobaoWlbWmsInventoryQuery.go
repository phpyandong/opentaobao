package wms

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wms"
)

/* 
菜鸟商品库存查询 
taobao.wlb.wms.inventory.query

支持按汇总（不分批次和渠道的总的库存数量）、渠道、批次三类方式查询商品实时库存
*/
func TaobaoWlbWmsInventoryQuery(clt *core.SDKClient, req *wms.TaobaoWlbWmsInventoryQueryRequest, session string) (*wms.TaobaoWlbWmsInventoryQueryAPIResponse, error) {
    var resp wms.TaobaoWlbWmsInventoryQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
