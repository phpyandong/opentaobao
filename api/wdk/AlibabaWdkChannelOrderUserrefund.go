package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
用户发起售后退款(整单/部分) 
alibaba.wdk.channel.order.userrefund

用户发起售后退款(整单/部分)
*/
func AlibabaWdkChannelOrderUserrefund(clt *core.SDKClient, req *wdk.AlibabaWdkChannelOrderUserrefundRequest, session string) (*wdk.AlibabaWdkChannelOrderUserrefundAPIResponse, error) {
    var resp wdk.AlibabaWdkChannelOrderUserrefundAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
