package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家代客售后提交逆向申请 APIRequest
alibaba.tcls.aelophy.refund.csapply

商家代客售后提交逆向申请
*/
type AlibabaTclsAelophyRefundCsapplyRequest struct {
    model.Params

    // 逆向申请入参
    refundCsApplyDTO   *RefundCsApplyDto 

}

func NewAlibabaTclsAelophyRefundCsapplyRequest() *AlibabaTclsAelophyRefundCsapplyRequest{
    return &AlibabaTclsAelophyRefundCsapplyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTclsAelophyRefundCsapplyRequest) GetApiMethodName() string {
    return "alibaba.tcls.aelophy.refund.csapply"
}

func (r AlibabaTclsAelophyRefundCsapplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTclsAelophyRefundCsapplyRequest) SetRefundCsApplyDTO(refundCsApplyDTO *RefundCsApplyDto) error {
    r.refundCsApplyDTO = refundCsApplyDTO
    r.Set("refund_cs_apply_d_t_o", refundCsApplyDTO)
    return nil
}

func (r AlibabaTclsAelophyRefundCsapplyRequest) GetRefundCsApplyDTO() *RefundCsApplyDto {
    return r.refundCsApplyDTO
}

