package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除网点 APIRequest
tmall.servicecenter.servicestore.deleteservicestore

删除网点信息。对于同一个服务商，通过 service_store_code 删除网点。
错误码
1, 服务商昵称无效
2, 缺少省份
3, 缺少城市
4, 缺少区域
5, 缺少详细地址
6, 传入地址不在标准地址库中，无法解析
7, 缺少网点编码
8, 缺少网点名称
9, 缺少网点电话
10, 网点已存在
11, 网点不存在
12, 系统错误
*/
type TmallServicecenterServicestoreDeleteservicestoreRequest struct {
    model.Params

    // 网点名称
    serviceStoreCode   string 

}

func NewTmallServicecenterServicestoreDeleteservicestoreRequest() *TmallServicecenterServicestoreDeleteservicestoreRequest{
    return &TmallServicecenterServicestoreDeleteservicestoreRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterServicestoreDeleteservicestoreRequest) GetApiMethodName() string {
    return "tmall.servicecenter.servicestore.deleteservicestore"
}

func (r TmallServicecenterServicestoreDeleteservicestoreRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterServicestoreDeleteservicestoreRequest) SetServiceStoreCode(serviceStoreCode string) error {
    r.serviceStoreCode = serviceStoreCode
    r.Set("service_store_code", serviceStoreCode)
    return nil
}

func (r TmallServicecenterServicestoreDeleteservicestoreRequest) GetServiceStoreCode() string {
    return r.serviceStoreCode
}

