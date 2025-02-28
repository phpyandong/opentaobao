package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商获取物料申请单列表 APIResponse
alibaba.cuntao.interact.requisition.get

供应商获取物料申请单列表
*/
type AlibabaCuntaoInteractRequisitionGetAPIResponse struct {
    model.CommonResponse
    AlibabaCuntaoInteractRequisitionGetResponse
}

type AlibabaCuntaoInteractRequisitionGetResponse struct {
    XMLName xml.Name `xml:"alibaba_cuntao_interact_requisition_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlibabaCuntaoInteractRequisitionGetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
