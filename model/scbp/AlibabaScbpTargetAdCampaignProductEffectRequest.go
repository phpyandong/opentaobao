package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-获取计划中产品推广效果 APIRequest
alibaba.scbp.target.ad.campaign.product.effect

定向推广-获取计划中产品推广效果
*/
type AlibabaScbpTargetAdCampaignProductEffectRequest struct {
    model.Params

    // TopP4pQuickEffectQuery
    topP4pQuickEffectQuery   *TopP4pQuickEffectQuery 

}

func NewAlibabaScbpTargetAdCampaignProductEffectRequest() *AlibabaScbpTargetAdCampaignProductEffectRequest{
    return &AlibabaScbpTargetAdCampaignProductEffectRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpTargetAdCampaignProductEffectRequest) GetApiMethodName() string {
    return "alibaba.scbp.target.ad.campaign.product.effect"
}

func (r AlibabaScbpTargetAdCampaignProductEffectRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpTargetAdCampaignProductEffectRequest) SetTopP4pQuickEffectQuery(topP4pQuickEffectQuery *TopP4pQuickEffectQuery) error {
    r.topP4pQuickEffectQuery = topP4pQuickEffectQuery
    r.Set("top_p4p_quick_effect_query", topP4pQuickEffectQuery)
    return nil
}

func (r AlibabaScbpTargetAdCampaignProductEffectRequest) GetTopP4pQuickEffectQuery() *TopP4pQuickEffectQuery {
    return r.topP4pQuickEffectQuery
}

