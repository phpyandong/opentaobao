package media

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询某个用户图片空间的使用情况 APIResponse
taobao.vas.service.getServTimes

查询某个用户图片空间的使用情况
*/
type TaobaoVasServiceGetServTimesAPIResponse struct {
    model.CommonResponse
    TaobaoVasServiceGetServTimesResponse
}

type TaobaoVasServiceGetServTimesResponse struct {
    XMLName xml.Name `xml:"vas_service_getServTimes_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 总次数（容量）
    
    TotalNum   int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`

    
    // 剩余次数（容量）
    
    LeftNum   int64 `json:"left_num,omitempty" xml:"left_num,omitempty"`

    
}
