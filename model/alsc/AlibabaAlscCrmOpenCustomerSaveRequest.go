package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
保存和更新顾客 APIRequest
alibaba.alsc.crm.open.customer.save

用来保存顾客，如果已经存在的话，则更新顾客
*/
type AlibabaAlscCrmOpenCustomerSaveRequest struct {
    model.Params

    // 入参
    paramCustomerSaveOpenReq   *CustomerSaveOpenReq 

}

func NewAlibabaAlscCrmOpenCustomerSaveRequest() *AlibabaAlscCrmOpenCustomerSaveRequest{
    return &AlibabaAlscCrmOpenCustomerSaveRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmOpenCustomerSaveRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.open.customer.save"
}

func (r AlibabaAlscCrmOpenCustomerSaveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmOpenCustomerSaveRequest) SetParamCustomerSaveOpenReq(paramCustomerSaveOpenReq *CustomerSaveOpenReq) error {
    r.paramCustomerSaveOpenReq = paramCustomerSaveOpenReq
    r.Set("param_customer_save_open_req", paramCustomerSaveOpenReq)
    return nil
}

func (r AlibabaAlscCrmOpenCustomerSaveRequest) GetParamCustomerSaveOpenReq() *CustomerSaveOpenReq {
    return r.paramCustomerSaveOpenReq
}

