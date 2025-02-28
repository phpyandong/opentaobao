package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新增门店扩展属性 APIResponse
taobao.place.store.extend.add

新增授权用户的门店扩展属性
*/
type TaobaoPlaceStoreExtendAddAPIResponse struct {
    model.CommonResponse
    TaobaoPlaceStoreExtendAddResponse
}

type TaobaoPlaceStoreExtendAddResponse struct {
    XMLName xml.Name `xml:"place_store_extend_add_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否失败
    
    Failure   bool `json:"failure,omitempty" xml:"failure,omitempty"`

    
    // 返回结果：true成功；false失败
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
    // 个数
    
    TotalNum   int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`

    
}
