package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
创建推广单元 
alibaba.scbp.ad.group.create.ad.group.batch

创建推广单元
*/
func AlibabaScbpAdGroupCreateAdGroupBatch(clt *core.SDKClient, req *scbp.AlibabaScbpAdGroupCreateAdGroupBatchRequest, session string) (*scbp.AlibabaScbpAdGroupCreateAdGroupBatchAPIResponse, error) {
    var resp scbp.AlibabaScbpAdGroupCreateAdGroupBatchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
