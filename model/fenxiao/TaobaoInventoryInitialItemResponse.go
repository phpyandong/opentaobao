package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品库存初始化 APIResponse
taobao.inventory.initial.item

建议使用新接口：taobao.inventory.merchant.adjust ，该接口会逐步停用。
商家仓商品初始化在各个仓中库存
*/
type TaobaoInventoryInitialItemAPIResponse struct {
    model.CommonResponse
    TaobaoInventoryInitialItemResponse
}

type TaobaoInventoryInitialItemResponse struct {
    XMLName xml.Name `xml:"inventory_initial_item_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 提示信息
    
    TipInfos   []TipInfo `json:"tip_infos,omitempty" xml:"tip_infos>tip_info,omitempty"`
    
    
}
