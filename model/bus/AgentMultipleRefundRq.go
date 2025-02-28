package bus

// AgentMultipleRefundRq 
type AgentMultipleRefundRq struct {

    // 商家订单号
    
    AgentOrderId   string `json:"agent_order_id,omitempty" xml:"agent_order_id,omitempty"`
    

    // 退款原因必填
    
    AgentRefundReason   string `json:"agent_refund_reason,omitempty" xml:"agent_refund_reason,omitempty"`
    

    // 退款流水号唯一标识，防止重复退款
    
    AgentRefundTransNo   string `json:"agent_refund_trans_no,omitempty" xml:"agent_refund_trans_no,omitempty"`
    

    // 飞猪平台订单号
    
    MainBizOrderId   int64 `json:"main_biz_order_id,omitempty" xml:"main_biz_order_id,omitempty"`
    

    // 退款明细数据
    
    RefundTicketDetailList   []AgentMultipleRefundTicketInfo `json:"refund_ticket_detail_list,omitempty" xml:"refund_ticket_detail_list,omitempty"`
    

    // 退款总金额(票款+服务费)(分)
    
    TotalRefundAmount   int64 `json:"total_refund_amount,omitempty" xml:"total_refund_amount,omitempty"`
    

    // 退服务费总金额(分)
    
    TotalServiceChargeRefundAmount   int64 `json:"total_service_charge_refund_amount,omitempty" xml:"total_service_charge_refund_amount,omitempty"`
    

}
