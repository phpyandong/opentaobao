package tbk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-推广者-店铺搜索 APIRequest
taobao.tbk.shop.get

淘宝客店铺查询
*/
type TaobaoTbkShopGetRequest struct {
    model.Params

    // 需返回的字段列表
    fields   string 

    // 查询词
    q   string 

    // 排序_des（降序），排序_asc（升序），佣金比率（commission_rate）， 商品数量（auction_count），销售总数量（total_auction）
    sort   string 

    // 是否商城的店铺，设置为true表示该是属于淘宝商城的店铺，设置为false或不设置表示不判断这个属性
    isTmall   bool 

    // 信用等级下限，1~20
    startCredit   int64 

    // 信用等级上限，1~20
    endCredit   int64 

    // 淘客佣金比率下限，1~10000
    startCommissionRate   int64 

    // 淘客佣金比率上限，1~10000
    endCommissionRate   int64 

    // 店铺商品总数下限
    startTotalAction   int64 

    // 店铺商品总数上限
    endTotalAction   int64 

    // 累计推广商品下限
    startAuctionCount   int64 

    // 累计推广商品上限
    endAuctionCount   int64 

    // 链接形式：1：PC，2：无线，默认：１
    platform   int64 

    // 第几页，默认1，1~100
    pageNo   int64 

    // 页大小，默认20，1~100
    pageSize   int64 

}

func NewTaobaoTbkShopGetRequest() *TaobaoTbkShopGetRequest{
    return &TaobaoTbkShopGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTbkShopGetRequest) GetApiMethodName() string {
    return "taobao.tbk.shop.get"
}

func (r TaobaoTbkShopGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTbkShopGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoTbkShopGetRequest) GetFields() string {
    return r.fields
}

func (r *TaobaoTbkShopGetRequest) SetQ(q string) error {
    r.q = q
    r.Set("q", q)
    return nil
}

func (r TaobaoTbkShopGetRequest) GetQ() string {
    return r.q
}

func (r *TaobaoTbkShopGetRequest) SetSort(sort string) error {
    r.sort = sort
    r.Set("sort", sort)
    return nil
}

func (r TaobaoTbkShopGetRequest) GetSort() string {
    return r.sort
}

func (r *TaobaoTbkShopGetRequest) SetIsTmall(isTmall bool) error {
    r.isTmall = isTmall
    r.Set("is_tmall", isTmall)
    return nil
}

func (r TaobaoTbkShopGetRequest) GetIsTmall() bool {
    return r.isTmall
}

func (r *TaobaoTbkShopGetRequest) SetStartCredit(startCredit int64) error {
    r.startCredit = startCredit
    r.Set("start_credit", startCredit)
    return nil
}

func (r TaobaoTbkShopGetRequest) GetStartCredit() int64 {
    return r.startCredit
}

func (r *TaobaoTbkShopGetRequest) SetEndCredit(endCredit int64) error {
    r.endCredit = endCredit
    r.Set("end_credit", endCredit)
    return nil
}

func (r TaobaoTbkShopGetRequest) GetEndCredit() int64 {
    return r.endCredit
}

func (r *TaobaoTbkShopGetRequest) SetStartCommissionRate(startCommissionRate int64) error {
    r.startCommissionRate = startCommissionRate
    r.Set("start_commission_rate", startCommissionRate)
    return nil
}

func (r TaobaoTbkShopGetRequest) GetStartCommissionRate() int64 {
    return r.startCommissionRate
}

func (r *TaobaoTbkShopGetRequest) SetEndCommissionRate(endCommissionRate int64) error {
    r.endCommissionRate = endCommissionRate
    r.Set("end_commission_rate", endCommissionRate)
    return nil
}

func (r TaobaoTbkShopGetRequest) GetEndCommissionRate() int64 {
    return r.endCommissionRate
}

func (r *TaobaoTbkShopGetRequest) SetStartTotalAction(startTotalAction int64) error {
    r.startTotalAction = startTotalAction
    r.Set("start_total_action", startTotalAction)
    return nil
}

func (r TaobaoTbkShopGetRequest) GetStartTotalAction() int64 {
    return r.startTotalAction
}

func (r *TaobaoTbkShopGetRequest) SetEndTotalAction(endTotalAction int64) error {
    r.endTotalAction = endTotalAction
    r.Set("end_total_action", endTotalAction)
    return nil
}

func (r TaobaoTbkShopGetRequest) GetEndTotalAction() int64 {
    return r.endTotalAction
}

func (r *TaobaoTbkShopGetRequest) SetStartAuctionCount(startAuctionCount int64) error {
    r.startAuctionCount = startAuctionCount
    r.Set("start_auction_count", startAuctionCount)
    return nil
}

func (r TaobaoTbkShopGetRequest) GetStartAuctionCount() int64 {
    return r.startAuctionCount
}

func (r *TaobaoTbkShopGetRequest) SetEndAuctionCount(endAuctionCount int64) error {
    r.endAuctionCount = endAuctionCount
    r.Set("end_auction_count", endAuctionCount)
    return nil
}

func (r TaobaoTbkShopGetRequest) GetEndAuctionCount() int64 {
    return r.endAuctionCount
}

func (r *TaobaoTbkShopGetRequest) SetPlatform(platform int64) error {
    r.platform = platform
    r.Set("platform", platform)
    return nil
}

func (r TaobaoTbkShopGetRequest) GetPlatform() int64 {
    return r.platform
}

func (r *TaobaoTbkShopGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoTbkShopGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoTbkShopGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoTbkShopGetRequest) GetPageSize() int64 {
    return r.pageSize
}

