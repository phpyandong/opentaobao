package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）批量获取创意 APIResponse
taobao.simba.salestar.creatives.get

取得一个推广组的所有创意或者根据一个创意Id列表取得一组创意；<br/>如果同时提供了推广组Id和创意id列表，则优先使用推广组Id；
*/
type TaobaoSimbaSalestarCreativesGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaSalestarCreativesGetResponse
}

type TaobaoSimbaSalestarCreativesGetResponse struct {
    XMLName xml.Name `xml:"simba_salestar_creatives_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 创意对象列表
    
    Creatives   []Creative `json:"creatives,omitempty" xml:"creatives>creative,omitempty"`
    
    
}
