package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV报名官方活动(中心化流量) APIResponse
alibaba.interact.activity.apply

支持商家将使用isv创建的活动所对应的权益信息同步到手淘，供过滤是否在中心化流量入口透出
*/
type AlibabaInteractActivityApplyAPIResponse struct {
    model.CommonResponse
    AlibabaInteractActivityApplyResponse
}

type AlibabaInteractActivityApplyResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_activity_apply_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 服务结果对象
    
    Data   *ActivityWriteResult `json:"data,omitempty" xml:"data,omitempty"`

    
    // top接口执行成功与否
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 出错提示信息
    
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`

    
}
