package hotel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店搜索List接口 APIResponse
alitrip.hotel.search.list.get

酒店搜索List接口
*/
type AlitripHotelSearchListGetAPIResponse struct {
    model.CommonResponse
    AlitripHotelSearchListGetResponse
}

type AlitripHotelSearchListGetResponse struct {
    XMLName xml.Name `xml:"alitrip_hotel_search_list_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlitripHotelSearchListGetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
