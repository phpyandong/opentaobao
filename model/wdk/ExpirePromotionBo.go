package wdk

// ExpirePromotionBo 
type ExpirePromotionBo struct {

    // 短保时间段信息
    
    PeriodInfos   []ExpirePeriodInfo `json:"period_infos,omitempty" xml:"period_infos,omitempty"`
    

    // 门店id
    
    ShopIds   []string `json:"shop_ids,omitempty" xml:"shop_ids>string,omitempty"`
    

    // 商家编码
    
    MerchantCode   string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
    

    // 商品skucode
    
    SkuCode   string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
    

}
