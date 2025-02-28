package wms

// Invoicewlbwmsconsignordernotify 
type Invoicewlbwmsconsignordernotify struct {

    // Erp发票ID
    
    BillId   int64 `json:"bill_id,omitempty" xml:"bill_id,omitempty"`
    

    // 发票类型(VINVOICE - 增值税普通发票， EVINVOICE - 电子增票，电子发票仓库不打印)
    
    BillType   string `json:"bill_type,omitempty" xml:"bill_type,omitempty"`
    

    // 发票抬头
    
    BillTitle   string `json:"bill_title,omitempty" xml:"bill_title,omitempty"`
    

    // 发票金额
    
    BillAccount   string `json:"bill_account,omitempty" xml:"bill_account,omitempty"`
    

    // 发票内容
    
    BillContent   string `json:"bill_content,omitempty" xml:"bill_content,omitempty"`
    

    // 发票明细列表
    
    DetailList   []Detaillistwlbwmsconsignordernotify `json:"detail_list,omitempty" xml:"detail_list,omitempty"`
    

}
