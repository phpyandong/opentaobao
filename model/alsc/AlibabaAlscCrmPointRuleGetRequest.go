package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询积分规则 APIRequest
alibaba.alsc.crm.point.rule.get

新增积分规则查询接口,传入includeLogicalDelete和maxUpdateTime时走同步下行逻辑不然则走普通积分规则查询接口
*/
type AlibabaAlscCrmPointRuleGetRequest struct {
    model.Params

    // 入参
    paramQueryPointRuleOpenReq   *QueryPointRuleOpenReq 

}

func NewAlibabaAlscCrmPointRuleGetRequest() *AlibabaAlscCrmPointRuleGetRequest{
    return &AlibabaAlscCrmPointRuleGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmPointRuleGetRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.point.rule.get"
}

func (r AlibabaAlscCrmPointRuleGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmPointRuleGetRequest) SetParamQueryPointRuleOpenReq(paramQueryPointRuleOpenReq *QueryPointRuleOpenReq) error {
    r.paramQueryPointRuleOpenReq = paramQueryPointRuleOpenReq
    r.Set("param_query_point_rule_open_req", paramQueryPointRuleOpenReq)
    return nil
}

func (r AlibabaAlscCrmPointRuleGetRequest) GetParamQueryPointRuleOpenReq() *QueryPointRuleOpenReq {
    return r.paramQueryPointRuleOpenReq
}

