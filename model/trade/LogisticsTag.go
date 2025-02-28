package trade

// LogisticsTag 
type LogisticsTag struct {

    // 主订单或子订单的订单号
    
    OrderId   string `json:"order_id,omitempty" xml:"order_id,omitempty"`
    

    // 服务标签
    
    LogisticServiceTagList   []LogisticServiceTag `json:"logistic_service_tag_list,omitempty" xml:"logistic_service_tag_list,omitempty"`
    

}
