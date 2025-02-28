package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品关联绑定接口 APIResponse
taobao.qimen.itemstore.banding

商家在ERP等系统中调用该接口，将线上商品和线下门店“新建/删除”关联。这里的线上。每次只能单个商品关联多个门店，门店上限200
*/
type TaobaoQimenItemstoreBandingAPIResponse struct {
    model.CommonResponse
    TaobaoQimenItemstoreBandingResponse
}

type TaobaoQimenItemstoreBandingResponse struct {
    XMLName xml.Name `xml:"qimen_itemstore_banding_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 响应信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 响应描述
    
    Flag   string `json:"flag,omitempty" xml:"flag,omitempty"`

    
    // 响应编码
    
    QimenCode   string `json:"qimen_code,omitempty" xml:"qimen_code,omitempty"`

    
}
