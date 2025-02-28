package category

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商家被授权品牌列表和类目列表 APIRequest
taobao.itemcats.authorize.get

查询B商家被授权品牌列表、类目列表和 c 商家新品类目列表
*/
type TaobaoItemcatsAuthorizeGetRequest struct {
    model.Params

    // 需要返回的字段。目前支持有：<br/>brand.vid, brand.name, <br/>item_cat.cid, item_cat.name, item_cat.status,item_cat.sort_order,item_cat.parent_cid,item_cat.is_parent,<br/>xinpin_item_cat.cid, <br/>xinpin_item_cat.name, <br/>xinpin_item_cat.status,<br/>xinpin_item_cat.sort_order,<br/>xinpin_item_cat.parent_cid,<br/>xinpin_item_cat.is_parent
    fields   []string 

}

func NewTaobaoItemcatsAuthorizeGetRequest() *TaobaoItemcatsAuthorizeGetRequest{
    return &TaobaoItemcatsAuthorizeGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoItemcatsAuthorizeGetRequest) GetApiMethodName() string {
    return "taobao.itemcats.authorize.get"
}

func (r TaobaoItemcatsAuthorizeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoItemcatsAuthorizeGetRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoItemcatsAuthorizeGetRequest) GetFields() []string {
    return r.fields
}

