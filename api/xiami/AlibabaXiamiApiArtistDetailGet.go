package xiami

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xiami"
)

/* 
艺人详情 
alibaba.xiami.api.artist.detail.get

艺人详情
*/
func AlibabaXiamiApiArtistDetailGet(clt *core.SDKClient, req *xiami.AlibabaXiamiApiArtistDetailGetRequest, session string) (*xiami.AlibabaXiamiApiArtistDetailGetAPIResponse, error) {
    var resp xiami.AlibabaXiamiApiArtistDetailGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
