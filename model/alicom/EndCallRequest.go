package alicom

// EndCallRequest 
type EndCallRequest struct {

    // 呼叫释放原因,1,未分配的号码（空号） 3,无至目的地的路由 4，停机 6,不可接受的信道 16,正常清除 17,用户忙 18,无用户响应 19,已有用户提醒，但无应答 21,呼叫拒绝 22,号码改变 26,清除未选择的用户 27,终点故障 28,无效号码格式（不完全的号码） 29,设施被拒绝 30,对状态询问的响应 31,正常--未规定 34,无电路/信道可用 38,网络故障 41,临时故障 42,交换设备拥塞 43,接入信息被丢弃 44,请求的电路/信道不可用 47,资源不可用--未规定 49,服务质量不可用 50,未预订所请求的设施 55,IncomingcallsbarredwithintheCUG 57,承载能力未认可(未开通通话功能） 58,承载能力目前不可用 63,无适用的业务或任选项目-未规定 65,承载业务不能实现 68,ACMequaltoorgreaterthanACMmax 69,所请求的设施不能实现 70,仅能获得受限数字信息承载能力 79,业务不能实现-未规定) 81,无效处理识别码 87UsernotmemberofCUG 88,非兼容目的地址 91,无效过渡网选择 95,无效消息-未规定 96,必选消息单元差错 97,消息类型不存在或不能实现 98,消息与控制状态不兼容-消息类型不存在或不能实现 99,信息单元不存在或不能实现 100,无效信息单元内容 101,消息与呼叫状态不兼容 102,定时器超时恢复 111,协议差错-未规定 127,互通-未规定 9999（短信话单时，传此值）
    
    ReleaseCause   int64 `json:"release_cause,omitempty" xml:"release_cause,omitempty"`
    

    // 唯一呼叫ID，需要和转呼控制接口的call_id对应起来；最大可支持字符串长度256
    
    CallId   string `json:"call_id,omitempty" xml:"call_id,omitempty"`
    

    // 被叫响铃时间，如果没有响铃时间，ring_time填写call_out_time的时间；短信话单时，此值传短信接收时间
    
    RingTime   string `json:"ring_time,omitempty" xml:"ring_time,omitempty"`
    

    // 被叫接听时间；短信话单时，此值传短信接收时间
    
    StartTime   string `json:"start_time,omitempty" xml:"start_time,omitempty"`
    

    // 中间号
    
    SecretNo   string `json:"secret_no,omitempty" xml:"secret_no,omitempty"`
    

    // 呼叫被叫侧发起的时间；短信话单时，此值传短信接收时间
    
    CallOutTime   string `json:"call_out_time,omitempty" xml:"call_out_time,omitempty"`
    

    // 0-平台释放 1-主叫释放 2-被叫释放；短信话单时，传0
    
    ReleaseDir   int64 `json:"release_dir,omitempty" xml:"release_dir,omitempty"`
    

    // 通话释放时间,如果未接通start_time应该等于release_time；短信话单时，此值传短信接收时间
    
    ReleaseTime   string `json:"release_time,omitempty" xml:"release_time,omitempty"`
    

    // 阿里侧对应的绑定关系ID，可不传
    
    SubsId   string `json:"subs_id,omitempty" xml:"subs_id,omitempty"`
    

    // 分配给供应商Key
    
    VendorKey   string `json:"vendor_key,omitempty" xml:"vendor_key,omitempty"`
    

    // 被叫空闲振铃时间；可选参数，无此相关功能，可不传
    
    FreeRingTime   string `json:"free_ring_time,omitempty" xml:"free_ring_time,omitempty"`
    

    // 短信条数
    
    SmsNumber   int64 `json:"sms_number,omitempty" xml:"sms_number,omitempty"`
    

    // 录音下载URL,公网可以访问
    
    RecordUrl   string `json:"record_url,omitempty" xml:"record_url,omitempty"`
    

    // 呼叫结果
    
    CallResult   string `json:"call_result,omitempty" xml:"call_result,omitempty"`
    

    // 放音录音下载URL,公网可以访问
    
    RingingRecordUrl   string `json:"ringing_record_url,omitempty" xml:"ringing_record_url,omitempty"`
    

    // 话单类型 0-通话 1-短信
    
    CallType   string `json:"call_type,omitempty" xml:"call_type,omitempty"`
    

    // 主叫号码
    
    CallNo   string `json:"call_no,omitempty" xml:"call_no,omitempty"`
    

    // 被叫号码
    
    CalledNo   string `json:"called_no,omitempty" xml:"called_no,omitempty"`
    

    // 分机号
    
    ExtensionNo   string `json:"extension_no,omitempty" xml:"extension_no,omitempty"`
    

}
