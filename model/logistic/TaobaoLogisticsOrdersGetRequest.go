package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询物流订单 APIRequest
taobao.logistics.orders.get

批量查询物流订单。
*/
type TaobaoLogisticsOrdersGetRequest struct {
    model.Params

    // 需返回的字段列表.可选值:Shipping 物流数据结构中的以下字段: <br><br/>tid,order_code,seller_nick,buyer_nick,delivery_start, delivery_end,out_sid,item_title,receiver_name, created,modified,status,type,freight_payer,seller_confirm,company_name,sub_tids,is_spilt；<br>多个字段之间用","分隔。如tid,seller_nick,buyer_nick,delivery_start。
    fields   string 

    // 交易ID.如果加入tid参数的话,不用传其他的参数,若传入tid：非拆单场景，仅会返回一条物流订单信息；拆单场景，会返回多条物流订单信息
    tid   int64 

    // 买家昵称
    buyerNick   string 

    // 物流状态.查看数据结构 Shipping 中的status字段.
    status   string 

    // 卖家是否发货.可选值:yes(是),no(否).如:yes
    sellerConfirm   string 

    // 收货人姓名
    receiverName   string 

    // 创建时间开始
    startCreated   string 

    // 创建时间结束
    endCreated   string 

    // 谁承担运费.可选值:buyer(买家),seller(卖家).如:buyer
    freightPayer   string 

    // 物流方式.可选值:post(平邮),express(快递),ems(EMS).如:post
    type   string 

    // 页码.该字段没传 或 值<1 ,则默认page_no为1
    pageNo   int64 

    // 每页条数.该字段没传 或 值<1 ,则默认page_size为40
    pageSize   int64 

}

func NewTaobaoLogisticsOrdersGetRequest() *TaobaoLogisticsOrdersGetRequest{
    return &TaobaoLogisticsOrdersGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoLogisticsOrdersGetRequest) GetApiMethodName() string {
    return "taobao.logistics.orders.get"
}

func (r TaobaoLogisticsOrdersGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoLogisticsOrdersGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoLogisticsOrdersGetRequest) GetFields() string {
    return r.fields
}

func (r *TaobaoLogisticsOrdersGetRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoLogisticsOrdersGetRequest) GetTid() int64 {
    return r.tid
}

func (r *TaobaoLogisticsOrdersGetRequest) SetBuyerNick(buyerNick string) error {
    r.buyerNick = buyerNick
    r.Set("buyer_nick", buyerNick)
    return nil
}

func (r TaobaoLogisticsOrdersGetRequest) GetBuyerNick() string {
    return r.buyerNick
}

func (r *TaobaoLogisticsOrdersGetRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TaobaoLogisticsOrdersGetRequest) GetStatus() string {
    return r.status
}

func (r *TaobaoLogisticsOrdersGetRequest) SetSellerConfirm(sellerConfirm string) error {
    r.sellerConfirm = sellerConfirm
    r.Set("seller_confirm", sellerConfirm)
    return nil
}

func (r TaobaoLogisticsOrdersGetRequest) GetSellerConfirm() string {
    return r.sellerConfirm
}

func (r *TaobaoLogisticsOrdersGetRequest) SetReceiverName(receiverName string) error {
    r.receiverName = receiverName
    r.Set("receiver_name", receiverName)
    return nil
}

func (r TaobaoLogisticsOrdersGetRequest) GetReceiverName() string {
    return r.receiverName
}

func (r *TaobaoLogisticsOrdersGetRequest) SetStartCreated(startCreated string) error {
    r.startCreated = startCreated
    r.Set("start_created", startCreated)
    return nil
}

func (r TaobaoLogisticsOrdersGetRequest) GetStartCreated() string {
    return r.startCreated
}

func (r *TaobaoLogisticsOrdersGetRequest) SetEndCreated(endCreated string) error {
    r.endCreated = endCreated
    r.Set("end_created", endCreated)
    return nil
}

func (r TaobaoLogisticsOrdersGetRequest) GetEndCreated() string {
    return r.endCreated
}

func (r *TaobaoLogisticsOrdersGetRequest) SetFreightPayer(freightPayer string) error {
    r.freightPayer = freightPayer
    r.Set("freight_payer", freightPayer)
    return nil
}

func (r TaobaoLogisticsOrdersGetRequest) GetFreightPayer() string {
    return r.freightPayer
}

func (r *TaobaoLogisticsOrdersGetRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TaobaoLogisticsOrdersGetRequest) GetType() string {
    return r.type
}

func (r *TaobaoLogisticsOrdersGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoLogisticsOrdersGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoLogisticsOrdersGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoLogisticsOrdersGetRequest) GetPageSize() int64 {
    return r.pageSize
}

