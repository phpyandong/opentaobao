package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
推广计划报表基础数据对象 
taobao.simba.rpt.campaignbase.get

推广计划报表基础数据对象
*/
func TaobaoSimbaRptCampaignbaseGet(clt *core.SDKClient, req *simba.TaobaoSimbaRptCampaignbaseGetRequest, session string) (*simba.TaobaoSimbaRptCampaignbaseGetAPIResponse, error) {
    var resp simba.TaobaoSimbaRptCampaignbaseGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
