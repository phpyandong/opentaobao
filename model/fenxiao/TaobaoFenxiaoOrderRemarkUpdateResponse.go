package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改采购单备注 APIResponse
taobao.fenxiao.order.remark.update

供应商修改采购单备注
*/
type TaobaoFenxiaoOrderRemarkUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoFenxiaoOrderRemarkUpdateResponse
}

type TaobaoFenxiaoOrderRemarkUpdateResponse struct {
    XMLName xml.Name `xml:"fenxiao_order_remark_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 操作是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
