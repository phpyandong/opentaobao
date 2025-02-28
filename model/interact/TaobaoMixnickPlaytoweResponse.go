package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
互动mixNick转微淘 APIResponse
taobao.mixnick.playtowe

微淘应用的混淆nick转为互动类型混淆nick
*/
type TaobaoMixnickPlaytoweAPIResponse struct {
    model.CommonResponse
    TaobaoMixnickPlaytoweResponse
}

type TaobaoMixnickPlaytoweResponse struct {
    XMLName xml.Name `xml:"mixnick_playtowe_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 微淘混淆nick
    
    WeMixnick   string `json:"we_mixnick,omitempty" xml:"we_mixnick,omitempty"`

    
}
