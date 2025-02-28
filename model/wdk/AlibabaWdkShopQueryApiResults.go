package wdk

// AlibabaWdkShopQueryApiResults 
type AlibabaWdkShopQueryApiResults struct {

    // 返回门店信息列表
    
    Models   []ShopDo `json:"models,omitempty" xml:"models,omitempty"`
    

    // 异常编码
    
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    

    // 异常信息
    
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    

    // true
    
    Success   string `json:"success,omitempty" xml:"success,omitempty"`
    

}
