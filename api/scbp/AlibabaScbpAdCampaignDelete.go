package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
删除计划 
alibaba.scbp.ad.campaign.delete

删除计划
*/
func AlibabaScbpAdCampaignDelete(clt *core.SDKClient, req *scbp.AlibabaScbpAdCampaignDeleteRequest, session string) (*scbp.AlibabaScbpAdCampaignDeleteAPIResponse, error) {
    var resp scbp.AlibabaScbpAdCampaignDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
