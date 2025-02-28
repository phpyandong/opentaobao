package baodian

// GameGiftRecordDto 
type GameGiftRecordDto struct {

    // 记录id
    
    RecordId   int64 `json:"record_id,omitempty" xml:"record_id,omitempty"`
    

    // 记录状态
    
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    

    // cp item id
    
    CpItemId   string `json:"cp_item_id,omitempty" xml:"cp_item_id,omitempty"`
    

}
