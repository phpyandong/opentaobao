package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
物流流转信息查询 APIResponse
taobao.logistics.trace.search

用户根据淘宝交易号查询物流流转信息，如2010-8-10 15：23：00到达杭州集散地。<br/>此接口的返回信息都由物流公司提供。（备注：使用线下发货（offline.send）的运单，不支持运单状态的实时跟踪，只要一发货，状态就会变为<status>对方已签收</status>，该字段仅对线上发货（online.send）的运单有效。）
*/
type TaobaoLogisticsTraceSearchAPIResponse struct {
    model.CommonResponse
    TaobaoLogisticsTraceSearchResponse
}

type TaobaoLogisticsTraceSearchResponse struct {
    XMLName xml.Name `xml:"logistics_trace_search_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 运单号
    
    OutSid   string `json:"out_sid,omitempty" xml:"out_sid,omitempty"`

    
    // 物流公司名称
    
    CompanyName   string `json:"company_name,omitempty" xml:"company_name,omitempty"`

    
    // 交易号
    
    Tid   int64 `json:"tid,omitempty" xml:"tid,omitempty"`

    
    // 订单的物流状态（仅支持线上发货online订单，线下发货offline发出后直接变为已签收）
* 等候发送给物流公司
*已提交给物流公司,等待物流公司接单
*已经确认消息接收，等待物流公司接单
*物流公司已接单
*物流公司不接单
*物流公司揽收失败
*物流公司揽收成功
*签收失败
*对方已签收
*对方拒绝签收
    
    Status   string `json:"status,omitempty" xml:"status,omitempty"`

    
    // 流转信息列表
    
    TraceList   []TransitStepInfo `json:"trace_list,omitempty" xml:"trace_list>transit_step_info,omitempty"`
    
    
}
