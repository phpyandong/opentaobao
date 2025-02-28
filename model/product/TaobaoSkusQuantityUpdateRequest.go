package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
SKU库存修改 APIRequest
taobao.skus.quantity.update

提供按照全量/增量的方式批量修改SKU库存的功能
*/
type TaobaoSkusQuantityUpdateRequest struct {
    model.Params

    // 商品数字ID，必填参数
    numIid   int64 

    // 库存更新方式，可选。1为全量更新，2为增量更新。如果不填，默认为全量更新。当选择全量更新时，如果库存更新值传入的是负数，会出错并返回错误码；当选择增量更新时，如果库存更新值为负数且绝对值大于当前库存，则sku库存会设置为0.
    type   int64 

    // sku库存批量修改入参，用于指定一批sku和每个sku的库存修改值，特殊可填。格式为skuId:库存修改值;skuId:库存修改值。最多支持20个SKU同时修改。
    skuidQuantities   string 

    // 特殊可选，skuIdQuantities为空的时候用该字段通过outerId来指定sku和其库存修改值。格式为outerId:库存修改值;outerId:库存修改值。当skuIdQuantities不为空的时候该字段失效。当一个outerId对应多个sku时，所有匹配到的sku都会被修改库存。最多支持20个SKU同时修改。
    outeridQuantities   string 

}

func NewTaobaoSkusQuantityUpdateRequest() *TaobaoSkusQuantityUpdateRequest{
    return &TaobaoSkusQuantityUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSkusQuantityUpdateRequest) GetApiMethodName() string {
    return "taobao.skus.quantity.update"
}

func (r TaobaoSkusQuantityUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSkusQuantityUpdateRequest) SetNumIid(numIid int64) error {
    r.numIid = numIid
    r.Set("num_iid", numIid)
    return nil
}

func (r TaobaoSkusQuantityUpdateRequest) GetNumIid() int64 {
    return r.numIid
}

func (r *TaobaoSkusQuantityUpdateRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TaobaoSkusQuantityUpdateRequest) GetType() int64 {
    return r.type
}

func (r *TaobaoSkusQuantityUpdateRequest) SetSkuidQuantities(skuidQuantities string) error {
    r.skuidQuantities = skuidQuantities
    r.Set("skuid_quantities", skuidQuantities)
    return nil
}

func (r TaobaoSkusQuantityUpdateRequest) GetSkuidQuantities() string {
    return r.skuidQuantities
}

func (r *TaobaoSkusQuantityUpdateRequest) SetOuteridQuantities(outeridQuantities string) error {
    r.outeridQuantities = outeridQuantities
    r.Set("outerid_quantities", outeridQuantities)
    return nil
}

func (r TaobaoSkusQuantityUpdateRequest) GetOuteridQuantities() string {
    return r.outeridQuantities
}

