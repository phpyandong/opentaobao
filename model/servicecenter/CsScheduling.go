package servicecenter

// CsScheduling 
type CsScheduling struct {

    // 排班日期
    
    Date   string `json:"date,omitempty" xml:"date,omitempty"`
    

    // 排班日期字符串
    
    StringDate   string `json:"string_date,omitempty" xml:"string_date,omitempty"`
    

    // 一天内排班信息
    
    Schedulings   []Scheduling `json:"schedulings,omitempty" xml:"schedulings,omitempty"`
    

    // 订单ID
    
    OrderId   int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
    

    // 排班记录更新时间
    
    ModifiedTime   string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
    

}
