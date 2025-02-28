package crm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询卖家的分组 APIResponse
taobao.crm.groups.get

查询卖家的分组，返回查询到的分组列表，分页返回分组
*/
type TaobaoCrmGroupsGetAPIResponse struct {
    model.CommonResponse
    TaobaoCrmGroupsGetResponse
}

type TaobaoCrmGroupsGetResponse struct {
    XMLName xml.Name `xml:"crm_groups_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 查询到的当前卖家的当前页的会员
    
    Groups   []Group `json:"groups,omitempty" xml:"groups>group,omitempty"`
    
    
    // 记录总数
    
    TotalResult   int64 `json:"total_result,omitempty" xml:"total_result,omitempty"`

    
}
