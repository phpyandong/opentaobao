package travel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/travel"
)

/* 
发布SKU信息（如果properties重复 则更新） 
alitrip.travel.gereralsku.update

发布SKU信息（如果properties重复 则更新）
*/
func AlitripTravelGereralskuUpdate(clt *core.SDKClient, req *travel.AlitripTravelGereralskuUpdateRequest, session string) (*travel.AlitripTravelGereralskuUpdateAPIResponse, error) {
    var resp travel.AlitripTravelGereralskuUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
