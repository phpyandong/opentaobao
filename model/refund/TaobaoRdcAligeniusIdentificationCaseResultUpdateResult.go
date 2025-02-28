package refund

// TaobaoRdcAligeniusIdentificationCaseResultUpdateResult 
type TaobaoRdcAligeniusIdentificationCaseResultUpdateResult struct {

    // resultData
    
    ResultData   *Resultdata `json:"result_data,omitempty" xml:"result_data,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    

    // 错误信息
    
    ErrorInfo   string `json:"error_info,omitempty" xml:"error_info,omitempty"`
    

}
