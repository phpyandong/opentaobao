package subuser

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询指定账户的子账号列表 APIResponse
taobao.sellercenter.subusers.get

根据主账号nick查询该账号下所有的子账号列表，只能查询属于自己的账号信息 (主账号以及所属子账号)
*/
type TaobaoSellercenterSubusersGetAPIResponse struct {
    model.CommonResponse
    TaobaoSellercenterSubusersGetResponse
}

type TaobaoSellercenterSubusersGetResponse struct {
    XMLName xml.Name `xml:"sellercenter_subusers_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 子账号基本信息列表。具体信息为id、子账号用户名、主账号id、主账号昵称、当前状态值、是否分流
    
    Subusers   []SubUserInfo `json:"subusers,omitempty" xml:"subusers>sub_user_info,omitempty"`
    
    
}
