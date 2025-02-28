package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
本地生活风控数据接入 APIResponse
alibaba.alsc.kms.access

第三方使用本地生活数据对外提供服务，上报访问日志信息接口
*/
type AlibabaAlscKmsAccessAPIResponse struct {
    model.CommonResponse
    AlibabaAlscKmsAccessResponse
}

type AlibabaAlscKmsAccessResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_kms_access_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // code
    
    Resultcode   string `json:"resultcode,omitempty" xml:"resultcode,omitempty"`

    
    // 是否成功
    
    Resultsuccess   bool `json:"resultsuccess,omitempty" xml:"resultsuccess,omitempty"`

    
    // message
    
    Resultmessage   string `json:"resultmessage,omitempty" xml:"resultmessage,omitempty"`

    
}
