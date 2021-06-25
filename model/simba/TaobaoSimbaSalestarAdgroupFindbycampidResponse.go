package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
(销量明星)批量获取推广计划下的推广组信息 APIResponse
taobao.simba.salestar.adgroup.findbycampid

批量得到推广计划下的推广组
*/
type TaobaoSimbaSalestarAdgroupFindbycampidAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaSalestarAdgroupFindbycampidResponse `json:"taobao_simba_salestar_adgroup_findbycampid_response,omitempty"`
}

type TaobaoSimbaSalestarAdgroupFindbycampidResponse struct {

    // 返回的推广组分页对象
    Adgroups   *ADGroupPage `json:"adgroups,omitempty"`

}
