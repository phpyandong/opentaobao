package waybill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取商家单一自定义区 APIResponse
cainiao.cloudprint.single.customarea.get

商家所有快递公司模板只有一个自定义区
*/
type CainiaoCloudprintSingleCustomareaGetAPIResponse struct {
    model.CommonResponse
    Response *CainiaoCloudprintSingleCustomareaGetResponse `json:"cainiao_cloudprint_single_customarea_get_response,omitempty"`
}

type CainiaoCloudprintSingleCustomareaGetResponse struct {

    // result
    Result   *CloudPrintBaseResult `json:"result,omitempty"`

}
