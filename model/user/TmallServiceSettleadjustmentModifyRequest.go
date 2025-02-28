package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改结算调整单 APIRequest
tmall.service.settleadjustment.modify

提供给服务商在对结算有异议时，发起结算调整单。
通过说明调整单ID，调整费用值，调整原因进行结算调整单修改。
*/
type TmallServiceSettleadjustmentModifyRequest struct {
    model.Params

    // 结算调整单父节点
    paramSettleAdjustmentRequest   *SettleAdjustmentRequest 

}

func NewTmallServiceSettleadjustmentModifyRequest() *TmallServiceSettleadjustmentModifyRequest{
    return &TmallServiceSettleadjustmentModifyRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServiceSettleadjustmentModifyRequest) GetApiMethodName() string {
    return "tmall.service.settleadjustment.modify"
}

func (r TmallServiceSettleadjustmentModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServiceSettleadjustmentModifyRequest) SetParamSettleAdjustmentRequest(paramSettleAdjustmentRequest *SettleAdjustmentRequest) error {
    r.paramSettleAdjustmentRequest = paramSettleAdjustmentRequest
    r.Set("param_settle_adjustment_request", paramSettleAdjustmentRequest)
    return nil
}

func (r TmallServiceSettleadjustmentModifyRequest) GetParamSettleAdjustmentRequest() *SettleAdjustmentRequest {
    return r.paramSettleAdjustmentRequest
}

