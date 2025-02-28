package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-获取推广计划产品列表 APIResponse
alibaba.scbp.target.ad.plan.product.list.get

定向推广-获取推广计划产品列表
*/
type AlibabaScbpTargetAdPlanProductListGetAPIResponse struct {
    model.CommonResponse
    AlibabaScbpTargetAdPlanProductListGetResponse
}

type AlibabaScbpTargetAdPlanProductListGetResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_target_ad_plan_product_list_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // TopP4pQuickCampaignProductView
    
    ProductList   []TopP4pQuickCampaignProductView `json:"product_list,omitempty" xml:"product_list>top_p4p_quick_campaign_product_view,omitempty"`
    
    
    // 总页数
    
    TotalPage   int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`

    
    // 总个数
    
    TotalNum   int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`

    
}
