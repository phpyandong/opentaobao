package wms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发货订单通知 APIRequest
taobao.wlb.wms.consign.order.notify

发货订单通知
*/
type TaobaoWlbWmsConsignOrderNotifyRequest struct {
    model.Params

    // ERP订单号
    orderCode   string 

    // 单据类型 201 一般交易出库单 202 B2B交易出库单 502 换货出库单 503 补发出库单
    orderType   int64 

    // 订单标识 (1: cod –货到付款，4:invoiceinfo-需要发票)
    orderFlag   string 

    // 订单来源（213 天猫，201 淘宝，214 京东，202 1688 阿里中文站 ，203 苏宁在线，204 亚马逊中国，205 当当，208 1号店，207 唯品会，209 国美在线，210 拍拍，206 易贝ebay，211 聚美优品，212 乐蜂网，215 邮乐，216 凡客，217 优购，218 银泰，219 易讯，221 聚尚网，222 蘑菇街，223 POS门店，301 其他）
    orderSource   int64 

    // 仓库编码，此字段为空时，由菜鸟路由仓库发货
    storeCode   string 

    // 快递公司编码，此字段为空时，由菜鸟选择快递配送
    tmsServiceCode   string 

    // 快递公司名称
    tmsServiceName   string 

    // 前物流订单号，订单类型为502 换货出库单 503 补发出库单时，需求传入此内容
    prevOrderCode   string 

    // 下单时间，订单在交易平台创建时间
    orderShopCreateTime   string 

    // 订单支付时间
    orderPayTime   string 

    // 订单创建时间,ERP创建订单时间
    orderCreateTime   string 

    // 订单审核时间,ERP创建支付时间
    orderExaminationTime   string 

    // 订单总金额,=总商品金额-订单优惠金额+快递费用，单位分
    orderAmount   int64 

    // 订单优惠金额，整单优惠金额，单位分
    discountAmount   int64 

    // 订单应收金额，消费者还需要付的金额，单位分
    arAmount   int64 

    // 订单已付金额，消费者已经支付的金额，单位分
    gotAmount   int64 

    // 快递费用，单位分
    postfee   int64 

    // COD服务费，单位分
    serviceFee   int64 

    // 配送要求
    deliverRequirements   *Deliverrequirementswlbwmsconsignordernotify 

    // 收件人信息
    receiverInfo   *Receiverwlbwmsconsignordernotify 

    // 发货方信息
    senderInfo   *Senderwlbwmsconsignordernotify 

    // 订单商品信息列表
    orderItemList   []Orderitemlistwlbwmsconsignordernotify 

    // 发票信息列表
    invoiceInfoList   []Invoicelistwlbwmsconsignordernotify 

    // 拓展属性
    extendFields   string 

    // 备注
    remark   string 

}

func NewTaobaoWlbWmsConsignOrderNotifyRequest() *TaobaoWlbWmsConsignOrderNotifyRequest{
    return &TaobaoWlbWmsConsignOrderNotifyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbWmsConsignOrderNotifyRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.consign.order.notify"
}

func (r TaobaoWlbWmsConsignOrderNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbWmsConsignOrderNotifyRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

func (r TaobaoWlbWmsConsignOrderNotifyRequest) GetOrderCode() string {
    return r.orderCode
}

func (r *TaobaoWlbWmsConsignOrderNotifyRequest) SetOrderType(orderType int64) error {
    r.orderType = orderType
    r.Set("order_type", orderType)
    return nil
}

func (r TaobaoWlbWmsConsignOrderNotifyRequest) GetOrderType() int64 {
    return r.orderType
}

func (r *TaobaoWlbWmsConsignOrderNotifyRequest) SetOrderFlag(orderFlag string) error {
    r.orderFlag = orderFlag
    r.Set("order_flag", orderFlag)
    return nil
}

func (r TaobaoWlbWmsConsignOrderNotifyRequest) GetOrderFlag() string {
    return r.orderFlag
}

func (r *TaobaoWlbWmsConsignOrderNotifyRequest) SetOrderSource(orderSource int64) error {
    r.orderSource = orderSource
    r.Set("order_source", orderSource)
    return nil
}

func (r TaobaoWlbWmsConsignOrderNotifyRequest) GetOrderSource() int64 {
    return r.orderSource
}

func (r *TaobaoWlbWmsConsignOrderNotifyRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

func (r TaobaoWlbWmsConsignOrderNotifyRequest) GetStoreCode() string {
    return r.storeCode
}

func (r *TaobaoWlbWmsConsignOrderNotifyRequest) SetTmsServiceCode(tmsServiceCode string) error {
    r.tmsServiceCode = tmsServiceCode
    r.Set("tms_service_code", tmsServiceCode)
    return nil
}

func (r TaobaoWlbWmsConsignOrderNotifyRequest) GetTmsServiceCode() string {
    return r.tmsServiceCode
}

func (r *TaobaoWlbWmsConsignOrderNotifyRequest) SetTmsServiceName(tmsServiceName string) error {
    r.tmsServiceName = tmsServiceName
    r.Set("tms_service_name", tmsServiceName)
    return nil
}

func (r TaobaoWlbWmsConsignOrderNotifyRequest) GetTmsServiceName() string {
    return r.tmsServiceName
}

func (r *TaobaoWlbWmsConsignOrderNotifyRequest) SetPrevOrderCode(prevOrderCode string) error {
    r.prevOrderCode = prevOrderCode
    r.Set("prev_order_code", prevOrderCode)
    return nil
}

func (r TaobaoWlbWmsConsignOrderNotifyRequest) GetPrevOrderCode() string {
    return r.prevOrderCode
}

func (r *TaobaoWlbWmsConsignOrderNotifyRequest) SetOrderShopCreateTime(orderShopCreateTime string) error {
    r.orderShopCreateTime = orderShopCreateTime
    r.Set("order_shop_create_time", orderShopCreateTime)
    return nil
}

func (r TaobaoWlbWmsConsignOrderNotifyRequest) GetOrderShopCreateTime() string {
    return r.orderShopCreateTime
}

func (r *TaobaoWlbWmsConsignOrderNotifyRequest) SetOrderPayTime(orderPayTime string) error {
    r.orderPayTime = orderPayTime
    r.Set("order_pay_time", orderPayTime)
    return nil
}

func (r TaobaoWlbWmsConsignOrderNotifyRequest) GetOrderPayTime() string {
    return r.orderPayTime
}

func (r *TaobaoWlbWmsConsignOrderNotifyRequest) SetOrderCreateTime(orderCreateTime string) error {
    r.orderCreateTime = orderCreateTime
    r.Set("order_create_time", orderCreateTime)
    return nil
}

func (r TaobaoWlbWmsConsignOrderNotifyRequest) GetOrderCreateTime() string {
    return r.orderCreateTime
}

func (r *TaobaoWlbWmsConsignOrderNotifyRequest) SetOrderExaminationTime(orderExaminationTime string) error {
    r.orderExaminationTime = orderExaminationTime
    r.Set("order_examination_time", orderExaminationTime)
    return nil
}

func (r TaobaoWlbWmsConsignOrderNotifyRequest) GetOrderExaminationTime() string {
    return r.orderExaminationTime
}

func (r *TaobaoWlbWmsConsignOrderNotifyRequest) SetOrderAmount(orderAmount int64) error {
    r.orderAmount = orderAmount
    r.Set("order_amount", orderAmount)
    return nil
}

func (r TaobaoWlbWmsConsignOrderNotifyRequest) GetOrderAmount() int64 {
    return r.orderAmount
}

func (r *TaobaoWlbWmsConsignOrderNotifyRequest) SetDiscountAmount(discountAmount int64) error {
    r.discountAmount = discountAmount
    r.Set("discount_amount", discountAmount)
    return nil
}

func (r TaobaoWlbWmsConsignOrderNotifyRequest) GetDiscountAmount() int64 {
    return r.discountAmount
}

func (r *TaobaoWlbWmsConsignOrderNotifyRequest) SetArAmount(arAmount int64) error {
    r.arAmount = arAmount
    r.Set("ar_amount", arAmount)
    return nil
}

func (r TaobaoWlbWmsConsignOrderNotifyRequest) GetArAmount() int64 {
    return r.arAmount
}

func (r *TaobaoWlbWmsConsignOrderNotifyRequest) SetGotAmount(gotAmount int64) error {
    r.gotAmount = gotAmount
    r.Set("got_amount", gotAmount)
    return nil
}

func (r TaobaoWlbWmsConsignOrderNotifyRequest) GetGotAmount() int64 {
    return r.gotAmount
}

func (r *TaobaoWlbWmsConsignOrderNotifyRequest) SetPostfee(postfee int64) error {
    r.postfee = postfee
    r.Set("postfee", postfee)
    return nil
}

func (r TaobaoWlbWmsConsignOrderNotifyRequest) GetPostfee() int64 {
    return r.postfee
}

func (r *TaobaoWlbWmsConsignOrderNotifyRequest) SetServiceFee(serviceFee int64) error {
    r.serviceFee = serviceFee
    r.Set("service_fee", serviceFee)
    return nil
}

func (r TaobaoWlbWmsConsignOrderNotifyRequest) GetServiceFee() int64 {
    return r.serviceFee
}

func (r *TaobaoWlbWmsConsignOrderNotifyRequest) SetDeliverRequirements(deliverRequirements *Deliverrequirementswlbwmsconsignordernotify) error {
    r.deliverRequirements = deliverRequirements
    r.Set("deliver_requirements", deliverRequirements)
    return nil
}

func (r TaobaoWlbWmsConsignOrderNotifyRequest) GetDeliverRequirements() *Deliverrequirementswlbwmsconsignordernotify {
    return r.deliverRequirements
}

func (r *TaobaoWlbWmsConsignOrderNotifyRequest) SetReceiverInfo(receiverInfo *Receiverwlbwmsconsignordernotify) error {
    r.receiverInfo = receiverInfo
    r.Set("receiver_info", receiverInfo)
    return nil
}

func (r TaobaoWlbWmsConsignOrderNotifyRequest) GetReceiverInfo() *Receiverwlbwmsconsignordernotify {
    return r.receiverInfo
}

func (r *TaobaoWlbWmsConsignOrderNotifyRequest) SetSenderInfo(senderInfo *Senderwlbwmsconsignordernotify) error {
    r.senderInfo = senderInfo
    r.Set("sender_info", senderInfo)
    return nil
}

func (r TaobaoWlbWmsConsignOrderNotifyRequest) GetSenderInfo() *Senderwlbwmsconsignordernotify {
    return r.senderInfo
}

func (r *TaobaoWlbWmsConsignOrderNotifyRequest) SetOrderItemList(orderItemList []Orderitemlistwlbwmsconsignordernotify) error {
    r.orderItemList = orderItemList
    r.Set("order_item_list", orderItemList)
    return nil
}

func (r TaobaoWlbWmsConsignOrderNotifyRequest) GetOrderItemList() []Orderitemlistwlbwmsconsignordernotify {
    return r.orderItemList
}

func (r *TaobaoWlbWmsConsignOrderNotifyRequest) SetInvoiceInfoList(invoiceInfoList []Invoicelistwlbwmsconsignordernotify) error {
    r.invoiceInfoList = invoiceInfoList
    r.Set("invoice_info_list", invoiceInfoList)
    return nil
}

func (r TaobaoWlbWmsConsignOrderNotifyRequest) GetInvoiceInfoList() []Invoicelistwlbwmsconsignordernotify {
    return r.invoiceInfoList
}

func (r *TaobaoWlbWmsConsignOrderNotifyRequest) SetExtendFields(extendFields string) error {
    r.extendFields = extendFields
    r.Set("extend_fields", extendFields)
    return nil
}

func (r TaobaoWlbWmsConsignOrderNotifyRequest) GetExtendFields() string {
    return r.extendFields
}

func (r *TaobaoWlbWmsConsignOrderNotifyRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

func (r TaobaoWlbWmsConsignOrderNotifyRequest) GetRemark() string {
    return r.remark
}

