package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询日消耗预算 APIResponse
alibaba.scbp.account.budget.get

查询日消耗预算
*/
type AlibabaScbpAccountBudgetGetAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAccountBudgetGetResponse
}

type AlibabaScbpAccountBudgetGetResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_account_budget_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回今日预算
    
    Budget   string `json:"budget,omitempty" xml:"budget,omitempty"`

    
}
