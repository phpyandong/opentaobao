package tbk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-公用-淘口令生成 APIResponse
taobao.tbk.tpwd.create

提供淘口令生成接口。提交需要生成淘口令的内容、logo、url等参数，生成淘口令，其中关键信息为￥SADadW￥，后续可基于淘口令进行文案包装组装用于传播。
*/
type TaobaoTbkTpwdCreateAPIResponse struct {
    model.CommonResponse
    TaobaoTbkTpwdCreateResponse
}

type TaobaoTbkTpwdCreateResponse struct {
    XMLName xml.Name `xml:"tbk_tpwd_create_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果对象
    
    Data   *TaobaoTbkTpwdCreateMapData `json:"data,omitempty" xml:"data,omitempty"`

    
}
