package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
品牌营销活动用户上传 
tmall.ccf.crowd.activityuser.upload

搜集ISV的活动用户信息，将其沉淀为活动人群数据
*/
func TmallCcfCrowdActivityuserUpload(clt *core.SDKClient, req *promotion.TmallCcfCrowdActivityuserUploadRequest, session string) (*promotion.TmallCcfCrowdActivityuserUploadAPIResponse, error) {
    var resp promotion.TmallCcfCrowdActivityuserUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
