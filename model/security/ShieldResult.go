package security

// ShieldResult 
type ShieldResult struct {

    // 任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
    
    TaskStatus   int64 `json:"task_status,omitempty" xml:"task_status,omitempty"`
    

    // 加固后的应用列表(任务完成时才返回)  普通加固时只有1个文件，多渠道加固时每个渠道1个文件
    
    AppList   []ChannelAppInfo `json:"app_list,omitempty" xml:"app_list,omitempty"`
    

    // 混淆率计算返回结果
    
    ObfuscateResult   *ObfuscateResult `json:"obfuscate_result,omitempty" xml:"obfuscate_result,omitempty"`
    

    // 加固出错的详细信息
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    

    // 混淆后的map文件的地址
    
    MapUrl   string `json:"map_url,omitempty" xml:"map_url,omitempty"`
    

}
