package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量获取商品详细信息 APIRequest
taobao.items.seller.list.get

批量获取商品详细信息
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong>
*/
type TaobaoItemsSellerListGetRequest struct {
    model.Params

    // 需要返回的商品字段列表。可选值：点击返回结果中的Item结构体中能展示出来的所有字段，多个字段用“,”分隔。注：返回所有sku信息的字段名称是sku而不是skus。
    fields   string 

    // 商品ID列表，多个ID用半角逗号隔开，一次最多不超过20个。注：获取不存在的商品ID或获取别人的商品都不会报错，但没有商品数据返回。
    numIids   []string 

}

func NewTaobaoItemsSellerListGetRequest() *TaobaoItemsSellerListGetRequest{
    return &TaobaoItemsSellerListGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoItemsSellerListGetRequest) GetApiMethodName() string {
    return "taobao.items.seller.list.get"
}

func (r TaobaoItemsSellerListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoItemsSellerListGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoItemsSellerListGetRequest) GetFields() string {
    return r.fields
}

func (r *TaobaoItemsSellerListGetRequest) SetNumIids(numIids []string) error {
    r.numIids = numIids
    r.Set("num_iids", numIids)
    return nil
}

func (r TaobaoItemsSellerListGetRequest) GetNumIids() []string {
    return r.numIids
}

