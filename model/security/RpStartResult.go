package security

// RpStartResult 
type RpStartResult struct {

    // biz
    
    Biz   string `json:"biz,omitempty" xml:"biz,omitempty"`
    

    // extraInfo
    
    ExtraInfo   string `json:"extra_info,omitempty" xml:"extra_info,omitempty"`
    

    // source
    
    Source   string `json:"source,omitempty" xml:"source,omitempty"`
    

    // steps
    
    Steps   []RpStepItem `json:"steps,omitempty" xml:"steps,omitempty"`
    

    // uploadToken
    
    UploadToken   *StsUploadToken `json:"upload_token,omitempty" xml:"upload_token,omitempty"`
    

}
