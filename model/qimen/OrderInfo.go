package qimen

// OrderInfo 
type OrderInfo struct {

    // 订单创建时间(YYYY-MM-DD HH:MM:SS)
    
    CreateTime   string `json:"createTime,omitempty" xml:"createTime,omitempty"`
    

    // 出库单仓储系统编码
    
    DeliveryOrderId   string `json:"deliveryOrderId,omitempty" xml:"deliveryOrderId,omitempty"`
    

    // 仓库编码(统仓统配使用)
    
    WarehouseCode   string `json:"warehouseCode,omitempty" xml:"warehouseCode,omitempty"`
    

    // 物流公司编码(统仓统配使用)
    
    LogisticsCode   string `json:"logisticsCode,omitempty" xml:"logisticsCode,omitempty"`
    

}
