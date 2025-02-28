package antifraud

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
反欺诈风险识别 APIResponse
taobao.antifraud.riskassessment.get

反欺诈服务是阿里大数据风控服务能力的对外输出，通过用户信誉、行为分析精准识别可信用户和风险用户并实时防御，解决交易、支付、活动等关键业务环节存在的欺诈威胁，保护企业品牌和数据，降低企业经济损失
*/
type TaobaoAntifraudRiskassessmentGetAPIResponse struct {
    model.CommonResponse
    TaobaoAntifraudRiskassessmentGetResponse
}

type TaobaoAntifraudRiskassessmentGetResponse struct {
    XMLName xml.Name `xml:"antifraud_riskassessment_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    RiskResult   *ResultWrapper `json:"risk_result,omitempty" xml:"risk_result,omitempty"`

    
}
