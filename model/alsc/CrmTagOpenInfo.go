package alsc

// CrmTagOpenInfo 
type CrmTagOpenInfo struct {

    // 是否已删除
    
    Deleted   bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
    

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    

    // 修改时间
    
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    

    // 会员计划ID
    
    PlanId   string `json:"plan_id,omitempty" xml:"plan_id,omitempty"`
    

    // 标签ID
    
    TagId   string `json:"tag_id,omitempty" xml:"tag_id,omitempty"`
    

    // 标签名称
    
    TagName   string `json:"tag_name,omitempty" xml:"tag_name,omitempty"`
    

    // 标签类型
    
    TagType   string `json:"tag_type,omitempty" xml:"tag_type,omitempty"`
    

}
