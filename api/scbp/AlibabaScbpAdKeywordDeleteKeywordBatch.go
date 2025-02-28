package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
删除关键词 
alibaba.scbp.ad.keyword.delete.keyword.batch

删除关键词
*/
func AlibabaScbpAdKeywordDeleteKeywordBatch(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordDeleteKeywordBatchRequest, session string) (*scbp.AlibabaScbpAdKeywordDeleteKeywordBatchAPIResponse, error) {
    var resp scbp.AlibabaScbpAdKeywordDeleteKeywordBatchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
