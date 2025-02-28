package wdk

// MaochaoOrderInsuranceRefundQueryResult 
type MaochaoOrderInsuranceRefundQueryResult struct {

    // 是否调用成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 返回码
    
    ReturnCode   string `json:"return_code,omitempty" xml:"return_code,omitempty"`
    

    // 返回码说明
    
    ReturnMsg   string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
    

    // 返回结果
    
    Model   *InsuranceRefundOrder `json:"model,omitempty" xml:"model,omitempty"`
    

}
