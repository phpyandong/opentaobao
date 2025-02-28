package tanx

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
对外部dsp提供交易id查询接口 APIResponse
taobao.tanx.deal.get

对外部dsp提供交易id查询接口
*/
type TaobaoTanxDealGetAPIResponse struct {
    model.CommonResponse
    TaobaoTanxDealGetResponse
}

type TaobaoTanxDealGetResponse struct {
    XMLName xml.Name `xml:"tanx_deal_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果代码
    
    Code   int64 `json:"code,omitempty" xml:"code,omitempty"`

    
    // 结果信息
    
    Messag   string `json:"messag,omitempty" xml:"messag,omitempty"`

    
    // 查询结果
    
    Sucess   bool `json:"sucess,omitempty" xml:"sucess,omitempty"`

    
    // 查询结果
    
    Result   *DealInfoDTO `json:"result,omitempty" xml:"result,omitempty"`

    
}
