package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询流量充值状态 APIResponse
alibaba.aliqin.flow.wallet.query.charge

查询流量充值状态
*/
type AlibabaAliqinFlowWalletQueryChargeAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinFlowWalletQueryChargeResponse
}

type AlibabaAliqinFlowWalletQueryChargeResponse struct {
    XMLName xml.Name `xml:"alibaba_aliqin_flow_wallet_query_charge_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 充值状态
    
    Charge   string `json:"charge,omitempty" xml:"charge,omitempty"`

    
}
