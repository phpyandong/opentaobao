package travel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】度假线路商品发布时基础信息获取接口：邮轮扩展信息获取 APIResponse
taobao.alitrip.travel.baseinfo.cruise.get

旅行度假新商品发布时可用的扩展接口，用于获取邮轮类目相关扩展信息。
*/
type TaobaoAlitripTravelBaseinfoCruiseGetAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripTravelBaseinfoCruiseGetResponse
}

type TaobaoAlitripTravelBaseinfoCruiseGetResponse struct {
    XMLName xml.Name `xml:"alitrip_travel_baseinfo_cruise_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 邮轮类目扩展信息的json格式字符串
    
    CruiseExtInfos   string `json:"cruise_ext_infos,omitempty" xml:"cruise_ext_infos,omitempty"`

    
}
