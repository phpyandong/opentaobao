package scbp

// TopP4pQuickProductQuery 
type TopP4pQuickProductQuery struct {

    // 定向推广计划id
    
    CampaignId   int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
    

    // 第几页
    
    ToPage   int64 `json:"to_page,omitempty" xml:"to_page,omitempty"`
    

    // 每页返回多少行
    
    PerPageSize   int64 `json:"per_page_size,omitempty" xml:"per_page_size,omitempty"`
    

}
