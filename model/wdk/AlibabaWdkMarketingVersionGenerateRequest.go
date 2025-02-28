package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
生成发布使用的版本号 APIRequest
alibaba.wdk.marketing.version.generate

生成发布使用的版本号
*/
type AlibabaWdkMarketingVersionGenerateRequest struct {
    model.Params

    // 档期版本号参数信息
    param   *SeasonVersionParam 

}

func NewAlibabaWdkMarketingVersionGenerateRequest() *AlibabaWdkMarketingVersionGenerateRequest{
    return &AlibabaWdkMarketingVersionGenerateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingVersionGenerateRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.version.generate"
}

func (r AlibabaWdkMarketingVersionGenerateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingVersionGenerateRequest) SetParam(param *SeasonVersionParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaWdkMarketingVersionGenerateRequest) GetParam() *SeasonVersionParam {
    return r.param
}

