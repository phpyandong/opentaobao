package legalcase

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
诉讼文件上传接口 APIResponse
alibaba.legal.suit.file.upload

上传文件接口
*/
type AlibabaLegalSuitFileUploadAPIResponse struct {
    model.CommonResponse
    AlibabaLegalSuitFileUploadResponse
}

type AlibabaLegalSuitFileUploadResponse struct {
    XMLName xml.Name `xml:"alibaba_legal_suit_file_upload_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 失败原因
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`

    
    // 是否调用成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 文件上传后的生成的id
    
    Content   string `json:"content,omitempty" xml:"content,omitempty"`

    
    // 失败的code枚举
    
    CodeError   string `json:"code_error,omitempty" xml:"code_error,omitempty"`

    
}
