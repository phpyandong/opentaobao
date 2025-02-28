package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取得一个推广组的所有关键词和类目出价的质量得分 APIRequest
taobao.simba.keywordscat.qscore.get

取得一个推广组的所有关键词和类目出价的质量得分列表
*/
type TaobaoSimbaKeywordscatQscoreGetRequest struct {
    model.Params

    // 主人昵称
    nick   string 

    // 推广组Id
    adgroupId   int64 

}

func NewTaobaoSimbaKeywordscatQscoreGetRequest() *TaobaoSimbaKeywordscatQscoreGetRequest{
    return &TaobaoSimbaKeywordscatQscoreGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaKeywordscatQscoreGetRequest) GetApiMethodName() string {
    return "taobao.simba.keywordscat.qscore.get"
}

func (r TaobaoSimbaKeywordscatQscoreGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaKeywordscatQscoreGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaKeywordscatQscoreGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaKeywordscatQscoreGetRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoSimbaKeywordscatQscoreGetRequest) GetAdgroupId() int64 {
    return r.adgroupId
}

