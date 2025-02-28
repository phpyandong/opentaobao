package simba

// TagOptions 
type TagOptions struct {

    // 标签的名称,如男,女,10-20等
    
    TagName   string `json:"tag_name,omitempty" xml:"tag_name,omitempty"`
    

    // 标签所属的分组id
    
    OptionGroupId   int64 `json:"option_group_id,omitempty" xml:"option_group_id,omitempty"`
    

    // 标签所属的维度id
    
    DimId   int64 `json:"dim_id,omitempty" xml:"dim_id,omitempty"`
    

    // 标签id
    
    TagId   string `json:"tag_id,omitempty" xml:"tag_id,omitempty"`
    

}
