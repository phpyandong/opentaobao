package wms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品同步 APIResponse
taobao.wlb.wms.sku.create

商品同步
*/
type TaobaoWlbWmsSkuCreateAPIResponse struct {
    model.CommonResponse
    TaobaoWlbWmsSkuCreateResponse
}

type TaobaoWlbWmsSkuCreateResponse struct {
    XMLName xml.Name `xml:"wlb_wms_sku_create_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 系统自动生成
    
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`

    
    // 错误信息
    
    WlErrorMsg   string `json:"wl_error_msg,omitempty" xml:"wl_error_msg,omitempty"`

    
    // 是否成功
    
    WlSuccess   bool `json:"wl_success,omitempty" xml:"wl_success,omitempty"`

    
    // 错误码
    
    WlErrorCode   string `json:"wl_error_code,omitempty" xml:"wl_error_code,omitempty"`

    
}
