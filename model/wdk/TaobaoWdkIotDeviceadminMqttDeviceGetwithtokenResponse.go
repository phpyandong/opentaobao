package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取mqtt设备信息 APIResponse
taobao.wdk.iot.deviceadmin.mqtt.device.getwithtoken

智能硬件设备动态注册和获取mqtt设备信息
*/
type TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenAPIResponse struct {
    model.CommonResponse
    TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenResponse
}

type TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenResponse struct {
    XMLName xml.Name `xml:"wdk_iot_deviceadmin_mqtt_device_getwithtoken_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *HmResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
