package user

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/user"
)

/* 
open account token验证 
taobao.open.account.token.validate

open account token验证
*/
func TaobaoOpenAccountTokenValidate(clt *core.SDKClient, req *user.TaobaoOpenAccountTokenValidateRequest, session string) (*user.TaobaoOpenAccountTokenValidateAPIResponse, error) {
    var resp user.TaobaoOpenAccountTokenValidateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
