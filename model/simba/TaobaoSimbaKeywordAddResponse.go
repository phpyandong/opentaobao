package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）关键词新增接口 APIResponse
taobao.simba.keyword.add

（新）关键词更新相关接口
*/
type TaobaoSimbaKeywordAddAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaKeywordAddResponse
}

type TaobaoSimbaKeywordAddResponse struct {
    XMLName xml.Name `xml:"simba_keyword_add_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 整体的返回值
    
    Results   []SiriusBidwordDto `json:"results,omitempty" xml:"results>sirius_bidword_dto,omitempty"`
    
    
    // 错误原因
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`

    
}
