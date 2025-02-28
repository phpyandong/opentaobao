package scbp

// CountryTagView 
type CountryTagView struct {

    // 最近7天效果数据
    
    Effect   *Effect7d `json:"effect,omitempty" xml:"effect,omitempty"`
    

    // 溢价百分比
    
    Discount   int64 `json:"discount,omitempty" xml:"discount,omitempty"`
    

    // 标签中文名
    
    TagName   string `json:"tag_name,omitempty" xml:"tag_name,omitempty"`
    

    // 标签id
    
    TagId   string `json:"tag_id,omitempty" xml:"tag_id,omitempty"`
    

}
