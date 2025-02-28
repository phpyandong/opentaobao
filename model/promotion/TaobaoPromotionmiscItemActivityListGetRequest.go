package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询无条件单品优惠活动列表 APIRequest
taobao.promotionmisc.item.activity.list.get

查询无条件单品优惠活动列表
*/
type TaobaoPromotionmiscItemActivityListGetRequest struct {
    model.Params

    // 页码。
    pageNo   int64 

    // 每页记录数，最大支持50 。
    pageSize   int64 

}

func NewTaobaoPromotionmiscItemActivityListGetRequest() *TaobaoPromotionmiscItemActivityListGetRequest{
    return &TaobaoPromotionmiscItemActivityListGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPromotionmiscItemActivityListGetRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.item.activity.list.get"
}

func (r TaobaoPromotionmiscItemActivityListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPromotionmiscItemActivityListGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoPromotionmiscItemActivityListGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoPromotionmiscItemActivityListGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoPromotionmiscItemActivityListGetRequest) GetPageSize() int64 {
    return r.pageSize
}

