package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取营销积木块 APIResponse
taobao.ump.mbb.getbyid

根据积木块id获取营销积木块。
*/
type TaobaoUmpMbbGetbyidAPIResponse struct {
    model.CommonResponse
    TaobaoUmpMbbGetbyidResponse
}

type TaobaoUmpMbbGetbyidResponse struct {
    XMLName xml.Name `xml:"ump_mbb_getbyid_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 营销积木块定义信息，可以通过ump sdk里面的MBB.fromJson来处理
    
    Mbb   string `json:"mbb,omitempty" xml:"mbb,omitempty"`

    
}
