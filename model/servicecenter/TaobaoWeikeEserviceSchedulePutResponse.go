package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
提交客服排班信息 APIResponse
taobao.weike.eservice.schedule.put

添加、更新、删除排班信息
*/
type TaobaoWeikeEserviceSchedulePutAPIResponse struct {
    model.CommonResponse
    TaobaoWeikeEserviceSchedulePutResponse
}

type TaobaoWeikeEserviceSchedulePutResponse struct {
    XMLName xml.Name `xml:"weike_eservice_schedule_put_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否执行成功
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
