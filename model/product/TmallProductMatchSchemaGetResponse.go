package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取匹配产品规则 APIResponse
tmall.product.match.schema.get

ISV发布商品前，需要先查找到产品ID，这个接口返回查找产品规则入参规则
*/
type TmallProductMatchSchemaGetAPIResponse struct {
    model.CommonResponse
    TmallProductMatchSchemaGetResponse
}

type TmallProductMatchSchemaGetResponse struct {
    XMLName xml.Name `xml:"tmall_product_match_schema_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回匹配product的规则文档
    
    MatchResult   string `json:"match_result,omitempty" xml:"match_result,omitempty"`

    
}
