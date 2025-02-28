package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
修改门店信息接口 
alibaba.ele.fengniao.chainstore.update

修改门店的经纬度，文本地址，电话，门店名
*/
func AlibabaEleFengniaoChainstoreUpdate(clt *core.SDKClient, req *logistic.AlibabaEleFengniaoChainstoreUpdateRequest, session string) (*logistic.AlibabaEleFengniaoChainstoreUpdateAPIResponse, error) {
    var resp logistic.AlibabaEleFengniaoChainstoreUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
