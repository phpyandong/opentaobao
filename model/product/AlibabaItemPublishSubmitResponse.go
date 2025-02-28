package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品发布 APIResponse
alibaba.item.publish.submit

新商品发布，提交商品发布信息
*/
type AlibabaItemPublishSubmitAPIResponse struct {
    model.CommonResponse
    AlibabaItemPublishSubmitResponse
}

type AlibabaItemPublishSubmitResponse struct {
    XMLName xml.Name `xml:"alibaba_item_publish_submit_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品创建时间
    
    CreateTime   string `json:"create_time,omitempty" xml:"create_time,omitempty"`

    
    // 商品ID
    
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`

    
    // 商品所属市场
    
    Market   string `json:"market,omitempty" xml:"market,omitempty"`

    
}
