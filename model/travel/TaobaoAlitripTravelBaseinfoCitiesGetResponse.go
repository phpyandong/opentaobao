package travel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】度假线路商品发布时基础信息获取接口：地址数据查询 APIResponse
taobao.alitrip.travel.baseinfo.cities.get

旅行度假新商品发布时可用的扩展接口，用于获取可用的出发地或目的地城市列表。
*/
type TaobaoAlitripTravelBaseinfoCitiesGetAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripTravelBaseinfoCitiesGetResponse
}

type TaobaoAlitripTravelBaseinfoCitiesGetResponse struct {
    XMLName xml.Name `xml:"alitrip_travel_baseinfo_cities_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 地区级联城市列表，返回数据为json数组结构的字符串
    
    IocInfos   string `json:"ioc_infos,omitempty" xml:"ioc_infos,omitempty"`

    
}
