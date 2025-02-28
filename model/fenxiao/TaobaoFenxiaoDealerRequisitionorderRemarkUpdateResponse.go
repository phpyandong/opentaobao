package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改经销采购单备注 APIResponse
taobao.fenxiao.dealer.requisitionorder.remark.update

供应商修改经销采购单备注
*/
type TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoFenxiaoDealerRequisitionorderRemarkUpdateResponse
}

type TaobaoFenxiaoDealerRequisitionorderRemarkUpdateResponse struct {
    XMLName xml.Name `xml:"fenxiao_dealer_requisitionorder_remark_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 操作是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
