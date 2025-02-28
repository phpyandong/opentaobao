package security

// RpClientInfo 
type RpClientInfo struct {

    // os版本号
    
    OsVersion   string `json:"os_version,omitempty" xml:"os_version,omitempty"`
    

    // umid_token
    
    UmidToken   string `json:"umid_token,omitempty" xml:"umid_token,omitempty"`
    

    // 无线端用于风控的token
    
    WuaToken   string `json:"wua_token,omitempty" xml:"wua_token,omitempty"`
    

    // ip地址
    
    Ip   string `json:"ip,omitempty" xml:"ip,omitempty"`
    

    // 时间戳
    
    TimestampInfo   string `json:"timestamp_info,omitempty" xml:"timestamp_info,omitempty"`
    

    // 无线端用于风控的wua数据
    
    Wua   string `json:"wua,omitempty" xml:"wua,omitempty"`
    

    // 应用名
    
    AppName   string `json:"app_name,omitempty" xml:"app_name,omitempty"`
    

    // 应用版本
    
    AppVersion   string `json:"app_version,omitempty" xml:"app_version,omitempty"`
    

    // session_id
    
    SessionId   string `json:"session_id,omitempty" xml:"session_id,omitempty"`
    

    // 客户端OS名称
    
    OsName   string `json:"os_name,omitempty" xml:"os_name,omitempty"`
    

    // 原始设备号
    
    DeviceId   string `json:"device_id,omitempty" xml:"device_id,omitempty"`
    

    // 客户端类别
    
    ClientType   string `json:"client_type,omitempty" xml:"client_type,omitempty"`
    

    // appKey
    
    AppKeyInfo   string `json:"app_key_info,omitempty" xml:"app_key_info,omitempty"`
    

    // 认证SDK版本号
    
    RpSdkVersion   string `json:"rp_sdk_version,omitempty" xml:"rp_sdk_version,omitempty"`
    

    // 认证SDK名字
    
    RpSdkName   string `json:"rp_sdk_name,omitempty" xml:"rp_sdk_name,omitempty"`
    

    // xxx
    
    ExtendMap   string `json:"extend_map,omitempty" xml:"extend_map,omitempty"`
    

    // 活体sdk版本
    
    LivenessSdkVersion   string `json:"liveness_sdk_version,omitempty" xml:"liveness_sdk_version,omitempty"`
    

    // 活体sdk名字
    
    LivenessSdkName   string `json:"liveness_sdk_name,omitempty" xml:"liveness_sdk_name,omitempty"`
    

    // 制造商
    
    Manufacturer   string `json:"manufacturer,omitempty" xml:"manufacturer,omitempty"`
    

    // 手机型号
    
    MobileModel   string `json:"mobile_model,omitempty" xml:"mobile_model,omitempty"`
    

    // 架构
    
    CpuArch   string `json:"cpu_arch,omitempty" xml:"cpu_arch,omitempty"`
    

    // 总内存
    
    TotalMemory   string `json:"total_memory,omitempty" xml:"total_memory,omitempty"`
    

    // 剩余内存
    
    AvailableMemory   string `json:"available_memory,omitempty" xml:"available_memory,omitempty"`
    

    // 支持
    
    SupportNeon   string `json:"support_neon,omitempty" xml:"support_neon,omitempty"`
    

    // umid
    
    Umid   string `json:"umid,omitempty" xml:"umid,omitempty"`
    

}
