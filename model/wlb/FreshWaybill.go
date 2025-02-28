package wlb

// FreshWaybill 
type FreshWaybill struct {

    // 获取的所有电子面单号，以“;”分隔
    
    WaybillCode   string `json:"waybill_code,omitempty" xml:"waybill_code,omitempty"`
    

    // 大头笔
    
    ShortAddress   string `json:"short_address,omitempty" xml:"short_address,omitempty"`
    

    // 简称
    
    Alias   string `json:"alias,omitempty" xml:"alias,omitempty"`
    

    // 预计到达时间
    
    Time   string `json:"time,omitempty" xml:"time,omitempty"`
    

    // 预留扩展字段
    
    Feature   string `json:"feature,omitempty" xml:"feature,omitempty"`
    

    // 交易号
    
    TradeId   string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
    

}
