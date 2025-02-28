package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据交易号获取物流宝订单 APIResponse
taobao.wlb.tradeorder.get

根据交易类型和交易id查询物流宝订单详情
*/
type TaobaoWlbTradeorderGetAPIResponse struct {
    model.CommonResponse
    TaobaoWlbTradeorderGetResponse
}

type TaobaoWlbTradeorderGetResponse struct {
    XMLName xml.Name `xml:"wlb_tradeorder_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 物流宝订单对象
    
    WlbOrderList   []WlbOrder `json:"wlb_order_list,omitempty" xml:"wlb_order_list>wlb_order,omitempty"`
    
    
}
