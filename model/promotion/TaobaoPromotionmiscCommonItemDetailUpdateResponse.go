package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改通用单品优惠详情 APIResponse
taobao.promotionmisc.common.item.detail.update

修改通用单品优惠详情。
1、该接口只修改活动下参与的商品的优惠信息，如需要增加、删除活动，请调用taobao.promotionmisc.common.item.activity.add和taobao.promotionmisc.common.item.activity.delete接口；
2、使用该接口时需要把未做修改的字段值也传入；
3、此接口受卖家最低折扣限制，如果优惠力度大于卖家设置的最低折扣则不能修改
*/
type TaobaoPromotionmiscCommonItemDetailUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoPromotionmiscCommonItemDetailUpdateResponse
}

type TaobaoPromotionmiscCommonItemDetailUpdateResponse struct {
    XMLName xml.Name `xml:"promotionmisc_common_item_detail_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否修改成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
