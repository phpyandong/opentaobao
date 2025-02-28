package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量删除分组 APIRequest
taobao.crm.members.groups.batchdelete

批量删除多个会员的公共分组，接口返回删除是否成功，该接口只删除多个会员的公共分组，不是公共分组的，不进行删除。如果入参只输入一个会员，则表示删除该会员的某些分组。
*/
type TaobaoCrmMembersGroupsBatchdeleteRequest struct {
    model.Params

    // 买家昵称列表
    buyerNicks   []string 

    // 会员需要删除的分组
    groupIds   []int64 

}

func NewTaobaoCrmMembersGroupsBatchdeleteRequest() *TaobaoCrmMembersGroupsBatchdeleteRequest{
    return &TaobaoCrmMembersGroupsBatchdeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoCrmMembersGroupsBatchdeleteRequest) GetApiMethodName() string {
    return "taobao.crm.members.groups.batchdelete"
}

func (r TaobaoCrmMembersGroupsBatchdeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoCrmMembersGroupsBatchdeleteRequest) SetBuyerNicks(buyerNicks []string) error {
    r.buyerNicks = buyerNicks
    r.Set("buyer_nicks", buyerNicks)
    return nil
}

func (r TaobaoCrmMembersGroupsBatchdeleteRequest) GetBuyerNicks() []string {
    return r.buyerNicks
}

func (r *TaobaoCrmMembersGroupsBatchdeleteRequest) SetGroupIds(groupIds []int64) error {
    r.groupIds = groupIds
    r.Set("group_ids", groupIds)
    return nil
}

func (r TaobaoCrmMembersGroupsBatchdeleteRequest) GetGroupIds() []int64 {
    return r.groupIds
}

