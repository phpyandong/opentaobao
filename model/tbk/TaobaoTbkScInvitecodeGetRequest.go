package tbk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-公用-私域用户邀请码生成 APIRequest
taobao.tbk.sc.invitecode.get

私域用户管理(即渠道管理或会员运营管理)功能中，通过此API可生成淘宝客自身的邀请码。
*/
type TaobaoTbkScInvitecodeGetRequest struct {
    model.Params

    // 渠道关系ID
    relationId   int64 

    // 渠道推广的物料类型
    relationApp   string 

    // 邀请码类型，1 - 渠道邀请，2 - 渠道裂变，3 -会员邀请
    codeType   int64 

}

func NewTaobaoTbkScInvitecodeGetRequest() *TaobaoTbkScInvitecodeGetRequest{
    return &TaobaoTbkScInvitecodeGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTbkScInvitecodeGetRequest) GetApiMethodName() string {
    return "taobao.tbk.sc.invitecode.get"
}

func (r TaobaoTbkScInvitecodeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTbkScInvitecodeGetRequest) SetRelationId(relationId int64) error {
    r.relationId = relationId
    r.Set("relation_id", relationId)
    return nil
}

func (r TaobaoTbkScInvitecodeGetRequest) GetRelationId() int64 {
    return r.relationId
}

func (r *TaobaoTbkScInvitecodeGetRequest) SetRelationApp(relationApp string) error {
    r.relationApp = relationApp
    r.Set("relation_app", relationApp)
    return nil
}

func (r TaobaoTbkScInvitecodeGetRequest) GetRelationApp() string {
    return r.relationApp
}

func (r *TaobaoTbkScInvitecodeGetRequest) SetCodeType(codeType int64) error {
    r.codeType = codeType
    r.Set("code_type", codeType)
    return nil
}

func (r TaobaoTbkScInvitecodeGetRequest) GetCodeType() int64 {
    return r.codeType
}

