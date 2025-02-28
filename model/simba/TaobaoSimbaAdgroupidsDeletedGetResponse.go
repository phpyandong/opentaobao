package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取删除的推广组ID APIResponse
taobao.simba.adgroupids.deleted.get

获取删除的推广组ID
*/
type TaobaoSimbaAdgroupidsDeletedGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaAdgroupidsDeletedGetResponse
}

type TaobaoSimbaAdgroupidsDeletedGetResponse struct {
    XMLName xml.Name `xml:"simba_adgroupids_deleted_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 推广组ID列表
    
    DeletedAdgroupIds   []int64 `json:"deleted_adgroup_ids,omitempty" xml:"deleted_adgroup_ids>int64,omitempty"`
    
    
}
