package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
添加单个物流宝商品 APIResponse
taobao.wlb.item.add

添加物流宝商品，支持物流宝子商品和属性添加
*/
type TaobaoWlbItemAddAPIResponse struct {
    model.CommonResponse
    TaobaoWlbItemAddResponse
}

type TaobaoWlbItemAddResponse struct {
    XMLName xml.Name `xml:"wlb_item_add_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 新增的商品
    
    ItemId   string `json:"item_id,omitempty" xml:"item_id,omitempty"`

    
}
