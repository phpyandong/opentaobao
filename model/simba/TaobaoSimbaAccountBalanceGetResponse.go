package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取实时余额，”元”为单位 APIResponse
taobao.simba.account.balance.get

获取实时余额，”元”为单位
*/
type TaobaoSimbaAccountBalanceGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaAccountBalanceGetResponse
}

type TaobaoSimbaAccountBalanceGetResponse struct {
    XMLName xml.Name `xml:"simba_account_balance_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 实时余额，”元”为单位
    
    Balance   string `json:"balance,omitempty" xml:"balance,omitempty"`

    
}
