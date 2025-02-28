package trade

// LogisticsInfo 
type LogisticsInfo struct {

    // 交易号
    
    TradeId   int64 `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
    

    // 子交易号
    
    SubTradeId   int64 `json:"sub_trade_id,omitempty" xml:"sub_trade_id,omitempty"`
    

    // 货品仓储ID
    
    ItemId   string `json:"item_id,omitempty" xml:"item_id,omitempty"`
    

    // 货品仓储code
    
    ItemCode   string `json:"item_code,omitempty" xml:"item_code,omitempty"`
    

    // 应发数量
    
    NeedConsignNum   int64 `json:"need_consign_num,omitempty" xml:"need_consign_num,omitempty"`
    

    // 如是菜鸟仓，则将菜鸟仓的区域仓code进行填充，如是商家仓发货则填充商家仓code
    
    StoreCode   string `json:"store_code,omitempty" xml:"store_code,omitempty"`
    

    // 子订单类型:标示该子交易单来源交易，还是BMS增加的，枚举值(00=交易，10=BMS绑定)
    
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
    

    // 商品数字编号
    
    NumIid   int64 `json:"num_iid,omitempty" xml:"num_iid,omitempty"`
    

    // 发货类型CN=菜鸟发货，SC的商家仓发货
    
    ConsignType   string `json:"consign_type,omitempty" xml:"consign_type,omitempty"`
    

    // 商品的最小库存单位Sku的id
    
    SkuId   string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
    

    // 商品比例
    
    ItemRatio   int64 `json:"item_ratio,omitempty" xml:"item_ratio,omitempty"`
    

    // 组合商品编码code
    
    CombineItemCode   string `json:"combine_item_code,omitempty" xml:"combine_item_code,omitempty"`
    

    // 组合商品id
    
    CombineItemId   string `json:"combine_item_id,omitempty" xml:"combine_item_id,omitempty"`
    

}
