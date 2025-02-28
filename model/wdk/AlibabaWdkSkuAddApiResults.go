package wdk

// AlibabaWdkSkuAddApiResults 
type AlibabaWdkSkuAddApiResults struct {

    // 错误码
    
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    

    // 错误信息
    
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    

    // 接口返回成功标志
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // models
    
    Models   []AlibabaWdkSkuAddApiResult `json:"models,omitempty" xml:"models,omitempty"`
    

}
