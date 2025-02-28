package mos

// Product 
type Product struct {

    // 分摊金额
    
    Amount   string `json:"amount,omitempty" xml:"amount,omitempty"`
    

    // 支付方式行号
    
    Fkfsorder   string `json:"fkfsorder,omitempty" xml:"fkfsorder,omitempty"`
    

    // 订单号或券号（支付宝订单号）
    
    Orderid   string `json:"orderid,omitempty" xml:"orderid,omitempty"`
    

    // 商品行号
    
    Comorder   string `json:"comorder,omitempty" xml:"comorder,omitempty"`
    

    // 大支付方式
    
    Payment   string `json:"payment,omitempty" xml:"payment,omitempty"`
    

    // 小支付方式
    
    Subpayment   string `json:"subpayment,omitempty" xml:"subpayment,omitempty"`
    

}
