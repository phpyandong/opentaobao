package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）商品发布增量更新接口 APIRequest
alibaba.icbu.product.schema.update

商品更新接口，方式为增量更新，只更新传入的字段
*/
type AlibabaIcbuProductSchemaUpdateRequest struct {
    model.Params

    // 发布入参
    paramProductTopPublishRequest   *ProductTopPublishRequest 

}

func NewAlibabaIcbuProductSchemaUpdateRequest() *AlibabaIcbuProductSchemaUpdateRequest{
    return &AlibabaIcbuProductSchemaUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIcbuProductSchemaUpdateRequest) GetApiMethodName() string {
    return "alibaba.icbu.product.schema.update"
}

func (r AlibabaIcbuProductSchemaUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIcbuProductSchemaUpdateRequest) SetParamProductTopPublishRequest(paramProductTopPublishRequest *ProductTopPublishRequest) error {
    r.paramProductTopPublishRequest = paramProductTopPublishRequest
    r.Set("param_product_top_publish_request", paramProductTopPublishRequest)
    return nil
}

func (r AlibabaIcbuProductSchemaUpdateRequest) GetParamProductTopPublishRequest() *ProductTopPublishRequest {
    return r.paramProductTopPublishRequest
}

