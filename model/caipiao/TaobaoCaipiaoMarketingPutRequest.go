package caipiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建或修改商家送彩票活动 APIRequest
taobao.caipiao.marketing.put

卖家通过此接口新增或修改送彩票活动的配置，比如活动时间、活动的条件等。

店铺装修、宝贝详情页的宣导素材示例：
https://gw.alicdn.com/tfs/TB1_nOiSXXXXXbgXXXXXXXXXXXX-750-280.png
https://gw.alicdn.com/tfs/TB1FZX6SXXXXXXzXFXXXXXXXXXX-790-280.png
https://gw.alicdn.com/tfs/TB1z4t8SXXXXXckXpXXXXXXXXXX-750-280.png
https://gw.alicdn.com/tfs/TB1BhqgSXXXXXcDXXXXXXXXXXXX-750-280.png
https://gw.alicdn.com/tfs/TB1TYt9SXXXXXXAXFXXXXXXXXXX-750-280.png
https://gw.alicdn.com/tfs/TB1tzpNSXXXXXacXVXXXXXXXXXX-790-280.png
https://gw.alicdn.com/tfs/TB1UXdxSXXXXXXsapXXXXXXXXXX-790-280.png
https://gw.alicdn.com/tfs/TB1_gV.SXXXXXbZXpXXXXXXXXXX-790-280.png
*/
type TaobaoCaipiaoMarketingPutRequest struct {
    model.Params

    // 活动详情设置
    detail   *WangcaiMarketingDetail 

}

func NewTaobaoCaipiaoMarketingPutRequest() *TaobaoCaipiaoMarketingPutRequest{
    return &TaobaoCaipiaoMarketingPutRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoCaipiaoMarketingPutRequest) GetApiMethodName() string {
    return "taobao.caipiao.marketing.put"
}

func (r TaobaoCaipiaoMarketingPutRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoCaipiaoMarketingPutRequest) SetDetail(detail *WangcaiMarketingDetail) error {
    r.detail = detail
    r.Set("detail", detail)
    return nil
}

func (r TaobaoCaipiaoMarketingPutRequest) GetDetail() *WangcaiMarketingDetail {
    return r.detail
}

