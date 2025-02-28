package waybill

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/waybill"
)

/* 
获取商家使用的标准模板 
cainiao.cloudprint.isvtemplates.get

获取商家使用的标准模板
*/
func CainiaoCloudprintIsvtemplatesGet(clt *core.SDKClient, req *waybill.CainiaoCloudprintIsvtemplatesGetRequest, session string) (*waybill.CainiaoCloudprintIsvtemplatesGetAPIResponse, error) {
    var resp waybill.CainiaoCloudprintIsvtemplatesGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
