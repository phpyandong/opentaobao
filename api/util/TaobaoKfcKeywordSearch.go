package util

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/util"
)

/* 
关键词过滤匹配 
taobao.kfc.keyword.search

对输入的文本信息进行禁忌关键词匹配，返回匹配的结果
*/
func TaobaoKfcKeywordSearch(clt *core.SDKClient, req *util.TaobaoKfcKeywordSearchRequest, session string) (*util.TaobaoKfcKeywordSearchAPIResponse, error) {
    var resp util.TaobaoKfcKeywordSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
