package trade

// AliexpressPaymentExchangeGetResult 
type AliexpressPaymentExchangeGetResult struct {

    // 是否不成功
    
    NotSuccess   bool `json:"not_success,omitempty" xml:"not_success,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 返汇率相关数据
    
    Module   *AliexpressPaymentExchangeGetModule `json:"module,omitempty" xml:"module,omitempty"`
    

    // 错误信息
    
    ErrorCode   *ErrorCode `json:"error_code,omitempty" xml:"error_code,omitempty"`
    

    // 是否重复重复
    
    Repeated   bool `json:"repeated,omitempty" xml:"repeated,omitempty"`
    

    // 是否重试
    
    Retry   bool `json:"retry,omitempty" xml:"retry,omitempty"`
    

}
