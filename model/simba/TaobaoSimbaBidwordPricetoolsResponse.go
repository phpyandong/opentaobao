package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
关键词出价指导工具（新） APIResponse
taobao.simba.bidword.pricetools

关键词出价指导工具（新）
*/
type TaobaoSimbaBidwordPricetoolsAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaBidwordPricetoolsResponse
}

type TaobaoSimbaBidwordPricetoolsResponse struct {
    XMLName xml.Name `xml:"simba_bidword_pricetools_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // true 表示符合准入，false不符合
    
    ResultList   *PriceSuggestionDto `json:"result_list,omitempty" xml:"result_list,omitempty"`

    
}
