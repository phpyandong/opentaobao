package alihealth

// ServiceResult 
type ServiceResult struct {

    // errMessage
    
    ErrMessage   string `json:"err_message,omitempty" xml:"err_message,omitempty"`
    

    // token
    
    Data   string `json:"data,omitempty" xml:"data,omitempty"`
    

    // errCode
    
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
