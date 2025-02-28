package refund

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
特殊退款类型的纠纷单列表查询 APIResponse
taobao.special.refunds.receive.get

特殊退款类型的纠纷单列表查询
*/
type TaobaoSpecialRefundsReceiveGetAPIResponse struct {
    model.CommonResponse
    TaobaoSpecialRefundsReceiveGetResponse
}

type TaobaoSpecialRefundsReceiveGetResponse struct {
    XMLName xml.Name `xml:"special_refunds_receive_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 搜索到的退款信息总数
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`

    
    // 是否存在下一页
    
    HasNext   bool `json:"has_next,omitempty" xml:"has_next,omitempty"`

    
    // 搜索到的退款信息列表
    
    Refunds   []Refund `json:"refunds,omitempty" xml:"refunds>refund,omitempty"`
    
    
}
