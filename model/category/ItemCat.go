package category

// ItemCat 
type ItemCat struct {

    // 商品所属类目ID
    
    Cid   int64 `json:"cid,omitempty" xml:"cid,omitempty"`
    

    // 父类目ID=0时，代表的是一级的类目
    
    ParentCid   int64 `json:"parent_cid,omitempty" xml:"parent_cid,omitempty"`
    

    // 类目名称
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 该类目是否为父类目(即：该类目是否还有子类目)
    
    IsParent   bool `json:"is_parent,omitempty" xml:"is_parent,omitempty"`
    

    // 状态。可选值:normal(正常),deleted(删除)
    
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    

    // 排列序号，表示同级类目的展现次序，如数值相等则按名称次序排列。取值范围:大于零的整数
    
    SortOrder   int64 `json:"sort_order,omitempty" xml:"sort_order,omitempty"`
    

    // Feature对象列表<br/>目前已有的属性：<br/>若Attr_key为 udsaleprop，attr_value为1 则允许卖家在改类目新增自定义销售属性,不然为不允许
    
    Features   []Feature `json:"features,omitempty" xml:"features,omitempty"`
    

    // 是否度量衡类目
    
    TaosirCat   bool `json:"taosir_cat,omitempty" xml:"taosir_cat,omitempty"`
    

}
