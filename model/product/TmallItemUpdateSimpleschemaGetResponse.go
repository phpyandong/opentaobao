package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
官网同购编辑商品的get接口 APIResponse
tmall.item.update.simpleschema.get

官网同购编辑商品的get接口
*/
type TmallItemUpdateSimpleschemaGetAPIResponse struct {
    model.CommonResponse
    TmallItemUpdateSimpleschemaGetResponse
}

type TmallItemUpdateSimpleschemaGetResponse struct {
    XMLName xml.Name `xml:"tmall_item_update_simpleschema_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Error   bool `json:"error,omitempty" xml:"error,omitempty"`

    
    // 错误信息
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`

    
    // 商品信息
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
