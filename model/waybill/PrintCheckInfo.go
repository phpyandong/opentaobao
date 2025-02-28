package waybill

// PrintCheckInfo 
type PrintCheckInfo struct {

    // 发货网点编码
    
    ShippingBranchCode   string `json:"shipping_branch_code,omitempty" xml:"shipping_branch_code,omitempty"`
    

    // 收件人姓名
    
    ConsigneeName   string `json:"consignee_name,omitempty" xml:"consignee_name,omitempty"`
    

    // 发件人联系方式
    
    SendPhone   string `json:"send_phone,omitempty" xml:"send_phone,omitempty"`
    

    // 包裹重量 单位为G(克)
    
    Weight   int64 `json:"weight,omitempty" xml:"weight,omitempty"`
    

    // 电子面单单号
    
    WaybillCode   string `json:"waybill_code,omitempty" xml:"waybill_code,omitempty"`
    

    // 收\发货地址
    
    ConsigneeAddress   *WaybillAddress `json:"consignee_address,omitempty" xml:"consignee_address,omitempty"`
    

    // 快递服务产品类型编码
    
    ProductType   string `json:"product_type,omitempty" xml:"product_type,omitempty"`
    

    // 发件人姓名
    
    SendName   string `json:"send_name,omitempty" xml:"send_name,omitempty"`
    

    // 收货网点编码
    
    ConsigneeBranchCode   string `json:"consignee_branch_code,omitempty" xml:"consignee_branch_code,omitempty"`
    

    // 物流服务能力集合
    
    LogisticsServiceList   []LogisticsService `json:"logistics_service_list,omitempty" xml:"logistics_service_list,omitempty"`
    

    // 收货网点信息
    
    ConsigneeBranchName   string `json:"consignee_branch_name,omitempty" xml:"consignee_branch_name,omitempty"`
    

    // 发货网点信息
    
    ShippingBranchName   string `json:"shipping_branch_name,omitempty" xml:"shipping_branch_name,omitempty"`
    

    // 拣货规则（大头笔信息）
    
    ShortAddress   string `json:"short_address,omitempty" xml:"short_address,omitempty"`
    

    // 包裹体积 单位为ML(毫升)或立方厘米
    
    Volume   int64 `json:"volume,omitempty" xml:"volume,omitempty"`
    

    // consigneePhone
    
    ConsigneePhone   string `json:"consignee_phone,omitempty" xml:"consignee_phone,omitempty"`
    

    // 收\发货地址
    
    ShippingAddress   *WaybillAddress `json:"shipping_address,omitempty" xml:"shipping_address,omitempty"`
    

    // 使用者ID
    
    RealUserId   int64 `json:"real_user_id,omitempty" xml:"real_user_id,omitempty"`
    

    // 集包地、目的地中心代码。打 印时根据该 code 生成目的地 中心的条码，条码生成的算法 与对应的电子面单条码一致
    
    PackageCenterCode   string `json:"package_center_code,omitempty" xml:"package_center_code,omitempty"`
    

    // 集包地、目的地中心名称
    
    PackageCenterName   string `json:"package_center_name,omitempty" xml:"package_center_name,omitempty"`
    

    // 打标设置字段，直接传给ali-lodop。不用管具体含义。
    
    PrintConfig   string `json:"print_config,omitempty" xml:"print_config,omitempty"`
    

}
