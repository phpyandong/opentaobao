package tmallservice

// TmallMallitemcenterServiceproductQueryResult 
type TmallMallitemcenterServiceproductQueryResult struct {

    // 返回数据
    
    ResultDataList   []Resultdata `json:"result_data_list,omitempty" xml:"result_data_list,omitempty"`
    

    // message
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    

    // 系统是否异常
    
    SystemError   bool `json:"system_error,omitempty" xml:"system_error,omitempty"`
    

    // 业务校验是否正常
    
    BusinessCheckFail   bool `json:"business_check_fail,omitempty" xml:"business_check_fail,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
