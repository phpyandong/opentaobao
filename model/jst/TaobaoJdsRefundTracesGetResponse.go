package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单条退款跟踪详情 APIResponse
taobao.jds.refund.traces.get

获取聚石塔数据共享的交易全链路的退款信息
*/
type TaobaoJdsRefundTracesGetAPIResponse struct {
    model.CommonResponse
    TaobaoJdsRefundTracesGetResponse
}

type TaobaoJdsRefundTracesGetResponse struct {
    XMLName xml.Name `xml:"jds_refund_traces_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 退款跟踪列表
    
    Traces   []RefundTrace `json:"traces,omitempty" xml:"traces>refund_trace,omitempty"`
    
    
    // 用户在全链路系统中的状态(比如是否存在)
    
    UserStatus   string `json:"user_status,omitempty" xml:"user_status,omitempty"`

    
}
