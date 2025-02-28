package wms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
驱动保税交易订单发货 APIRequest
cainiao.bim.tradeorder.consign

驱动保税交易订单发货
*/
type CainiaoBimTradeorderConsignRequest struct {
    model.Params

    // 交易单号
    tradeId   string 

    // 仓储编码，ERP指定仓库发货时需要传值，编码由菜鸟提供
    storeCode   string 

    // 选择的线路ID非必填字段
    resId   string 

}

func NewCainiaoBimTradeorderConsignRequest() *CainiaoBimTradeorderConsignRequest{
    return &CainiaoBimTradeorderConsignRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoBimTradeorderConsignRequest) GetApiMethodName() string {
    return "cainiao.bim.tradeorder.consign"
}

func (r CainiaoBimTradeorderConsignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoBimTradeorderConsignRequest) SetTradeId(tradeId string) error {
    r.tradeId = tradeId
    r.Set("trade_id", tradeId)
    return nil
}

func (r CainiaoBimTradeorderConsignRequest) GetTradeId() string {
    return r.tradeId
}

func (r *CainiaoBimTradeorderConsignRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

func (r CainiaoBimTradeorderConsignRequest) GetStoreCode() string {
    return r.storeCode
}

func (r *CainiaoBimTradeorderConsignRequest) SetResId(resId string) error {
    r.resId = resId
    r.Set("res_id", resId)
    return nil
}

func (r CainiaoBimTradeorderConsignRequest) GetResId() string {
    return r.resId
}

