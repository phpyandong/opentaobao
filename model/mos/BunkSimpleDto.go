package mos

// BunkSimpleDto 
type BunkSimpleDto struct {

    // 门店号
    
    StoreNo   string `json:"store_no,omitempty" xml:"store_no,omitempty"`
    

    // 铺位面积
    
    Acreage   string `json:"acreage,omitempty" xml:"acreage,omitempty"`
    

    // 铺位类型
    
    Type   int64 `json:"type,omitempty" xml:"type,omitempty"`
    

    // 铺位编号
    
    Number   string `json:"number,omitempty" xml:"number,omitempty"`
    

    // 铺位ID
    
    CpId   string `json:"cp_id,omitempty" xml:"cp_id,omitempty"`
    

    // 合同号
    
    ContractCode   string `json:"contract_code,omitempty" xml:"contract_code,omitempty"`
    

    // 楼层名称
    
    FloorName   string `json:"floor_name,omitempty" xml:"floor_name,omitempty"`
    

    // 楼层ID
    
    FloorId   string `json:"floor_id,omitempty" xml:"floor_id,omitempty"`
    

    // 合同系统编号
    
    ContractId   string `json:"contract_id,omitempty" xml:"contract_id,omitempty"`
    

}
