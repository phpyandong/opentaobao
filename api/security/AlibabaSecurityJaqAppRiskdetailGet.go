package security

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/security"
)

/* 
应用风险详细信息查询接口 
alibaba.security.jaq.app.riskdetail.get

用户通过alibaba.security.jaq.app.risk.scan接口提交应用进行风险扫描后，用此接口获取风险详细信息,包含漏洞列表、恶意代码列表、仿冒应用列表等信息
*/
func AlibabaSecurityJaqAppRiskdetailGet(clt *core.SDKClient, req *security.AlibabaSecurityJaqAppRiskdetailGetRequest, session string) (*security.AlibabaSecurityJaqAppRiskdetailGetAPIResponse, error) {
    var resp security.AlibabaSecurityJaqAppRiskdetailGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
