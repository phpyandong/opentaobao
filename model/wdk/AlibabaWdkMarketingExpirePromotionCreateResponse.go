package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
短保优惠创建 APIResponse
alibaba.wdk.marketing.expire.promotion.create

过期优惠优惠信息录入
*/
type AlibabaWdkMarketingExpirePromotionCreateAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingExpirePromotionCreateResponse
}

type AlibabaWdkMarketingExpirePromotionCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_expire_promotion_create_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // message
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // data
    
    Datas   []T `json:"datas,omitempty" xml:"datas>t,omitempty"`
    
    
    // errorCode
    
    FailCode   string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`

    
    // success
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
