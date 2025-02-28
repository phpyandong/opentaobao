package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查看功能权限 APIRequest
taobao.simba.customers.sid.get

查询用户是否拥有某个功能权限
*/
type TaobaoSimbaCustomersSidGetRequest struct {
    model.Params

}

func NewTaobaoSimbaCustomersSidGetRequest() *TaobaoSimbaCustomersSidGetRequest{
    return &TaobaoSimbaCustomersSidGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaCustomersSidGetRequest) GetApiMethodName() string {
    return "taobao.simba.customers.sid.get"
}

func (r TaobaoSimbaCustomersSidGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


