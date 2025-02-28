package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
翱象商家自有渠道 逆向单申请 APIRequest
alibaba.tcls.aelophy.merchant.channel.refund.apply

翱象小程序 用户逆向单申请
*/
type AlibabaTclsAelophyMerchantChannelRefundApplyRequest struct {
    model.Params

    // 请求对象
    refundApplyInfo   *RefundApplyInfo 

}

func NewAlibabaTclsAelophyMerchantChannelRefundApplyRequest() *AlibabaTclsAelophyMerchantChannelRefundApplyRequest{
    return &AlibabaTclsAelophyMerchantChannelRefundApplyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTclsAelophyMerchantChannelRefundApplyRequest) GetApiMethodName() string {
    return "alibaba.tcls.aelophy.merchant.channel.refund.apply"
}

func (r AlibabaTclsAelophyMerchantChannelRefundApplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTclsAelophyMerchantChannelRefundApplyRequest) SetRefundApplyInfo(refundApplyInfo *RefundApplyInfo) error {
    r.refundApplyInfo = refundApplyInfo
    r.Set("refund_apply_info", refundApplyInfo)
    return nil
}

func (r AlibabaTclsAelophyMerchantChannelRefundApplyRequest) GetRefundApplyInfo() *RefundApplyInfo {
    return r.refundApplyInfo
}

