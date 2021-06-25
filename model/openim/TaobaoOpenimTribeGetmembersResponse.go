package openim

import (
    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群成员获取 APIResponse
taobao.openim.tribe.getmembers

OPENIM群成员获取
*/
type TaobaoOpenimTribeGetmembersAPIResponse struct {
    model.CommonResponse
    Response *TaobaoOpenimTribeGetmembersResponse `json:"taobao_openim_tribe_getmembers_response,omitempty"`
}

type TaobaoOpenimTribeGetmembersResponse struct {

    // OPENIM群成员列表
    TribeUserList   []TribeUser `json:"tribe_user_list,omitempty"`

}
