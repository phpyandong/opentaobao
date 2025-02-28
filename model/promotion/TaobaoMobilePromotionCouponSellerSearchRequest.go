package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询绑定卖家优惠券相关信息(手淘专用) APIRequest
taobao.mobile.promotion.coupon.seller.search

查询绑定卖家相关优惠券信息 如isv 百川 等外部业务方
*/
type TaobaoMobilePromotionCouponSellerSearchRequest struct {
    model.Params

    // 请求id 排查线索 需保证单次调用唯一
    traceId   string 

    // 券id集合
    spreadIds   string 

    // 每页数据 最大20左右
    pageSize   int64 

    // 当前第几页 从第一页开始
    currentPage   int64 

}

func NewTaobaoMobilePromotionCouponSellerSearchRequest() *TaobaoMobilePromotionCouponSellerSearchRequest{
    return &TaobaoMobilePromotionCouponSellerSearchRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMobilePromotionCouponSellerSearchRequest) GetApiMethodName() string {
    return "taobao.mobile.promotion.coupon.seller.search"
}

func (r TaobaoMobilePromotionCouponSellerSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMobilePromotionCouponSellerSearchRequest) SetTraceId(traceId string) error {
    r.traceId = traceId
    r.Set("trace_id", traceId)
    return nil
}

func (r TaobaoMobilePromotionCouponSellerSearchRequest) GetTraceId() string {
    return r.traceId
}

func (r *TaobaoMobilePromotionCouponSellerSearchRequest) SetSpreadIds(spreadIds string) error {
    r.spreadIds = spreadIds
    r.Set("spread_ids", spreadIds)
    return nil
}

func (r TaobaoMobilePromotionCouponSellerSearchRequest) GetSpreadIds() string {
    return r.spreadIds
}

func (r *TaobaoMobilePromotionCouponSellerSearchRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoMobilePromotionCouponSellerSearchRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoMobilePromotionCouponSellerSearchRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

func (r TaobaoMobilePromotionCouponSellerSearchRequest) GetCurrentPage() int64 {
    return r.currentPage
}

