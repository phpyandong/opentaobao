package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分销渠道订单详情查询 APIResponse
taobao.xhotel.distribution.order.detail.search

该接口用于提供酒店分销渠道的订单详情查询
*/
type TaobaoXhotelDistributionOrderDetailSearchAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelDistributionOrderDetailSearchResponse
}

type TaobaoXhotelDistributionOrderDetailSearchResponse struct {
    XMLName xml.Name `xml:"xhotel_distribution_order_detail_search_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误码
    
    Error   string `json:"error,omitempty" xml:"error,omitempty"`

    
    // 错误描述
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`

    
    // 订单详情对象
    
    TopDistributionOrderDetail   *TopDistributionOrderDetail `json:"top_distribution_order_detail,omitempty" xml:"top_distribution_order_detail,omitempty"`

    
}
