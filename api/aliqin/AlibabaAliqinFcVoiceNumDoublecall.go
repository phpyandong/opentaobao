package aliqin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aliqin"
)

/* 
多方通话 
alibaba.aliqin.fc.voice.num.doublecall

根据传入的主叫号码与被叫号码，由系统依次向主叫号码与被叫号码发起呼叫，双方都应答后，开始一对一通话并开始计费。使用前需要在阿里大于管理中心添加呼叫双方的显示号码。
*/
func AlibabaAliqinFcVoiceNumDoublecall(clt *core.SDKClient, req *aliqin.AlibabaAliqinFcVoiceNumDoublecallRequest, session string) (*aliqin.AlibabaAliqinFcVoiceNumDoublecallAPIResponse, error) {
    var resp aliqin.AlibabaAliqinFcVoiceNumDoublecallAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
