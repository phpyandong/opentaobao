package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
营销详情添加 APIResponse
taobao.ump.detail.list.add

批量添加营销活动。替代单条添加营销详情的的API。此接口适用针对某个营销活动的多档设置，会按顺序插入detail。若在整个事务过程中出现断点，会将已插入完成的detail_id返回，注意记录这些id，并将其删除，会对交易过程中的优惠产生影响。
*/
type TaobaoUmpDetailListAddAPIResponse struct {
    model.CommonResponse
    TaobaoUmpDetailListAddResponse
}

type TaobaoUmpDetailListAddResponse struct {
    XMLName xml.Name `xml:"ump_detail_list_add_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回对应的营销详情的id列表！若有某一条插入失败，会将插入成功的detail_id放到errorMessage里面返回，此时errorMessage里面会包含格式为(id1,id2,id3)的插入成功id列表。这些ids会对交易产生影响，通过截取此信息，将对应detail删除！
    
    DetailIdList   []int64 `json:"detail_id_list,omitempty" xml:"detail_id_list>int64,omitempty"`
    
    
}
