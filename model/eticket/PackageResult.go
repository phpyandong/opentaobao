package eticket

// PackageResult 
type PackageResult struct {

    // 操作结果码
    
    Code   int64 `json:"code,omitempty" xml:"code,omitempty"`
    

    // 操作是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 操作结果信息
    
    Info   string `json:"info,omitempty" xml:"info,omitempty"`
    

    // 包基本信息
    
    PackageBase   *PackageBase `json:"package_base,omitempty" xml:"package_base,omitempty"`
    

    // 包基本信息列表
    
    PackageBaseList   []PackageBase `json:"package_base_list,omitempty" xml:"package_base_list,omitempty"`
    

}
