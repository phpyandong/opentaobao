package subuser

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取指定用户的权限集合 APIRequest
taobao.sellercenter.user.permissions.get

获取指定用户的权限集合，并不组装成树。如果是主账号，返回所有的权限列表；如果是子账号，返回所有已授权的权限。只能查询属于自己的账号信息 (如果是主账号，则是主账号以及所属子账号，如果是子账号则是对应主账号以及所属子账号)
*/
type TaobaoSellercenterUserPermissionsGetRequest struct {
    model.Params

    // 用户标识，次入参必须为子账号比如zhangsan:cool。如果只输入主账号zhangsan，将报错。
    nick   string 

}

func NewTaobaoSellercenterUserPermissionsGetRequest() *TaobaoSellercenterUserPermissionsGetRequest{
    return &TaobaoSellercenterUserPermissionsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSellercenterUserPermissionsGetRequest) GetApiMethodName() string {
    return "taobao.sellercenter.user.permissions.get"
}

func (r TaobaoSellercenterUserPermissionsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSellercenterUserPermissionsGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSellercenterUserPermissionsGetRequest) GetNick() string {
    return r.nick
}

