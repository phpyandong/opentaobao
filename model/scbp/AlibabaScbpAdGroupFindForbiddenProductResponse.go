package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询屏蔽品 APIResponse
alibaba.scbp.ad.group.find.forbidden.product

查询屏蔽品
*/
type AlibabaScbpAdGroupFindForbiddenProductAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdGroupFindForbiddenProductResponse
}

type AlibabaScbpAdGroupFindForbiddenProductResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_group_find_forbidden_product_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回列表
    
    ResultList   []ForbiddenProductDto `json:"result_list,omitempty" xml:"result_list>forbidden_product_dto,omitempty"`
    
    
}
