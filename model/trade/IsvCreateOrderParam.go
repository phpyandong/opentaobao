package trade

// IsvCreateOrderParam 
type IsvCreateOrderParam struct {

    // 买家备注
    
    BuyerRemarks   string `json:"buyer_remarks,omitempty" xml:"buyer_remarks,omitempty"`
    

    // 外部订单ID，仅仅做追踪使用
    
    OutOrderId   string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
    

    // 收货地址信息
    
    SmAddrModel   *SmAddrModel `json:"sm_addr_model,omitempty" xml:"sm_addr_model,omitempty"`
    

    // 订单行详情
    
    SubOrderInfoList   []IsvSimpleSubOrderModel `json:"sub_order_info_list,omitempty" xml:"sub_order_info_list,omitempty"`
    

}
