package icbu

// SkuAttributeValue 
type SkuAttributeValue struct {

    // 默认的属性值名称
    
    SystemValueName   string `json:"system_value_name,omitempty" xml:"system_value_name,omitempty"`
    

    // 自定义的图片URL
    
    ImageUrl   string `json:"image_url,omitempty" xml:"image_url,omitempty"`
    

    // 默认的展示样式
    
    MarkInfo   string `json:"mark_info,omitempty" xml:"mark_info,omitempty"`
    

    // 属性值ID
    
    ValueId   int64 `json:"value_id,omitempty" xml:"value_id,omitempty"`
    

    // 自定义的属性值名称
    
    CustomValueName   string `json:"custom_value_name,omitempty" xml:"custom_value_name,omitempty"`
    

}
