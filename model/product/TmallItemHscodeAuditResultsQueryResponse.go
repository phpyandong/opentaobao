package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品hscode信息审核状态查询接口 APIResponse
tmall.item.hscode.audit.results.query

通过此接口查询天猫跨境商品的hscode信息审核状态，卖家可以参考返回结果判断是否需要调整商品hscode相关信息。
*/
type TmallItemHscodeAuditResultsQueryAPIResponse struct {
    model.CommonResponse
    TmallItemHscodeAuditResultsQueryResponse
}

type TmallItemHscodeAuditResultsQueryResponse struct {
    XMLName xml.Name `xml:"tmall_item_hscode_audit_results_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品或sku的hscode信息审核状态。
    
    ResultList   []HscodeAuditInfo `json:"result_list,omitempty" xml:"result_list>hscode_audit_info,omitempty"`
    
    
}
