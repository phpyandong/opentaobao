package waybill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
电子面单云打印更新接口 APIResponse
cainiao.waybill.ii.update

商家更新电子面单号对应的面单信息。
*/
type CainiaoWaybillIiUpdateAPIResponse struct {
    model.CommonResponse
    CainiaoWaybillIiUpdateResponse
}

type CainiaoWaybillIiUpdateResponse struct {
    XMLName xml.Name `xml:"cainiao_waybill_ii_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 面单号
    
    WaybillCode   string `json:"waybill_code,omitempty" xml:"waybill_code,omitempty"`

    
    // 模板内容
    
    PrintData   string `json:"print_data,omitempty" xml:"print_data,omitempty"`

    
}
