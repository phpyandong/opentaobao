package tanx

import (
    "github.com/bububa/opentaobao/model"
)

/* 
批量获取交易列表 APIResponse
taobao.tanx.deals.get

批量获取交易信息
*/
type TaobaoTanxDealsGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTanxDealsGetResponse `json:"taobao_tanx_deals_get_response,omitempty"`
}

type TaobaoTanxDealsGetResponse struct {

    // 查询是否成功
    Sucess   bool `json:"sucess,omitempty"`

    // 查询结果编码
    Code   int64 `json:"code,omitempty"`

    // 查询结果信息
    Message   string `json:"message,omitempty"`

    // 查询结果
    Deals   []DealInfoDTO `json:"deals,omitempty"`

}
