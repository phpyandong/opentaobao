package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建满就送活动 APIResponse
taobao.promotionmisc.mjs.activity.add

创建满就送活动。<br/>1、可以选择是全店参加或者部分商品参加：participate_range：0表示全部参与； 1表示部分商品参与。 2、如果是部分商品参加，则需要通过taobao.promotionmisc.activity.range.add接口来指定需要参加的商品。 3、该接口创建的优惠受店铺最低折扣限制，如优惠不生效，请让卖家检查该优惠是否低于店铺的最低折扣设置。
*/
type TaobaoPromotionmiscMjsActivityAddAPIResponse struct {
    model.CommonResponse
    TaobaoPromotionmiscMjsActivityAddResponse
}

type TaobaoPromotionmiscMjsActivityAddResponse struct {
    XMLName xml.Name `xml:"promotionmisc_mjs_activity_add_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否保存成功。
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 活动id。
    
    ActivityId   int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`

    
}
