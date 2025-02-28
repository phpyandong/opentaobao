package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ASCP渠道中心销售单创建接口 APIResponse
tmall.ascp.orders.sale.create

ASCP渠道中心销售单创建接口
*/
type TmallAscpOrdersSaleCreateAPIResponse struct {
    model.CommonResponse
    TmallAscpOrdersSaleCreateResponse
}

type TmallAscpOrdersSaleCreateResponse struct {
    XMLName xml.Name `xml:"tmall_ascp_orders_sale_create_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TmallAscpOrdersSaleCreateResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
