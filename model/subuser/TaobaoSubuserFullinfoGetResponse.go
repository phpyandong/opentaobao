package subuser

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取指定账户子账号的详细信息 APIResponse
taobao.subuser.fullinfo.get

获取指定账户子账号的详细信息，其中包括子账号的账号信息以及员工、部门、职务信息（只能通过主账号登陆并查询属于该主账号下的某个子账号详细信息）
*/
type TaobaoSubuserFullinfoGetAPIResponse struct {
    model.CommonResponse
    TaobaoSubuserFullinfoGetResponse
}

type TaobaoSubuserFullinfoGetResponse struct {
    XMLName xml.Name `xml:"subuser_fullinfo_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 子账号详细信息，其中包括账号基本信息、员工信息和部门职务信息
    
    SubFullinfo   *SubUserFullInfo `json:"sub_fullinfo,omitempty" xml:"sub_fullinfo,omitempty"`

    
}
