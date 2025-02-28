package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
除度假线路、门票以外的其他类目商品维护接口（商品ID重复将自动更新） APIRequest
alitrip.travel.gereralitem.update

除度假线路、门票以外的商品维护接口；目前该接口支持以下类目；
（123740001：客栈周边交通服务、125038002：旅行设备/GPS/相机租赁、50018298：船票、124084006：酒店商品升级差价、125228016：预约卡券、50011954：旅游服务、50012913：酒店优惠券、50214003：旅游会员卡/酒店会员卡、50012917：巴士/地铁/交通卡/一卡通、50134002：代客烧香/还愿、50026091：境外火车票、123742001：客栈周边餐饮服务、50019817：海外套餐（该类目已废弃）、125210016：团建/outing、124212017：其他预定、50025888：机场行李托运取送寄存、50025831：旅游年票/年卡、124142009：旅游商品升级差价/押金、123744001：客栈周边其他服务、50012762：广深口岸港澳送关服务、50025880：旅行拍照/婚纱摄影、123166001：酒店餐饮美食（该类目已废弃）、50668002：手绘地图/明信片、50024210：旅游购物/纪念品、50024208：酒店用品、50024215：购物折扣卡券、50025878：酒店SPA/足浴/温泉、50024212：旅游必备品、123738001：客栈周边票务服务、123164002：游泳健身（该类目已废弃）、50686003：机票增值产品、123164001：酒店SPA（该类目已废弃）、124832008：美食卡券/酒店餐饮卡券、125408001：旅游定制服务、50018112：旅行社/网站优惠券、124258004：酒店客房优惠券（该类目已废弃）、50104001：机场周边停车位、124730002：内机机场/火车站送关服务、124090010：其他）
*/
type AlitripTravelGereralitemUpdateRequest struct {
    model.Params

    // 必填，商品基本信息
    baseInfo   *BaseInfo 

    // 必填，预定规则结构。示例： [{ "rule_type": "fee_excluded", "rule_desc": "费用包含描述"},{ "rule_type": "fee_included", "rule_desc": "费用不含描述"},{ "rule_type": "order_info", "rule_desc": "预定须知描述"}]
    bookingRules   []BookingRuleInfo 

    // 更新sku信息，仅限日历商品使用
    dateSkuInfoList   []DateSkuInfo 

    // 新版电子凭证信息。如果传递了此参数，则sales_info中旧版电子凭证信息将被忽略，否则电子凭证信息将以旧版电子凭证参数为准。如果新、旧版参数都没传，则该商品不支持电子凭证
    itemEleCertInfo   *ItemEleCertInfo 

    // 选填，退款规则结构
    itemRefundInfo   *ItemRefundInfo 

    // poi信息，个别类目必填，如演艺类目下场馆信息
    poi   *Poi 

    // 更新sku信息，仅限非日历（普通）商品使用
    commonSkuList   []NoDateSkuInfo 

}

func NewAlitripTravelGereralitemUpdateRequest() *AlitripTravelGereralitemUpdateRequest{
    return &AlitripTravelGereralitemUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripTravelGereralitemUpdateRequest) GetApiMethodName() string {
    return "alitrip.travel.gereralitem.update"
}

func (r AlitripTravelGereralitemUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripTravelGereralitemUpdateRequest) SetBaseInfo(baseInfo *BaseInfo) error {
    r.baseInfo = baseInfo
    r.Set("base_info", baseInfo)
    return nil
}

func (r AlitripTravelGereralitemUpdateRequest) GetBaseInfo() *BaseInfo {
    return r.baseInfo
}

func (r *AlitripTravelGereralitemUpdateRequest) SetBookingRules(bookingRules []BookingRuleInfo) error {
    r.bookingRules = bookingRules
    r.Set("booking_rules", bookingRules)
    return nil
}

func (r AlitripTravelGereralitemUpdateRequest) GetBookingRules() []BookingRuleInfo {
    return r.bookingRules
}

func (r *AlitripTravelGereralitemUpdateRequest) SetDateSkuInfoList(dateSkuInfoList []DateSkuInfo) error {
    r.dateSkuInfoList = dateSkuInfoList
    r.Set("date_sku_info_list", dateSkuInfoList)
    return nil
}

func (r AlitripTravelGereralitemUpdateRequest) GetDateSkuInfoList() []DateSkuInfo {
    return r.dateSkuInfoList
}

func (r *AlitripTravelGereralitemUpdateRequest) SetItemEleCertInfo(itemEleCertInfo *ItemEleCertInfo) error {
    r.itemEleCertInfo = itemEleCertInfo
    r.Set("item_ele_cert_info", itemEleCertInfo)
    return nil
}

func (r AlitripTravelGereralitemUpdateRequest) GetItemEleCertInfo() *ItemEleCertInfo {
    return r.itemEleCertInfo
}

func (r *AlitripTravelGereralitemUpdateRequest) SetItemRefundInfo(itemRefundInfo *ItemRefundInfo) error {
    r.itemRefundInfo = itemRefundInfo
    r.Set("item_refund_info", itemRefundInfo)
    return nil
}

func (r AlitripTravelGereralitemUpdateRequest) GetItemRefundInfo() *ItemRefundInfo {
    return r.itemRefundInfo
}

func (r *AlitripTravelGereralitemUpdateRequest) SetPoi(poi *Poi) error {
    r.poi = poi
    r.Set("poi", poi)
    return nil
}

func (r AlitripTravelGereralitemUpdateRequest) GetPoi() *Poi {
    return r.poi
}

func (r *AlitripTravelGereralitemUpdateRequest) SetCommonSkuList(commonSkuList []NoDateSkuInfo) error {
    r.commonSkuList = commonSkuList
    r.Set("common_sku_list", commonSkuList)
    return nil
}

func (r AlitripTravelGereralitemUpdateRequest) GetCommonSkuList() []NoDateSkuInfo {
    return r.commonSkuList
}

