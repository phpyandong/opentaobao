package promotion

// RedRiskUpdateFactor 
type RedRiskUpdateFactor struct {

    // 商品id
    ItemId   int64 `json:"item_id,omitempty"`

    // 需要删除的sku红线价格
    RemoveSkuIds   []Number `json:"remove_sku_ids,omitempty"`

    // 商品红线价格
    AmountAt   int64 `json:"amount_at,omitempty"`

    // 新增sku级别红线价格
    SkuRiskFactors   []SkuRedRiskFactor `json:"sku_risk_factors,omitempty"`

    // 风险等级设置
    RiskLevels   []RiskLevelParam `json:"risk_levels,omitempty"`

}
