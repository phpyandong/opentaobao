package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
工单查询接口 APIResponse
tmall.servicecenter.workcard.query

工单查询接口
*/
type TmallServicecenterWorkcardQueryAPIResponse struct {
    model.CommonResponse
    TmallServicecenterWorkcardQueryResponse
}

type TmallServicecenterWorkcardQueryResponse struct {
    XMLName xml.Name `xml:"tmall_servicecenter_workcard_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 请求结果
    
    Result   *TmallServicecenterWorkcardQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
