package interact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/interact"
)

/* 
天猫奖池鉴权接口 
alibaba.interact.isvlottery.isvdraw

鉴权接口，为tida.isvdraw接口鉴权
*/
func AlibabaInteractIsvlotteryIsvdraw(clt *core.SDKClient, req *interact.AlibabaInteractIsvlotteryIsvdrawRequest, session string) (*interact.AlibabaInteractIsvlotteryIsvdrawAPIResponse, error) {
    var resp interact.AlibabaInteractIsvlotteryIsvdrawAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
