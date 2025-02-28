package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
消息号下行回复接口 APIResponse
taobao.messageaccount.messsage.reply

外部 isv 调用该进口来进行消息号消息的回复
*/
type TaobaoMessageaccountMesssageReplyAPIResponse struct {
    model.CommonResponse
    TaobaoMessageaccountMesssageReplyResponse
}

type TaobaoMessageaccountMesssageReplyResponse struct {
    XMLName xml.Name `xml:"messageaccount_messsage_reply_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TaobaoMessageaccountMesssageReplyResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
