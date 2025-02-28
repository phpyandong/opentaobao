package wdk

// WdkOpenOrderFinanceBillQueryRequest 
type WdkOpenOrderFinanceBillQueryRequest struct {

    // 分页查询数量
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    

    // 传入上一次查询结果的next_id，第一次查询时传0
    
    NextId   int64 `json:"next_id,omitempty" xml:"next_id,omitempty"`
    

    // 经营店id
    
    StoreId   string `json:"store_id,omitempty" xml:"store_id,omitempty"`
    

    // 账单日期
    
    Dt   string `json:"dt,omitempty" xml:"dt,omitempty"`
    

    // 售卖商家code，联营商模式必填，非联营商模式不填
    
    SellerMerchantCode   string `json:"seller_merchant_code,omitempty" xml:"seller_merchant_code,omitempty"`
    

}
