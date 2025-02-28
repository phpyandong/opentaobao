package iot

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/iot"
)

/* 
查询音频 
alibaba.ailabs.aligenie.tracks.search

搜索类目下的音频信息
*/
func AlibabaAilabsAligenieTracksSearch(clt *core.SDKClient, req *iot.AlibabaAilabsAligenieTracksSearchRequest, session string) (*iot.AlibabaAilabsAligenieTracksSearchAPIResponse, error) {
    var resp iot.AlibabaAilabsAligenieTracksSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
