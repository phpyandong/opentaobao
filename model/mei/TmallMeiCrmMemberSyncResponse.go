package mei

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
同步推送会员信息 APIResponse
tmall.mei.crm.member.sync

品牌方通过该api主动推送会员信息。使用场景包括
1.用户在线上注册后，线下补充信息后，同步到线上。
2.其他情况的主动推送变更。
*/
type TmallMeiCrmMemberSyncAPIResponse struct {
    model.CommonResponse
    TmallMeiCrmMemberSyncResponse
}

type TmallMeiCrmMemberSyncResponse struct {
    XMLName xml.Name `xml:"tmall_mei_crm_member_sync_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 处理的其他信息
    
    MeiExtraInfo   string `json:"mei_extra_info,omitempty" xml:"mei_extra_info,omitempty"`

    
}
