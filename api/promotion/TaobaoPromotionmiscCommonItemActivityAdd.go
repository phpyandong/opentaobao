package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
创建通用单品优惠活动 
taobao.promotionmisc.common.item.activity.add

创建通用单品优惠活动。
1、该接口只创建活动的基本信息，如需要增加、删除参与该活动的商品及优惠，请调用taobao.promotionmisc.common.item.detail.add和taobao.promotionmisc.common.item.detail.delete接口
2、同一卖家下的活动数量限制为30个，超过限制需先调用taobao.promotionmisc.common.item.activity.delete接口删除无用的活动后才可再创建新的活动
*/
func TaobaoPromotionmiscCommonItemActivityAdd(clt *core.SDKClient, req *promotion.TaobaoPromotionmiscCommonItemActivityAddRequest, session string) (*promotion.TaobaoPromotionmiscCommonItemActivityAddAPIResponse, error) {
    var resp promotion.TaobaoPromotionmiscCommonItemActivityAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
