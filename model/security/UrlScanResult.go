package security

// UrlScanResult 
type UrlScanResult struct {

    // 请求标志唯一id
    
    EventId   string `json:"event_id,omitempty" xml:"event_id,omitempty"`
    

    // 扫描详细结果
    
    UrlResultItemList   []UrlScanResultItem `json:"url_result_item_list,omitempty" xml:"url_result_item_list,omitempty"`
    

}
