package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据商品ID查询所有库存变更记录 APIRequest
taobao.wlb.inventorylog.query

通过商品ID等几个条件来分页查询库存变更记录
*/
type TaobaoWlbInventorylogQueryRequest struct {
    model.Params

    // 商品ID
    itemId   int64 

    // 仓库编码
    storeCode   string 

    // 单号
    orderCode   string 

    // 起始修改时间,大于等于该时间
    gmtStart   string 

    // 结束修改时间,小于等于该时间
    gmtEnd   string 

    // 当前页
    pageNo   int64 

    // 分页记录个数
    pageSize   int64 

    // 可指定授权的用户来查询
    opUserId   int64 

    // 库存操作作类型(可以为空) CHU_KU 1-出库 RU_KU 2-入库 FREEZE 3-冻结 THAW 4-解冻 CHECK_FREEZE 5-冻结确认 CHANGE_KU 6-库存类型变更 若值不在范围内，则按CHU_KU处理
    opType   string 

}

func NewTaobaoWlbInventorylogQueryRequest() *TaobaoWlbInventorylogQueryRequest{
    return &TaobaoWlbInventorylogQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbInventorylogQueryRequest) GetApiMethodName() string {
    return "taobao.wlb.inventorylog.query"
}

func (r TaobaoWlbInventorylogQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbInventorylogQueryRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoWlbInventorylogQueryRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TaobaoWlbInventorylogQueryRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

func (r TaobaoWlbInventorylogQueryRequest) GetStoreCode() string {
    return r.storeCode
}

func (r *TaobaoWlbInventorylogQueryRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

func (r TaobaoWlbInventorylogQueryRequest) GetOrderCode() string {
    return r.orderCode
}

func (r *TaobaoWlbInventorylogQueryRequest) SetGmtStart(gmtStart string) error {
    r.gmtStart = gmtStart
    r.Set("gmt_start", gmtStart)
    return nil
}

func (r TaobaoWlbInventorylogQueryRequest) GetGmtStart() string {
    return r.gmtStart
}

func (r *TaobaoWlbInventorylogQueryRequest) SetGmtEnd(gmtEnd string) error {
    r.gmtEnd = gmtEnd
    r.Set("gmt_end", gmtEnd)
    return nil
}

func (r TaobaoWlbInventorylogQueryRequest) GetGmtEnd() string {
    return r.gmtEnd
}

func (r *TaobaoWlbInventorylogQueryRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoWlbInventorylogQueryRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoWlbInventorylogQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoWlbInventorylogQueryRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoWlbInventorylogQueryRequest) SetOpUserId(opUserId int64) error {
    r.opUserId = opUserId
    r.Set("op_user_id", opUserId)
    return nil
}

func (r TaobaoWlbInventorylogQueryRequest) GetOpUserId() int64 {
    return r.opUserId
}

func (r *TaobaoWlbInventorylogQueryRequest) SetOpType(opType string) error {
    r.opType = opType
    r.Set("op_type", opType)
    return nil
}

func (r TaobaoWlbInventorylogQueryRequest) GetOpType() string {
    return r.opType
}

