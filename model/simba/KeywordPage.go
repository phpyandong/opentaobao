package simba

// KeywordPage 
type KeywordPage struct {

    // 每页数据大小
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    

    // 返回第几页的数据从1开始
    
    PageNo   int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
    

    // 所查询的数据总数，只有当返回第一页数据时有值，当要求返回的数据非第一页时，此返回值无效
    
    TotalItem   int64 `json:"total_item,omitempty" xml:"total_item,omitempty"`
    

    // 关键词列表
    
    KeywordList   []Keyword `json:"keyword_list,omitempty" xml:"keyword_list,omitempty"`
    

}
