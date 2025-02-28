package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
设备安装活动激活 APIResponse
alibaba.baichuan.aso.activate

设备安装活动激活
*/
type AlibabaBaichuanAsoActivateAPIResponse struct {
    model.CommonResponse
    AlibabaBaichuanAsoActivateResponse
}

type AlibabaBaichuanAsoActivateResponse struct {
    XMLName xml.Name `xml:"alibaba_baichuan_aso_activate_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AsoActivateDeviceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
