package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询活动参与的商品 APIResponse
taobao.promotionmisc.activity.range.list.get

查询活动参与的商品
*/
type TaobaoPromotionmiscActivityRangeListGetAPIResponse struct {
    model.CommonResponse
    TaobaoPromotionmiscActivityRangeListGetResponse
}

type TaobaoPromotionmiscActivityRangeListGetResponse struct {
    XMLName xml.Name `xml:"promotionmisc_activity_range_list_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 活动参与的商品列表
    
    PromotionRangeList   []PromotionRange `json:"promotion_range_list,omitempty" xml:"promotion_range_list>promotion_range,omitempty"`
    
    
}
