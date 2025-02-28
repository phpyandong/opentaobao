package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
顺丰仓作业回传 APIResponse
alibaba.wdk.fulfill.sf.btoc.fms.wms.work.order.callback

顺丰仓作业单回传接口
*/
type AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIResponse struct {
    model.CommonResponse
    AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackResponse
}

type AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_fulfill_sf_btoc_fms_wms_work_order_callback_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 响应提示信息
    
    RespMessage   string `json:"resp_message,omitempty" xml:"resp_message,omitempty"`

    
    // 响应code(SUCCESS:回传成功； SYSTEM_ERROR :系统异常； PARAM_ERROR :参数错误； BUSINESS_ERROR:业务异常；)
    
    RespCode   string `json:"resp_code,omitempty" xml:"resp_code,omitempty"`

    
}
