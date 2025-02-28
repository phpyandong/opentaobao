package alsc

// LevelConsumeGrowRuleOpenInfo 
type LevelConsumeGrowRuleOpenInfo struct {

    // 创建人
    
    CreateBy   string `json:"create_by,omitempty" xml:"create_by,omitempty"`
    

    // 是否已经删除
    
    Deleted   bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
    

    // 扩展字段
    
    ExtInfo   string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
    

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    

    // 修改时间
    
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    

    // 等级ID
    
    LevelId   string `json:"level_id,omitempty" xml:"level_id,omitempty"`
    

    // 等级名称
    
    LevelName   string `json:"level_name,omitempty" xml:"level_name,omitempty"`
    

    // 等级编码
    
    LevelNo   string `json:"level_no,omitempty" xml:"level_no,omitempty"`
    

    // 每消费金额，单位：分
    
    PerConsume   int64 `json:"per_consume,omitempty" xml:"per_consume,omitempty"`
    

    // 每消费金额对应可获得的成长值
    
    PerGrowth   int64 `json:"per_growth,omitempty" xml:"per_growth,omitempty"`
    

    // 更新人
    
    UpdateBy   string `json:"update_by,omitempty" xml:"update_by,omitempty"`
    

}
