package aliyun

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据渠道id和状态列出appkey APIResponse
account.aliyuncs.com.ListAppkeyByOwnerAndBid.2013-07-01

根据渠道id和状态列出appkey
*/
type AccountAliyuncsComListAppkeyByOwnerAndBid2013-07-01APIResponse struct {
    model.CommonResponse
    AccountAliyuncsComListAppkeyByOwnerAndBid2013-07-01Response
}

type AccountAliyuncsComListAppkeyByOwnerAndBid2013-07-01Response struct {
    XMLName xml.Name `xml:"account_aliyuncs_com_ListAppkeyByOwnerAndBid_2013-07-01_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // return result
    
    ResultData   string `json:"result_data,omitempty" xml:"result_data,omitempty"`

    
}
