package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AG退货入仓状态写接口 APIRequest
taobao.nextone.logistics.warehouse.update

商家上传退货入仓状态给ag
*/
type TaobaoNextoneLogisticsWarehouseUpdateRequest struct {
    model.Params

    // 退款编号
    refundId   int64 

    // 退货入仓状态 1.已入仓
    warehouseStatus   int64 

}

func NewTaobaoNextoneLogisticsWarehouseUpdateRequest() *TaobaoNextoneLogisticsWarehouseUpdateRequest{
    return &TaobaoNextoneLogisticsWarehouseUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoNextoneLogisticsWarehouseUpdateRequest) GetApiMethodName() string {
    return "taobao.nextone.logistics.warehouse.update"
}

func (r TaobaoNextoneLogisticsWarehouseUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoNextoneLogisticsWarehouseUpdateRequest) SetRefundId(refundId int64) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

func (r TaobaoNextoneLogisticsWarehouseUpdateRequest) GetRefundId() int64 {
    return r.refundId
}

func (r *TaobaoNextoneLogisticsWarehouseUpdateRequest) SetWarehouseStatus(warehouseStatus int64) error {
    r.warehouseStatus = warehouseStatus
    r.Set("warehouse_status", warehouseStatus)
    return nil
}

func (r TaobaoNextoneLogisticsWarehouseUpdateRequest) GetWarehouseStatus() int64 {
    return r.warehouseStatus
}

