package category

// ItemPickPagingResult 
type ItemPickPagingResult struct {

    // result
    
    Results   []CategoryDto `json:"results,omitempty" xml:"results,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    

    // 错误信息
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    

}
