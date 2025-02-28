package subuser

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取指定卖家的角色列表 APIResponse
taobao.sellercenter.roles.get

获取指定卖家的角色列表，只能获取属于登陆者自己的信息。
*/
type TaobaoSellercenterRolesGetAPIResponse struct {
    model.CommonResponse
    TaobaoSellercenterRolesGetResponse
}

type TaobaoSellercenterRolesGetResponse struct {
    XMLName xml.Name `xml:"sellercenter_roles_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 卖家子账号角色列表。<br/>返回对象为 role数据对象中的role_id，role_name，description，seller_id，create_time，modified_time。不包含permissions(权限点)
    
    Roles   []Role `json:"roles,omitempty" xml:"roles>role,omitempty"`
    
    
}
