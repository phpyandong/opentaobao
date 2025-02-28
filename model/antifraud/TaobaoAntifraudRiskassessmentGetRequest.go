package antifraud

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
反欺诈风险识别 APIRequest
taobao.antifraud.riskassessment.get

反欺诈服务是阿里大数据风控服务能力的对外输出，通过用户信誉、行为分析精准识别可信用户和风险用户并实时防御，解决交易、支付、活动等关键业务环节存在的欺诈威胁，保护企业品牌和数据，降低企业经济损失
*/
type TaobaoAntifraudRiskassessmentGetRequest struct {
    model.Params

    // 风控查询参数
    collinadataContext   *CollinadataContext 

}

func NewTaobaoAntifraudRiskassessmentGetRequest() *TaobaoAntifraudRiskassessmentGetRequest{
    return &TaobaoAntifraudRiskassessmentGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAntifraudRiskassessmentGetRequest) GetApiMethodName() string {
    return "taobao.antifraud.riskassessment.get"
}

func (r TaobaoAntifraudRiskassessmentGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAntifraudRiskassessmentGetRequest) SetCollinadataContext(collinadataContext *CollinadataContext) error {
    r.collinadataContext = collinadataContext
    r.Set("collinadata_context", collinadataContext)
    return nil
}

func (r TaobaoAntifraudRiskassessmentGetRequest) GetCollinadataContext() *CollinadataContext {
    return r.collinadataContext
}

