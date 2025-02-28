package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更改交易的收货地址 APIRequest
taobao.trade.shippingaddress.update

只能更新一笔交易里面的买家收货地址 <br/>只能更新发货前（即买家已付款，等待卖家发货状态）的交易的买家收货地址 <br/>更新后的发货地址可以通过taobao.trade.fullinfo.get查到 <br/>参数中所说的字节为GBK编码的（英文和数字占1字节，中文占2字节）
*/
type TaobaoTradeShippingaddressUpdateRequest struct {
    model.Params

    // 交易编号。
    tid   int64 

    // 收货人全名。最大长度为50个字节。
    receiverName   string 

    // 固定电话。最大长度为30个字节。
    receiverPhone   string 

    // 移动电话。最大长度为11个字节。
    receiverMobile   string 

    // 省份。最大长度为32个字节。如：浙江
    receiverState   string 

    // 城市。最大长度为32个字节。如：杭州
    receiverCity   string 

    // 区/县。最大长度为32个字节。如：西湖区
    receiverDistrict   string 

    // 收货地址。最大长度为228个字节。
    receiverAddress   string 

    // 邮政编码。必须由6个数字组成。
    receiverZip   string 

    // 四级地址。最大长度为32个字节。如：五常街道
    receiverTown   string 

}

func NewTaobaoTradeShippingaddressUpdateRequest() *TaobaoTradeShippingaddressUpdateRequest{
    return &TaobaoTradeShippingaddressUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTradeShippingaddressUpdateRequest) GetApiMethodName() string {
    return "taobao.trade.shippingaddress.update"
}

func (r TaobaoTradeShippingaddressUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTradeShippingaddressUpdateRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoTradeShippingaddressUpdateRequest) GetTid() int64 {
    return r.tid
}

func (r *TaobaoTradeShippingaddressUpdateRequest) SetReceiverName(receiverName string) error {
    r.receiverName = receiverName
    r.Set("receiver_name", receiverName)
    return nil
}

func (r TaobaoTradeShippingaddressUpdateRequest) GetReceiverName() string {
    return r.receiverName
}

func (r *TaobaoTradeShippingaddressUpdateRequest) SetReceiverPhone(receiverPhone string) error {
    r.receiverPhone = receiverPhone
    r.Set("receiver_phone", receiverPhone)
    return nil
}

func (r TaobaoTradeShippingaddressUpdateRequest) GetReceiverPhone() string {
    return r.receiverPhone
}

func (r *TaobaoTradeShippingaddressUpdateRequest) SetReceiverMobile(receiverMobile string) error {
    r.receiverMobile = receiverMobile
    r.Set("receiver_mobile", receiverMobile)
    return nil
}

func (r TaobaoTradeShippingaddressUpdateRequest) GetReceiverMobile() string {
    return r.receiverMobile
}

func (r *TaobaoTradeShippingaddressUpdateRequest) SetReceiverState(receiverState string) error {
    r.receiverState = receiverState
    r.Set("receiver_state", receiverState)
    return nil
}

func (r TaobaoTradeShippingaddressUpdateRequest) GetReceiverState() string {
    return r.receiverState
}

func (r *TaobaoTradeShippingaddressUpdateRequest) SetReceiverCity(receiverCity string) error {
    r.receiverCity = receiverCity
    r.Set("receiver_city", receiverCity)
    return nil
}

func (r TaobaoTradeShippingaddressUpdateRequest) GetReceiverCity() string {
    return r.receiverCity
}

func (r *TaobaoTradeShippingaddressUpdateRequest) SetReceiverDistrict(receiverDistrict string) error {
    r.receiverDistrict = receiverDistrict
    r.Set("receiver_district", receiverDistrict)
    return nil
}

func (r TaobaoTradeShippingaddressUpdateRequest) GetReceiverDistrict() string {
    return r.receiverDistrict
}

func (r *TaobaoTradeShippingaddressUpdateRequest) SetReceiverAddress(receiverAddress string) error {
    r.receiverAddress = receiverAddress
    r.Set("receiver_address", receiverAddress)
    return nil
}

func (r TaobaoTradeShippingaddressUpdateRequest) GetReceiverAddress() string {
    return r.receiverAddress
}

func (r *TaobaoTradeShippingaddressUpdateRequest) SetReceiverZip(receiverZip string) error {
    r.receiverZip = receiverZip
    r.Set("receiver_zip", receiverZip)
    return nil
}

func (r TaobaoTradeShippingaddressUpdateRequest) GetReceiverZip() string {
    return r.receiverZip
}

func (r *TaobaoTradeShippingaddressUpdateRequest) SetReceiverTown(receiverTown string) error {
    r.receiverTown = receiverTown
    r.Set("receiver_town", receiverTown)
    return nil
}

func (r TaobaoTradeShippingaddressUpdateRequest) GetReceiverTown() string {
    return r.receiverTown
}

