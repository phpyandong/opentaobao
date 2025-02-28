package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关键词过滤 APIRequest
taobao.openim.snfilterword.setfilter

设置openim关键词过滤
*/
type TaobaoOpenimSnfilterwordSetfilterRequest struct {
    model.Params

    // 上传者身份信息，区分不同上传者;只是记录，没有身份校验功能
    creator   string 

    // 需要过滤的关键词
    filterword   string 

    // 过滤原因描述
    desc   string 

}

func NewTaobaoOpenimSnfilterwordSetfilterRequest() *TaobaoOpenimSnfilterwordSetfilterRequest{
    return &TaobaoOpenimSnfilterwordSetfilterRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenimSnfilterwordSetfilterRequest) GetApiMethodName() string {
    return "taobao.openim.snfilterword.setfilter"
}

func (r TaobaoOpenimSnfilterwordSetfilterRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenimSnfilterwordSetfilterRequest) SetCreator(creator string) error {
    r.creator = creator
    r.Set("creator", creator)
    return nil
}

func (r TaobaoOpenimSnfilterwordSetfilterRequest) GetCreator() string {
    return r.creator
}

func (r *TaobaoOpenimSnfilterwordSetfilterRequest) SetFilterword(filterword string) error {
    r.filterword = filterword
    r.Set("filterword", filterword)
    return nil
}

func (r TaobaoOpenimSnfilterwordSetfilterRequest) GetFilterword() string {
    return r.filterword
}

func (r *TaobaoOpenimSnfilterwordSetfilterRequest) SetDesc(desc string) error {
    r.desc = desc
    r.Set("desc", desc)
    return nil
}

func (r TaobaoOpenimSnfilterwordSetfilterRequest) GetDesc() string {
    return r.desc
}

