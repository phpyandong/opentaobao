package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据交易单号判断是否为菜鸟发货订单 APIResponse
cainiao.cboss.workplatform.logistics.iscainiaoorder

根据交易单号判断是否为菜鸟发货订单
*/
type CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIResponse struct {
    model.CommonResponse
    CainiaoCbossWorkplatformLogisticsIscainiaoorderResponse
}

type CainiaoCbossWorkplatformLogisticsIscainiaoorderResponse struct {
    XMLName xml.Name `xml:"cainiao_cboss_workplatform_logistics_iscainiaoorder_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // isCainiaoOrder
    
    IsCainiaoOrder   bool `json:"is_cainiao_order,omitempty" xml:"is_cainiao_order,omitempty"`

    
    // success
    
    ResSuccess   bool `json:"res_success,omitempty" xml:"res_success,omitempty"`

    
    // errorCode
    
    ResErrorCode   string `json:"res_error_code,omitempty" xml:"res_error_code,omitempty"`

    
    // errorMsg
    
    ResErrorMsg   string `json:"res_error_msg,omitempty" xml:"res_error_msg,omitempty"`

    
}
