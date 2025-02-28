package bill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
tae查询虚拟账户明细数据 APIResponse
taobao.tae.book.bills.get

tae查询虚拟账户明细数据
*/
type TaobaoTaeBookBillsGetAPIResponse struct {
    model.CommonResponse
    TaobaoTaeBookBillsGetResponse
}

type TaobaoTaeBookBillsGetResponse struct {
    XMLName xml.Name `xml:"tae_book_bills_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 虚拟账户账单列表
    
    Bills   []TopAcctCashJourDto `json:"bills,omitempty" xml:"bills>top_acct_cash_jour_dto,omitempty"`
    
    
    // 是否有下一页
    
    HasNext   bool `json:"has_next,omitempty" xml:"has_next,omitempty"`

    
    // 当前查询的结果数,非总数
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`

    
}
