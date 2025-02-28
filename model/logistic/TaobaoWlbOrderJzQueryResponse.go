package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
家装业务查询物流公司api APIResponse
taobao.wlb.order.jz.query

家装业务查询物流公司api
*/
type TaobaoWlbOrderJzQueryAPIResponse struct {
    model.CommonResponse
    TaobaoWlbOrderJzQueryResponse
}

type TaobaoWlbOrderJzQueryResponse struct {
    XMLName xml.Name `xml:"wlb_order_jz_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误编码
    
    ResultErrorCode   string `json:"result_error_code,omitempty" xml:"result_error_code,omitempty"`

    
    // 错误信息
    
    ResultErrorMsg   string `json:"result_error_msg,omitempty" xml:"result_error_msg,omitempty"`

    
    // 结果信息
    
    Result   *JzTopDto `json:"result,omitempty" xml:"result,omitempty"`

    
    // 是否成功
    
    ResultSuccess   bool `json:"result_success,omitempty" xml:"result_success,omitempty"`

    
}
