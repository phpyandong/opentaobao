package tbk

// TaobaoTbkDgPunishOrderGetRpcResult 
type TaobaoTbkDgPunishOrderGetRpcResult struct {

    // 结果
    
    Data   *PageResult `json:"data,omitempty" xml:"data,omitempty"`
    

    // 业务出错的描述
    
    BizErrorDesc   string `json:"biz_error_desc,omitempty" xml:"biz_error_desc,omitempty"`
    

    // 业务出错的状态码
    
    BizErrorCode   int64 `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
    

    // 执行结果
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
    

    // 执行结果状态码
    
    ResultCode   int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
    

}
