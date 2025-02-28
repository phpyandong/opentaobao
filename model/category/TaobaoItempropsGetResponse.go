package category

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取标准商品类目属性 APIResponse
taobao.itemprops.get

通过设置必要的参数，来获取商品后台标准类目属性，以及这些属性里面详细的属性值prop_values。
*/
type TaobaoItempropsGetAPIResponse struct {
    model.CommonResponse
    TaobaoItempropsGetResponse
}

type TaobaoItempropsGetResponse struct {
    XMLName xml.Name `xml:"itemprops_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 最近修改时间(只有取全量或增量的时候会返回该字段)。格式:yyyy-MM-dd HH:mm:ss
    
    LastModified   string `json:"last_modified,omitempty" xml:"last_modified,omitempty"`

    
    // 类目属性信息(如果是取全量或者增量，不包括属性值),根据fields传入的参数返回相应的结果
    
    ItemProps   []ItemProp `json:"item_props,omitempty" xml:"item_props>item_prop,omitempty"`
    
    
}
