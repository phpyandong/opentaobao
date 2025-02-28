package qt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝服务订购关系查询 APIRequest
taobao.ts.subscribe.get

ts订购关系状态查询. 暂只支持1口价服务.
*/
type TaobaoTsSubscribeGetRequest struct {
    model.Params

    // 服务收费项code
    servcieItemCode   string 

    // 订购用户昵称
    nick   string 

}

func NewTaobaoTsSubscribeGetRequest() *TaobaoTsSubscribeGetRequest{
    return &TaobaoTsSubscribeGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTsSubscribeGetRequest) GetApiMethodName() string {
    return "taobao.ts.subscribe.get"
}

func (r TaobaoTsSubscribeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTsSubscribeGetRequest) SetServcieItemCode(servcieItemCode string) error {
    r.servcieItemCode = servcieItemCode
    r.Set("servcie_item_code", servcieItemCode)
    return nil
}

func (r TaobaoTsSubscribeGetRequest) GetServcieItemCode() string {
    return r.servcieItemCode
}

func (r *TaobaoTsSubscribeGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoTsSubscribeGetRequest) GetNick() string {
    return r.nick
}

