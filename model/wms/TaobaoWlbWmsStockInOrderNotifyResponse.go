package wms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
入库通知单 APIResponse
taobao.wlb.wms.stock.in.order.notify

入库通知单
*/
type TaobaoWlbWmsStockInOrderNotifyAPIResponse struct {
    model.CommonResponse
    TaobaoWlbWmsStockInOrderNotifyResponse
}

type TaobaoWlbWmsStockInOrderNotifyResponse struct {
    XMLName xml.Name `xml:"wlb_wms_stock_in_order_notify_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    WlSuccess   bool `json:"wl_success,omitempty" xml:"wl_success,omitempty"`

    
    // 错误编码
    
    WlErrorCode   string `json:"wl_error_code,omitempty" xml:"wl_error_code,omitempty"`

    
    // 错误详细
    
    WlErrorMsg   string `json:"wl_error_msg,omitempty" xml:"wl_error_msg,omitempty"`

    
    // 仓储订单编码
    
    OrderCode   string `json:"order_code,omitempty" xml:"order_code,omitempty"`

    
}
