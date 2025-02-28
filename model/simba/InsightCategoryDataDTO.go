package simba

// InsightCategoryDataDTO 
type InsightCategoryDataDTO struct {

    // 展现量
    
    Impression   int64 `json:"impression,omitempty" xml:"impression,omitempty"`
    

    // 点击量
    
    Click   int64 `json:"click,omitempty" xml:"click,omitempty"`
    

    // 点击率
    
    Ctr   string `json:"ctr,omitempty" xml:"ctr,omitempty"`
    

    // 类目id
    
    CatId   int64 `json:"cat_id,omitempty" xml:"cat_id,omitempty"`
    

    // 类目名称
    
    CatName   string `json:"cat_name,omitempty" xml:"cat_name,omitempty"`
    

}
