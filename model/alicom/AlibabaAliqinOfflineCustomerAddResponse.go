package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
系外拉新代理商增加客户经理接口 APIResponse
alibaba.aliqin.offline.customer.add

阿里通信这边维护了代理商和其对应的客户经理的关系，用于业务处理，开放该接口用于代理商将他们系统下的客户经理信息同步给我们
*/
type AlibabaAliqinOfflineCustomerAddAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinOfflineCustomerAddResponse
}

type AlibabaAliqinOfflineCustomerAddResponse struct {
    XMLName xml.Name `xml:"alibaba_aliqin_offline_customer_add_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
