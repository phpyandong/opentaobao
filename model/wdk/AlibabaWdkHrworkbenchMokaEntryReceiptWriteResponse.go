package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
摩卡确认入职后往入职单据表写数据接口 APIResponse
alibaba.wdk.hrworkbench.moka.entry.receipt.write

摩卡确认入职后往入职单据表写数据接口
*/
type AlibabaWdkHrworkbenchMokaEntryReceiptWriteAPIResponse struct {
    model.CommonResponse
    AlibabaWdkHrworkbenchMokaEntryReceiptWriteResponse
}

type AlibabaWdkHrworkbenchMokaEntryReceiptWriteResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_hrworkbench_moka_entry_receipt_write_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // data
    
    Data   bool `json:"data,omitempty" xml:"data,omitempty"`

    
    // 失败
    
    Fail   bool `json:"fail,omitempty" xml:"fail,omitempty"`

    
    // 结果信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 成功并且结果非空
    
    SuccAndNotNull   bool `json:"succ_and_not_null,omitempty" xml:"succ_and_not_null,omitempty"`

    
    // 成功结果为空
    
    SuccAndNull   bool `json:"succ_and_null,omitempty" xml:"succ_and_null,omitempty"`

    
    // trace_id
    
    TraceId   string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`

    
}
