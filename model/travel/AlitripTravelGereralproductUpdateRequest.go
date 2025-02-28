package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通用类目产品发布编辑 APIRequest
alitrip.travel.gereralproduct.update

提供给飞猪供销平台供应商发布编辑通用类目产品的API
*/
type AlitripTravelGereralproductUpdateRequest struct {
    model.Params

    // 产品基本信息
    baseInfo   *GeneralProductBaseInfo 

    // 退款规则结构
    refundInfo   *ItemRefundInfo 

    // 必填，预定规则结构。示例： [{ "rule_type": "fee_excluded", "rule_desc": "费用包含描述"},{ "rule_type": "fee_included", "rule_desc": "费用不含描述"},{ "rule_type": "order_info", "rule_desc": "预定须知描述"}]
    bookingRules   []BookingRuleInfo 

    // 产品销售信息
    productSaleInfo   *ProductSaleInfo 

    // 更新sku信息，仅限日历商品使用
    dateSkuInfoList   []DateSkuInfo 

}

func NewAlitripTravelGereralproductUpdateRequest() *AlitripTravelGereralproductUpdateRequest{
    return &AlitripTravelGereralproductUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripTravelGereralproductUpdateRequest) GetApiMethodName() string {
    return "alitrip.travel.gereralproduct.update"
}

func (r AlitripTravelGereralproductUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripTravelGereralproductUpdateRequest) SetBaseInfo(baseInfo *GeneralProductBaseInfo) error {
    r.baseInfo = baseInfo
    r.Set("base_info", baseInfo)
    return nil
}

func (r AlitripTravelGereralproductUpdateRequest) GetBaseInfo() *GeneralProductBaseInfo {
    return r.baseInfo
}

func (r *AlitripTravelGereralproductUpdateRequest) SetRefundInfo(refundInfo *ItemRefundInfo) error {
    r.refundInfo = refundInfo
    r.Set("refund_info", refundInfo)
    return nil
}

func (r AlitripTravelGereralproductUpdateRequest) GetRefundInfo() *ItemRefundInfo {
    return r.refundInfo
}

func (r *AlitripTravelGereralproductUpdateRequest) SetBookingRules(bookingRules []BookingRuleInfo) error {
    r.bookingRules = bookingRules
    r.Set("booking_rules", bookingRules)
    return nil
}

func (r AlitripTravelGereralproductUpdateRequest) GetBookingRules() []BookingRuleInfo {
    return r.bookingRules
}

func (r *AlitripTravelGereralproductUpdateRequest) SetProductSaleInfo(productSaleInfo *ProductSaleInfo) error {
    r.productSaleInfo = productSaleInfo
    r.Set("product_sale_info", productSaleInfo)
    return nil
}

func (r AlitripTravelGereralproductUpdateRequest) GetProductSaleInfo() *ProductSaleInfo {
    return r.productSaleInfo
}

func (r *AlitripTravelGereralproductUpdateRequest) SetDateSkuInfoList(dateSkuInfoList []DateSkuInfo) error {
    r.dateSkuInfoList = dateSkuInfoList
    r.Set("date_sku_info_list", dateSkuInfoList)
    return nil
}

func (r AlitripTravelGereralproductUpdateRequest) GetDateSkuInfoList() []DateSkuInfo {
    return r.dateSkuInfoList
}

