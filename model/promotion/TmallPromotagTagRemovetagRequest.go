package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除标签定义 APIRequest
tmall.promotag.tag.removetag

用于删除标签定义，但是要确保目前该标签没有人群在使用。
*/
type TmallPromotagTagRemovetagRequest struct {
    model.Params

    // 需要删除的标签id
    tagId   int64 

}

func NewTmallPromotagTagRemovetagRequest() *TmallPromotagTagRemovetagRequest{
    return &TmallPromotagTagRemovetagRequest{
        Params: model.NewParams(),
    }
}

func (r TmallPromotagTagRemovetagRequest) GetApiMethodName() string {
    return "tmall.promotag.tag.removetag"
}

func (r TmallPromotagTagRemovetagRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallPromotagTagRemovetagRequest) SetTagId(tagId int64) error {
    r.tagId = tagId
    r.Set("tag_id", tagId)
    return nil
}

func (r TmallPromotagTagRemovetagRequest) GetTagId() int64 {
    return r.tagId
}

