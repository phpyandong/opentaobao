package dt

// PriceTagReqParam 
type PriceTagReqParam struct {

    // 图片名称，不能重复
    
    ObjKeyName   string `json:"obj_key_name,omitempty" xml:"obj_key_name,omitempty"`
    

    // 业务编码
    
    BusiCode   string `json:"busi_code,omitempty" xml:"busi_code,omitempty"`
    

    // 业务来源
    
    Source   string `json:"source,omitempty" xml:"source,omitempty"`
    

    // 条码
    
    BarCode   string `json:"bar_code,omitempty" xml:"bar_code,omitempty"`
    

    // 扩展信息
    
    ExtendInfoMap   string `json:"extend_info_map,omitempty" xml:"extend_info_map,omitempty"`
    

}
