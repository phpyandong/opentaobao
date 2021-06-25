package util

// ItemPickPagingResult 
type ItemPickPagingResult struct {

    // 返回数据集合
    Results   []CountryDto `json:"results,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty"`

    // 页码
    PageNo   int64 `json:"page_no,omitempty"`

    // 每页数据条数
    PageSize   int64 `json:"page_size,omitempty"`

    // 总记录条数
    TotalCount   int64 `json:"total_count,omitempty"`

}
