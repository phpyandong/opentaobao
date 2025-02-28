package bill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询虚拟账户明细数据(自研发商家专用) APIResponse
taobao.bill.book.bills.get

查询虚拟账户明细数据
*/
type TaobaoBillBookBillsGetAPIResponse struct {
    model.CommonResponse
    TaobaoBillBookBillsGetResponse
}

type TaobaoBillBookBillsGetResponse struct {
    XMLName xml.Name `xml:"bill_book_bills_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否有下一页
    
    HasNext   bool `json:"has_next,omitempty" xml:"has_next,omitempty"`

    
    // 虚拟账户账单列表
    
    Bills   []BookBill `json:"bills,omitempty" xml:"bills>book_bill,omitempty"`
    
    
    // 当前查询的结果数,非总数
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`

    
}
