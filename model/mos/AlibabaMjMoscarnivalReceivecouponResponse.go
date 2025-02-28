package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据手机号码领券 APIResponse
alibaba.mj.moscarnival.receivecoupon

根据手机号码领券
*/
type AlibabaMjMoscarnivalReceivecouponAPIResponse struct {
    model.CommonResponse
    AlibabaMjMoscarnivalReceivecouponResponse
}

type AlibabaMjMoscarnivalReceivecouponResponse struct {
    XMLName xml.Name `xml:"alibaba_mj_moscarnival_receivecoupon_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Result   *AlibabaMjMoscarnivalReceivecouponResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
