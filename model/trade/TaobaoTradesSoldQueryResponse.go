package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据收件人信息查询交易单号 APIResponse
taobao.trades.sold.query

根据收件人信息查询交易单号。
*/
type TaobaoTradesSoldQueryAPIResponse struct {
    model.CommonResponse
    TaobaoTradesSoldQueryResponse
}

type TaobaoTradesSoldQueryResponse struct {
    XMLName xml.Name `xml:"trades_sold_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 订单ID列表。按照订单创建时间倒序，最多返回最近的100笔订单。
    
    TidList   []string `json:"tid_list,omitempty" xml:"tid_list>string,omitempty"`
    
    
}
