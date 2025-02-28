package product

// SupplyItemChangeMessage 
type SupplyItemChangeMessage struct {

    // 货号列表
    
    ProductCodes   []string `json:"product_codes,omitempty" xml:"product_codes>string,omitempty"`
    

    // 供应商门店，最大20个
    
    SupplyStoreId   string `json:"supply_store_id,omitempty" xml:"supply_store_id,omitempty"`
    

}
