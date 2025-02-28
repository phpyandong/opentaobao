package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批次库存查询接口 APIResponse
taobao.wlb.item.batch.query

根据用户id，item id list和store code来查询商品库存信息和批次信息
*/
type TaobaoWlbItemBatchQueryAPIResponse struct {
    model.CommonResponse
    TaobaoWlbItemBatchQueryResponse
}

type TaobaoWlbItemBatchQueryResponse struct {
    XMLName xml.Name `xml:"wlb_item_batch_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果记录的数量
    
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`

    
    // 商品库存及批次信息查询结果
    
    ItemInventoryBatchList   []WlbItemBatchInventory `json:"item_inventory_batch_list,omitempty" xml:"item_inventory_batch_list>wlb_item_batch_inventory,omitempty"`
    
    
}
