package train

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商获取订单退票信息 APIResponse
taobao.train.agent.get.refund

代理商获取订单信息回调API
*/
type TaobaoTrainAgentGetRefundAPIResponse struct {
    model.CommonResponse
    TaobaoTrainAgentGetRefundResponse
}

type TaobaoTrainAgentGetRefundResponse struct {
    XMLName xml.Name `xml:"train_agent_get_refund_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 系统自动生成
    
    TopRefundApplyList   string `json:"top_refund_apply_list,omitempty" xml:"top_refund_apply_list,omitempty"`

    
}
