package tbk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-推广者-权益物料精选 APIResponse
taobao.tbk.dg.optimus.promotion

推广者使用。支持入参推广者对应的“推广位”和官方提供的“权益物料id”，获取指定权益物料。
*/
type TaobaoTbkDgOptimusPromotionAPIResponse struct {
    model.CommonResponse
    TaobaoTbkDgOptimusPromotionResponse
}

type TaobaoTbkDgOptimusPromotionResponse struct {
    XMLName xml.Name `xml:"tbk_dg_optimus_promotion_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // resultList
    
    ResultList   []TaobaoTbkDgOptimusPromotionMapData `json:"result_list,omitempty" xml:"result_list>taobao_tbk_dg_optimus_promotion_map_data,omitempty"`
    
    
}
