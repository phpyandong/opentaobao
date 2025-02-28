package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单和商品同步接口 APIRequest
alibaba.wdk.order.sync.withitem

轻pos,将订单和商品的信息一起传到盒马这边，进行创单和添加商品处理。
*/
type AlibabaWdkOrderSyncWithitemRequest struct {
    model.Params

    // 商家传过来的交易和商品信息
    posOrderAndItemSync   *PosOrderAndItemSyncDo 

}

func NewAlibabaWdkOrderSyncWithitemRequest() *AlibabaWdkOrderSyncWithitemRequest{
    return &AlibabaWdkOrderSyncWithitemRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkOrderSyncWithitemRequest) GetApiMethodName() string {
    return "alibaba.wdk.order.sync.withitem"
}

func (r AlibabaWdkOrderSyncWithitemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkOrderSyncWithitemRequest) SetPosOrderAndItemSync(posOrderAndItemSync *PosOrderAndItemSyncDo) error {
    r.posOrderAndItemSync = posOrderAndItemSync
    r.Set("pos_order_and_item_sync", posOrderAndItemSync)
    return nil
}

func (r AlibabaWdkOrderSyncWithitemRequest) GetPosOrderAndItemSync() *PosOrderAndItemSyncDo {
    return r.posOrderAndItemSync
}

