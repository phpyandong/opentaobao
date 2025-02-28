package hotel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店详细信息查询 APIResponse
taobao.xhotel.info.list.get

获取酒店详情信息
*/
type TaobaoXhotelInfoListGetAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelInfoListGetResponse
}

type TaobaoXhotelInfoListGetResponse struct {
    XMLName xml.Name `xml:"xhotel_info_list_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 酒店总数
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`

    
    // 标准酒店信息
    
    Hotels   []SHotelInfoObject `json:"hotels,omitempty" xml:"hotels>s_hotel_info_object,omitempty"`
    
    
}
