package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询无条件单品优惠活动 APIResponse
taobao.promotionmisc.item.activity.get

查询无条件单品优惠活动
*/
type TaobaoPromotionmiscItemActivityGetAPIResponse struct {
    model.CommonResponse
    TaobaoPromotionmiscItemActivityGetResponse
}

type TaobaoPromotionmiscItemActivityGetResponse struct {
    XMLName xml.Name `xml:"promotionmisc_item_activity_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 单品优惠活动信息。
    
    ItemPromotion   *ItemPromotion `json:"item_promotion,omitempty" xml:"item_promotion,omitempty"`

    
}
