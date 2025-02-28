package subuser

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询指定的子账号的权限和角色信息 APIRequest
taobao.sellercenter.subuser.permissions.roles.get

查询指定的子账号的被直接赋予的权限信息和角色信息。<br/>返回对象中包括直接赋予子账号的权限点信息、被赋予的角色以及角色的对应权限点信息。
*/
type TaobaoSellercenterSubuserPermissionsRolesGetRequest struct {
    model.Params

    // 子账号昵称(子账号标识)
    nick   string 

}

func NewTaobaoSellercenterSubuserPermissionsRolesGetRequest() *TaobaoSellercenterSubuserPermissionsRolesGetRequest{
    return &TaobaoSellercenterSubuserPermissionsRolesGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSellercenterSubuserPermissionsRolesGetRequest) GetApiMethodName() string {
    return "taobao.sellercenter.subuser.permissions.roles.get"
}

func (r TaobaoSellercenterSubuserPermissionsRolesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSellercenterSubuserPermissionsRolesGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSellercenterSubuserPermissionsRolesGetRequest) GetNick() string {
    return r.nick
}

