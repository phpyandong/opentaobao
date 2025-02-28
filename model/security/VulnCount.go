package security

// VulnCount 
type VulnCount struct {

    // 漏洞总数量
    
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
    

    // 高风险漏洞数量
    
    HighLevel   int64 `json:"high_level,omitempty" xml:"high_level,omitempty"`
    

    // 中风险漏洞数量
    
    MidLevel   int64 `json:"mid_level,omitempty" xml:"mid_level,omitempty"`
    

    // 低风险漏洞数量
    
    LowLevel   int64 `json:"low_level,omitempty" xml:"low_level,omitempty"`
    

    // 安全红线漏洞数量
    
    RedLine   int64 `json:"red_line,omitempty" xml:"red_line,omitempty"`
    

}
