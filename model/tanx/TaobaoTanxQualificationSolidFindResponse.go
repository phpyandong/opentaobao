package tanx

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
客户固态共享资质 APIResponse
taobao.tanx.qualification.solid.find

接口会返回该广告主下的所有审核通过并且可被共享的资质，这些资质在过期之前可以不需要再次上传。
*/
type TaobaoTanxQualificationSolidFindAPIResponse struct {
    model.CommonResponse
    TaobaoTanxQualificationSolidFindResponse
}

type TaobaoTanxQualificationSolidFindResponse struct {
    XMLName xml.Name `xml:"tanx_qualification_solid_find_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 返回固化资质列表
    
    QualificationList   []QualificationDto `json:"qualification_list,omitempty" xml:"qualification_list>qualification_dto,omitempty"`
    
    
    // 返回查询总数
    
    Count   string `json:"count,omitempty" xml:"count,omitempty"`

    
}
