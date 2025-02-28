package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改支付密码 APIRequest
alibaba.alsc.crm.customer.updateppw

修改支付密码
*/
type AlibabaAlscCrmCustomerUpdateppwRequest struct {
    model.Params

    // 修改密码
    updatePayPasswdReq   *UpdatePayPasswdReq 

}

func NewAlibabaAlscCrmCustomerUpdateppwRequest() *AlibabaAlscCrmCustomerUpdateppwRequest{
    return &AlibabaAlscCrmCustomerUpdateppwRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmCustomerUpdateppwRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.customer.updateppw"
}

func (r AlibabaAlscCrmCustomerUpdateppwRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmCustomerUpdateppwRequest) SetUpdatePayPasswdReq(updatePayPasswdReq *UpdatePayPasswdReq) error {
    r.updatePayPasswdReq = updatePayPasswdReq
    r.Set("update_pay_passwd_req", updatePayPasswdReq)
    return nil
}

func (r AlibabaAlscCrmCustomerUpdateppwRequest) GetUpdatePayPasswdReq() *UpdatePayPasswdReq {
    return r.updatePayPasswdReq
}

