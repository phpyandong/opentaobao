package travel

// TopElementParam 
type TopElementParam struct {

    // 资源元素类型。1-景点，2-酒店，999-其他
    
    ElementType   int64 `json:"element_type,omitempty" xml:"element_type,omitempty"`
    

    // 元素所在城市
    
    City   string `json:"city,omitempty" xml:"city,omitempty"`
    

    // 元素名称
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 元素的子类型
    
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
    

    // 元素的外部商家编码
    
    OuterId   string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
    

    // 元素的说明描述
    
    Desc   string `json:"desc,omitempty" xml:"desc,omitempty"`
    

}
