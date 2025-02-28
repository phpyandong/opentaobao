package category

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取标准类目属性值 APIResponse
taobao.itempropvalues.get

获取标准类目属性值
*/
type TaobaoItempropvaluesGetAPIResponse struct {
    model.CommonResponse
    TaobaoItempropvaluesGetResponse
}

type TaobaoItempropvaluesGetResponse struct {
    XMLName xml.Name `xml:"itempropvalues_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 最近修改时间。格式:yyyy-MM-dd HH:mm:ss
    
    LastModified   string `json:"last_modified,omitempty" xml:"last_modified,omitempty"`

    
    // 属性值,根据fields传入的参数返回相应的结果
    
    PropValues   []PropValue `json:"prop_values,omitempty" xml:"prop_values>prop_value,omitempty"`
    
    
}
