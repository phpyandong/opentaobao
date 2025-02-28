package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
获取计划关键词 
alibaba.scbp.ad.keyword.list.campaign.keywords

获取计划关键词
*/
func AlibabaScbpAdKeywordListCampaignKeywords(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordListCampaignKeywordsRequest, session string) (*scbp.AlibabaScbpAdKeywordListCampaignKeywordsAPIResponse, error) {
    var resp scbp.AlibabaScbpAdKeywordListCampaignKeywordsAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
