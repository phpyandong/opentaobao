package servicecenter

// IncomeConfirmDto 
type IncomeConfirmDto struct {

    // 确认金额（单位：分）
    
    Fee   int64 `json:"fee,omitempty" xml:"fee,omitempty"`
    

    // 卖家nick
    
    Nick   string `json:"nick,omitempty" xml:"nick,omitempty"`
    

    // 确认扩展信息
    
    Extend   string `json:"extend,omitempty" xml:"extend,omitempty"`
    

    // appkey
    
    Appkey   string `json:"appkey,omitempty" xml:"appkey,omitempty"`
    

    // 服务市场有效订单ID
    
    OrderId   int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
    

    // 外部订单ID
    
    OutOrderId   string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
    

    // 外部确认账单ID
    
    OutConfirmId   string `json:"out_confirm_id,omitempty" xml:"out_confirm_id,omitempty"`
    

}
