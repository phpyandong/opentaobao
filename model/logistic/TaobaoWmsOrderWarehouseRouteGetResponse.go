package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取订单仓库路由信息 APIResponse
taobao.wms.order.warehouse.route.get

获取订单仓库路由信息
*/
type TaobaoWmsOrderWarehouseRouteGetAPIResponse struct {
    model.CommonResponse
    TaobaoWmsOrderWarehouseRouteGetResponse
}

type TaobaoWmsOrderWarehouseRouteGetResponse struct {
    XMLName xml.Name `xml:"wms_order_warehouse_route_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品列表
    
    Items   []OrderWarehouseRouteGetItems `json:"items,omitempty" xml:"items>order_warehouse_route_get_items,omitempty"`
    
    
    // 是否成功
    
    WlSuccess   bool `json:"wl_success,omitempty" xml:"wl_success,omitempty"`

    
    // 错误信息
    
    WlErrorCode   string `json:"wl_error_code,omitempty" xml:"wl_error_code,omitempty"`

    
    // 错误信息
    
    WlErrorMsg   string `json:"wl_error_msg,omitempty" xml:"wl_error_msg,omitempty"`

    
    // 订单编号
    
    OrderCode   string `json:"order_code,omitempty" xml:"order_code,omitempty"`

    
}
