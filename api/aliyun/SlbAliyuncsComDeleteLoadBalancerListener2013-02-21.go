package aliyun

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aliyun"
)

/* 
删除 slb listener 
slb.aliyuncs.com.DeleteLoadBalancerListener.2013-02-21

delete_vip
*/
func SlbAliyuncsComDeleteLoadBalancerListener2013-02-21(clt *core.SDKClient, req *aliyun.SlbAliyuncsComDeleteLoadBalancerListener2013-02-21Request, session string) (*aliyun.SlbAliyuncsComDeleteLoadBalancerListener2013-02-21APIResponse, error) {
    var resp aliyun.SlbAliyuncsComDeleteLoadBalancerListener2013-02-21APIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
