package wdk

// OutputItemInfoDto 
type OutputItemInfoDto struct {

    // 数量
    
    Count   string `json:"count,omitempty" xml:"count,omitempty"`
    

    // 部门code，该商品所属的部门编码
    
    DeptCode   string `json:"dept_code,omitempty" xml:"dept_code,omitempty"`
    

    // 库存单位
    
    InventoryUnit   string `json:"inventory_unit,omitempty" xml:"inventory_unit,omitempty"`
    

    // 商品code，盒马系统中的商品编码
    
    ItemCode   string `json:"item_code,omitempty" xml:"item_code,omitempty"`
    

    // 商品价格，单位为分
    
    Price   string `json:"price,omitempty" xml:"price,omitempty"`
    

    // 原因
    
    Reason   string `json:"reason,omitempty" xml:"reason,omitempty"`
    

    // 备注
    
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    

    // 规格
    
    Spec   string `json:"spec,omitempty" xml:"spec,omitempty"`
    

    // 采购单位
    
    Unit   string `json:"unit,omitempty" xml:"unit,omitempty"`
    

}
