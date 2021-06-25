package wlb

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商家库存调整 APIResponse
cainiao.merchant.inventory.adjust

商家仓库存调整接口，目前仅支持全量更新
*/
type CainiaoMerchantInventoryAdjustAPIResponse struct {
    model.CommonResponse
    Response *CainiaoMerchantInventoryAdjustResponse `json:"cainiao_merchant_inventory_adjust_response,omitempty"`
}

type CainiaoMerchantInventoryAdjustResponse struct {

    // result
    Result   *SingleResultDto `json:"result,omitempty"`

}
