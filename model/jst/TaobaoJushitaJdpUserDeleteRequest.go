package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除数据推送用户 APIRequest
taobao.jushita.jdp.user.delete

删除应用的数据推送用户，用户被删除后，重新添加时会重新同步历史数据。
*/
type TaobaoJushitaJdpUserDeleteRequest struct {
    model.Params

    // 要删除用户的昵称
    nick   string 

    // 需要删除的用户编号
    userId   int64 

}

func NewTaobaoJushitaJdpUserDeleteRequest() *TaobaoJushitaJdpUserDeleteRequest{
    return &TaobaoJushitaJdpUserDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJushitaJdpUserDeleteRequest) GetApiMethodName() string {
    return "taobao.jushita.jdp.user.delete"
}

func (r TaobaoJushitaJdpUserDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJushitaJdpUserDeleteRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoJushitaJdpUserDeleteRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoJushitaJdpUserDeleteRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r TaobaoJushitaJdpUserDeleteRequest) GetUserId() int64 {
    return r.userId
}

