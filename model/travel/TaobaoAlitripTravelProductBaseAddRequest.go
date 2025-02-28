package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商新增产品API APIRequest
taobao.alitrip.travel.product.base.add

飞猪供销平台供应商可通过该API发布新产品
*/
type TaobaoAlitripTravelProductBaseAddRequest struct {
    model.Params

    // 产品基本信息
    baseInfo   *ProductBaseInfo 

    // 选填，详细行程描述结构
    itineraries   []ItemItineraryInfo 

    // 选填，退款规则结构
    refundInfo   *ItemRefundInfo 

    // 必填，预定规则结构。示例： [{  "rule_type": "fee_excluded",  "rule_desc": "费用包含描述"},{  "rule_type": "fee_included",  "rule_desc": "费用不含描述"},{  "rule_type": "order_info",  "rule_desc": "预定须知描述"}]
    bookingRules   []BookingRuleInfo 

    // 特殊选填（当发布邮轮商品时必填，其他情况不填）邮轮商品相关信息，发布邮轮商品时必填
    cruiseProductExt   *CruiseProductExt 

    // 选填，商品的销售属性相关信息
    productSaleInfo   *ProductSaleInfo 

}

func NewTaobaoAlitripTravelProductBaseAddRequest() *TaobaoAlitripTravelProductBaseAddRequest{
    return &TaobaoAlitripTravelProductBaseAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripTravelProductBaseAddRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.product.base.add"
}

func (r TaobaoAlitripTravelProductBaseAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripTravelProductBaseAddRequest) SetBaseInfo(baseInfo *ProductBaseInfo) error {
    r.baseInfo = baseInfo
    r.Set("base_info", baseInfo)
    return nil
}

func (r TaobaoAlitripTravelProductBaseAddRequest) GetBaseInfo() *ProductBaseInfo {
    return r.baseInfo
}

func (r *TaobaoAlitripTravelProductBaseAddRequest) SetItineraries(itineraries []ItemItineraryInfo) error {
    r.itineraries = itineraries
    r.Set("itineraries", itineraries)
    return nil
}

func (r TaobaoAlitripTravelProductBaseAddRequest) GetItineraries() []ItemItineraryInfo {
    return r.itineraries
}

func (r *TaobaoAlitripTravelProductBaseAddRequest) SetRefundInfo(refundInfo *ItemRefundInfo) error {
    r.refundInfo = refundInfo
    r.Set("refund_info", refundInfo)
    return nil
}

func (r TaobaoAlitripTravelProductBaseAddRequest) GetRefundInfo() *ItemRefundInfo {
    return r.refundInfo
}

func (r *TaobaoAlitripTravelProductBaseAddRequest) SetBookingRules(bookingRules []BookingRuleInfo) error {
    r.bookingRules = bookingRules
    r.Set("booking_rules", bookingRules)
    return nil
}

func (r TaobaoAlitripTravelProductBaseAddRequest) GetBookingRules() []BookingRuleInfo {
    return r.bookingRules
}

func (r *TaobaoAlitripTravelProductBaseAddRequest) SetCruiseProductExt(cruiseProductExt *CruiseProductExt) error {
    r.cruiseProductExt = cruiseProductExt
    r.Set("cruise_product_ext", cruiseProductExt)
    return nil
}

func (r TaobaoAlitripTravelProductBaseAddRequest) GetCruiseProductExt() *CruiseProductExt {
    return r.cruiseProductExt
}

func (r *TaobaoAlitripTravelProductBaseAddRequest) SetProductSaleInfo(productSaleInfo *ProductSaleInfo) error {
    r.productSaleInfo = productSaleInfo
    r.Set("product_sale_info", productSaleInfo)
    return nil
}

func (r TaobaoAlitripTravelProductBaseAddRequest) GetProductSaleInfo() *ProductSaleInfo {
    return r.productSaleInfo
}

