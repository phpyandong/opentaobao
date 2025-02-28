package fenxiao

// ResultDTO 
type ResultDTO struct {

    // 错误码
    
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    

    // 执行结果
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
