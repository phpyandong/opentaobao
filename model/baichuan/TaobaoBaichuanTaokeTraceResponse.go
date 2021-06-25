package baichuan

import (
    "github.com/bububa/opentaobao/model"
)

/* 
百川淘客打点 APIResponse
taobao.baichuan.taoke.trace

百川淘客打点
*/
type TaobaoBaichuanTaokeTraceAPIResponse struct {
    model.CommonResponse
    Response *TaobaoBaichuanTaokeTraceResponse `json:"taobao_baichuan_taoke_trace_response,omitempty"`
}

type TaobaoBaichuanTaokeTraceResponse struct {

    // name
    Name   string `json:"name,omitempty"`

}
