package jms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询ONS分组 APIRequest
taobao.jushita.jms.group.get

查询当前appkey在ONS中已有的分组
*/
type TaobaoJushitaJmsGroupGetRequest struct {
    model.Params

    // 页码
    pageNo   int64 

    // 每页返回多少个分组
    pageSize   int64 

    // 要查询分组的名称，多个分组用半角逗号分隔，不传代表查询所有分组信息，但不会返回组下面的用户信息。如果应用没有设置分组则返回空。组名不能以default开头，default开头是系统默认的组。
    groupNames   []string 

}

func NewTaobaoJushitaJmsGroupGetRequest() *TaobaoJushitaJmsGroupGetRequest{
    return &TaobaoJushitaJmsGroupGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJushitaJmsGroupGetRequest) GetApiMethodName() string {
    return "taobao.jushita.jms.group.get"
}

func (r TaobaoJushitaJmsGroupGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJushitaJmsGroupGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoJushitaJmsGroupGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoJushitaJmsGroupGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoJushitaJmsGroupGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoJushitaJmsGroupGetRequest) SetGroupNames(groupNames []string) error {
    r.groupNames = groupNames
    r.Set("group_names", groupNames)
    return nil
}

func (r TaobaoJushitaJmsGroupGetRequest) GetGroupNames() []string {
    return r.groupNames
}

