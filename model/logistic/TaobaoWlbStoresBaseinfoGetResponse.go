package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家按照仓的类型获取仓库接口 APIResponse
taobao.wlb.stores.baseinfo.get

通过USERID和仓库类型，获取商家自有仓库或菜鸟仓库或全部仓库
*/
type TaobaoWlbStoresBaseinfoGetAPIResponse struct {
    model.CommonResponse
    TaobaoWlbStoresBaseinfoGetResponse
}

type TaobaoWlbStoresBaseinfoGetResponse struct {
    XMLName xml.Name `xml:"wlb_stores_baseinfo_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 仓库列表
    
    StoreInfoList   []StoreInfo `json:"store_info_list,omitempty" xml:"store_info_list>store_info,omitempty"`
    
    
    // 返回的总数
    
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`

    
}
