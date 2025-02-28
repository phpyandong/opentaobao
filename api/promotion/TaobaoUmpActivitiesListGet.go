package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
营销活动列表查询 
taobao.ump.activities.list.get

按照营销活动id的列表ids，查询对应的营销活动列表。
*/
func TaobaoUmpActivitiesListGet(clt *core.SDKClient, req *promotion.TaobaoUmpActivitiesListGetRequest, session string) (*promotion.TaobaoUmpActivitiesListGetAPIResponse, error) {
    var resp promotion.TaobaoUmpActivitiesListGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
