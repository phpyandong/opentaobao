package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家地址库新增接口 APIResponse
taobao.logistics.address.add

通过此接口新增卖家地址库,卖家最多可添加5条地址库,新增第一条卖家地址，将会自动设为默认地址库
*/
type TaobaoLogisticsAddressAddAPIResponse struct {
    model.CommonResponse
    TaobaoLogisticsAddressAddResponse
}

type TaobaoLogisticsAddressAddResponse struct {
    XMLName xml.Name `xml:"logistics_address_add_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 只返回修改日期modify_date
    
    AddressResult   *AddressResult `json:"address_result,omitempty" xml:"address_result,omitempty"`

    
}
