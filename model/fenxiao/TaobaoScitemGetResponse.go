package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据id查询商品 APIResponse
taobao.scitem.get

根据id查询商品
*/
type TaobaoScitemGetAPIResponse struct {
    model.CommonResponse
    TaobaoScitemGetResponse
}

type TaobaoScitemGetResponse struct {
    XMLName xml.Name `xml:"scitem_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 后端商品
    
    ScItem   *ScItem `json:"sc_item,omitempty" xml:"sc_item,omitempty"`

    
}
