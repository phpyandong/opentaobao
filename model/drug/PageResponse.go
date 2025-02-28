package drug

// PageResponse 
type PageResponse struct {

    // 当前页码
    
    CurrentPage   int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
    

    // 总页数
    
    TotalPage   int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
    

    // 总条数
    
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
    

    // 结果列表
    
    Spus   []TopAlihealthSpuQuery `json:"spus,omitempty" xml:"spus,omitempty"`
    

}
