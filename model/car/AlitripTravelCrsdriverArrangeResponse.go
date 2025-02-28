package car

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
CRS接送机商家派司机接口 APIResponse
alitrip.travel.crsdriver.arrange

提供给CRS接送机商家派司机的API
*/
type AlitripTravelCrsdriverArrangeAPIResponse struct {
    model.CommonResponse
    AlitripTravelCrsdriverArrangeResponse
}

type AlitripTravelCrsdriverArrangeResponse struct {
    XMLName xml.Name `xml:"alitrip_travel_crsdriver_arrange_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果message
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 返回结果code
    
    MessageCode   int64 `json:"message_code,omitempty" xml:"message_code,omitempty"`

    
}
