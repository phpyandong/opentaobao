package xiami

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xiami"
)

/* 
热门歌曲 
alibaba.xiami.api.artist.hotSongs.get

热门歌曲
*/
func AlibabaXiamiApiArtistHotSongsGet(clt *core.SDKClient, req *xiami.AlibabaXiamiApiArtistHotSongsGetRequest, session string) (*xiami.AlibabaXiamiApiArtistHotSongsGetAPIResponse, error) {
    var resp xiami.AlibabaXiamiApiArtistHotSongsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
