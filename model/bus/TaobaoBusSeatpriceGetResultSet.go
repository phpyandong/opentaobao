package bus

// TaobaoBusSeatpriceGetResultSet 
type TaobaoBusSeatpriceGetResultSet struct {

    // 错误代码
    
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    

    // 错误描述
    
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    

    // 余票对象
    
    Module   *B2BBusSeatPriceDto `json:"module,omitempty" xml:"module,omitempty"`
    

    // serverIP
    
    ServerIP   string `json:"server_i_p,omitempty" xml:"server_i_p,omitempty"`
    

    // 是否查询成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
