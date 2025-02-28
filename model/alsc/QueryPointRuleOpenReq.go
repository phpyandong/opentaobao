package alsc

// QueryPointRuleOpenReq 
type QueryPointRuleOpenReq struct {

    // 品牌id(不能和outbrandid同时为空)
    
    BrandId   string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
    

    // 是否包含逻辑删除
    
    IncludeLogicalDelete   bool `json:"include_logical_delete,omitempty" xml:"include_logical_delete,omitempty"`
    

    // 会员等级ID
    
    LevelId   string `json:"level_id,omitempty" xml:"level_id,omitempty"`
    

    // maxUpdateTime
    
    MaxUpdateTime   string `json:"max_update_time,omitempty" xml:"max_update_time,omitempty"`
    

    // 外部品牌ID(不能和brandid同时为空)
    
    OutBrandId   string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
    

    // 外部门店ID(不能和shopid同时为空)
    
    OutShopId   string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
    

    // 门店ID(不能和outshopid同时为空)
    
    ShopId   string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
    

}
