package wms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
通过订单号获取单个出库单发货信息 APIResponse
taobao.wlb.wms.stock.out.bill.get

通过订单号获取单个出库单发货信息
*/
type TaobaoWlbWmsStockOutBillGetAPIResponse struct {
    model.CommonResponse
    TaobaoWlbWmsStockOutBillGetResponse
}

type TaobaoWlbWmsStockOutBillGetResponse struct {
    XMLName xml.Name `xml:"wlb_wms_stock_out_bill_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 出库信息
    
    StockOutInfo   *CainiaoStockOutBillStockoutinfo `json:"stock_out_info,omitempty" xml:"stock_out_info,omitempty"`

    
}
