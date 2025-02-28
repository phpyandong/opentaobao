package flight

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
航变信息录入接口 APIResponse
taobao.alitrip.flightchange.add

代理商调用航变录入接口,
    简要描述:完成航变信息的自动录入后飞猪机票平台会发生的动作是匹配所有该代理人的订单,如果接口参数指定了飞猪机票订单号，发生的动作是匹配该代理人的指定订单；
找到与航变航班相关的订单给旅客下发航变短信并出发IVR自动外呼旅客。
*/
type TaobaoAlitripFlightchangeAddAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripFlightchangeAddResponse
}

type TaobaoAlitripFlightchangeAddResponse struct {
    XMLName xml.Name `xml:"alitrip_flightchange_add_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TaobaoAlitripFlightchangeAddResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
