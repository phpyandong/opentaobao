package qimen

// SnReportRequest 
type SnReportRequest struct {

    // 总页数
    
    TotalPage   int64 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
    

    // 当前页(从1开始)
    
    CurrentPage   int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
    

    // 每页记录的条数
    
    PageSize   int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
    

    // 发货单信息
    
    DeliveryOrder   *DeliveryOrder `json:"deliveryOrder,omitempty" xml:"deliveryOrder,omitempty"`
    

    // 商品列表
    
    Items   []Item `json:"items,omitempty" xml:"items,omitempty"`
    

    // 扩展属性
    
    ExtendProps   *Map `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
    

}
