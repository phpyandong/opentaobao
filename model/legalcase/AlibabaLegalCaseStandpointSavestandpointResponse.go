package legalcase

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新增反馈口径 APIResponse
alibaba.legal.case.standpoint.savestandpoint

新增反馈口径 ,从外部接受反馈的口径
*/
type AlibabaLegalCaseStandpointSavestandpointAPIResponse struct {
    model.CommonResponse
    AlibabaLegalCaseStandpointSavestandpointResponse
}

type AlibabaLegalCaseStandpointSavestandpointResponse struct {
    XMLName xml.Name `xml:"alibaba_legal_case_standpoint_savestandpoint_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // success
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 错误编码
    
    Errcode   string `json:"errcode,omitempty" xml:"errcode,omitempty"`

    
    // 反馈的新增口径id
    
    Content   int64 `json:"content,omitempty" xml:"content,omitempty"`

    
    // 错误描述
    
    Errmsg   string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`

    
}
