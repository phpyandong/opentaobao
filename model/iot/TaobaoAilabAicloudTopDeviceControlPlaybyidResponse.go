package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
通过id播放歌曲 APIResponse
taobao.ailab.aicloud.top.device.control.playbyid

通过id播放歌曲
*/
type TaobaoAilabAicloudTopDeviceControlPlaybyidAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAilabAicloudTopDeviceControlPlaybyidResponse `json:"taobao_ailab_aicloud_top_device_control_playbyid_response,omitempty"`
}

type TaobaoAilabAicloudTopDeviceControlPlaybyidResponse struct {

    // 返回结果
    Result   *AiCloudResult `json:"result,omitempty"`

}
