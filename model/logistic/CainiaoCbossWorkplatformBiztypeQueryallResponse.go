package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟工单平台根据交易订单查询某条业务线上的所有业务类型 APIResponse
cainiao.cboss.workplatform.biztype.queryall

菜鸟工单平台根据交易订单查询某条业务线上的所有业务类型。 目前调用者ISV
*/
type CainiaoCbossWorkplatformBiztypeQueryallAPIResponse struct {
    model.CommonResponse
    CainiaoCbossWorkplatformBiztypeQueryallResponse
}

type CainiaoCbossWorkplatformBiztypeQueryallResponse struct {
    XMLName xml.Name `xml:"cainiao_cboss_workplatform_biztype_queryall_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // bizTypeJson
    
    BizTypeJson   string `json:"biz_type_json,omitempty" xml:"biz_type_json,omitempty"`

    
    // success
    
    ResSuccess   bool `json:"res_success,omitempty" xml:"res_success,omitempty"`

    
    // errorCode
    
    ResErrorCode   string `json:"res_error_code,omitempty" xml:"res_error_code,omitempty"`

    
    // errorMsg
    
    ResErrorMsg   string `json:"res_error_msg,omitempty" xml:"res_error_msg,omitempty"`

    
}
