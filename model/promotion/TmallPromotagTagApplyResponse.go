package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠标签申请 APIResponse
tmall.promotag.tag.apply

创建优惠标签
*/
type TmallPromotagTagApplyAPIResponse struct {
    model.CommonResponse
    TmallPromotagTagApplyResponse
}

type TmallPromotagTagApplyResponse struct {
    XMLName xml.Name `xml:"tmall_promotag_tag_apply_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否设置成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 优惠标签ID
    
    TagId   int64 `json:"tag_id,omitempty" xml:"tag_id,omitempty"`

    
}
