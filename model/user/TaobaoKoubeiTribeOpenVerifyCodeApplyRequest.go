package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
口碑综合体手机号获取验证码 APIRequest
taobao.koubei.tribe.open.verify.code.apply

口碑综合体通过手机号获取验证码对外开放接口
*/
type TaobaoKoubeiTribeOpenVerifyCodeApplyRequest struct {
    model.Params

    // 数据集id
    dataSetId   string 

    // 手机号
    phone   string 

}

func NewTaobaoKoubeiTribeOpenVerifyCodeApplyRequest() *TaobaoKoubeiTribeOpenVerifyCodeApplyRequest{
    return &TaobaoKoubeiTribeOpenVerifyCodeApplyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoKoubeiTribeOpenVerifyCodeApplyRequest) GetApiMethodName() string {
    return "taobao.koubei.tribe.open.verify.code.apply"
}

func (r TaobaoKoubeiTribeOpenVerifyCodeApplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoKoubeiTribeOpenVerifyCodeApplyRequest) SetDataSetId(dataSetId string) error {
    r.dataSetId = dataSetId
    r.Set("data_set_id", dataSetId)
    return nil
}

func (r TaobaoKoubeiTribeOpenVerifyCodeApplyRequest) GetDataSetId() string {
    return r.dataSetId
}

func (r *TaobaoKoubeiTribeOpenVerifyCodeApplyRequest) SetPhone(phone string) error {
    r.phone = phone
    r.Set("phone", phone)
    return nil
}

func (r TaobaoKoubeiTribeOpenVerifyCodeApplyRequest) GetPhone() string {
    return r.phone
}

