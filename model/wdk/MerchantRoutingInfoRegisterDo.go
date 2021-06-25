package wdk

// MerchantRoutingInfoRegisterDo 
type MerchantRoutingInfoRegisterDo struct {

    // 仓code，为空时路由为商家维度
    WarehouseCode   string `json:"warehouse_code,omitempty"`

    // 操作状态1-注册；2-删除
    Status   int64 `json:"status,omitempty"`

}
