package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更改交易的收货地址 APIResponse
taobao.trade.shippingaddress.update

只能更新一笔交易里面的买家收货地址 <br/>只能更新发货前（即买家已付款，等待卖家发货状态）的交易的买家收货地址 <br/>更新后的发货地址可以通过taobao.trade.fullinfo.get查到 <br/>参数中所说的字节为GBK编码的（英文和数字占1字节，中文占2字节）
*/
type TaobaoTradeShippingaddressUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoTradeShippingaddressUpdateResponse
}

type TaobaoTradeShippingaddressUpdateResponse struct {
    XMLName xml.Name `xml:"trade_shippingaddress_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 交易结构
    
    Trade   *Trade `json:"trade,omitempty" xml:"trade,omitempty"`

    
}
