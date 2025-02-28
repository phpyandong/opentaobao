package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
按照日期范围查询物流订单详情 APIResponse
taobao.wlb.orderdetail.date.get

外部ERP可通过该接口查询一段时间内的物流宝订单，以及订单详情
*/
type TaobaoWlbOrderdetailDateGetAPIResponse struct {
    model.CommonResponse
    TaobaoWlbOrderdetailDateGetResponse
}

type TaobaoWlbOrderdetailDateGetResponse struct {
    XMLName xml.Name `xml:"wlb_orderdetail_date_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 物流宝订单，并且包含订单详情
    
    OrderDetailList   []WlbOrderDetail `json:"order_detail_list,omitempty" xml:"order_detail_list>wlb_order_detail,omitempty"`
    
    
    // 总数
    
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`

    
}
