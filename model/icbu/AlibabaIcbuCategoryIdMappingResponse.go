package icbu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新旧属性的映射 APIResponse
alibaba.icbu.category.id.mapping

商品发布接口升级，需要传入新的类目。这个接口 根据旧的类目id，获取新的类目id
*/
type AlibabaIcbuCategoryIdMappingAPIResponse struct {
    model.CommonResponse
    AlibabaIcbuCategoryIdMappingResponse
}

type AlibabaIcbuCategoryIdMappingResponse struct {
    XMLName xml.Name `xml:"alibaba_icbu_category_id_mapping_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 转化的类目id
    
    MappingResult   int64 `json:"mapping_result,omitempty" xml:"mapping_result,omitempty"`

    
}
