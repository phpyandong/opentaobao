package servicecenter

// SettleAdjustmentResponse 
type SettleAdjustmentResponse struct {

    // bizOrderId
    
    BizOrderId   int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
    

    // comments
    
    Comments   string `json:"comments,omitempty" xml:"comments,omitempty"`
    

    // 费用，单位分
    
    Cost   int64 `json:"cost,omitempty" xml:"cost,omitempty"`
    

    // description
    
    Description   string `json:"description,omitempty" xml:"description,omitempty"`
    

    // gmtCreate
    
    CreateTime   string `json:"create_time,omitempty" xml:"create_time,omitempty"`
    

    // gmtModified
    
    ModifiedTime   string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
    

    // id
    
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    

    // parentBizOrderId
    
    ParentBizOrderId   int64 `json:"parent_biz_order_id,omitempty" xml:"parent_biz_order_id,omitempty"`
    

    // pictureUrls
    
    PictureUrls   string `json:"picture_urls,omitempty" xml:"picture_urls,omitempty"`
    

    // priceFactors
    
    PriceFactors   string `json:"price_factors,omitempty" xml:"price_factors,omitempty"`
    

    // serviceCode
    
    ServiceCode   string `json:"service_code,omitempty" xml:"service_code,omitempty"`
    

    // serviceOrderId
    
    ServiceOrderId   int64 `json:"service_order_id,omitempty" xml:"service_order_id,omitempty"`
    

    // workcardId
    
    WorkcardId   int64 `json:"workcard_id,omitempty" xml:"workcard_id,omitempty"`
    

    // 调整单态。 待商家确认:1, 商家已确认:2,  待小二判定:3,  小二判定有效:4,  小二判定无效:5,  小二无法判定:6, 服务商取消:7, 超时确认:8, 完成:9
    
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    

    // type
    
    Type   int64 `json:"type,omitempty" xml:"type,omitempty"`
    

}
