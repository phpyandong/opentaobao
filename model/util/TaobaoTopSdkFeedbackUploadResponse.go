package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
sdk信息回调 APIResponse
taobao.top.sdk.feedback.upload

sdk回调客户端基本信息到开放平台，用于做监控之类，有助于帮助isv监控系统稳定性
*/
type TaobaoTopSdkFeedbackUploadAPIResponse struct {
    model.CommonResponse
    TaobaoTopSdkFeedbackUploadResponse
}

type TaobaoTopSdkFeedbackUploadResponse struct {
    XMLName xml.Name `xml:"top_sdk_feedback_upload_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 控制回传间隔（单位：秒）
    
    UploadInterval   int64 `json:"upload_interval,omitempty" xml:"upload_interval,omitempty"`

    
}
