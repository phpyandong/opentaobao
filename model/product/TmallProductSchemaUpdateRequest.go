package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
产品更新接口 APIRequest
tmall.product.schema.update

产品更新接口
*/
type TmallProductSchemaUpdateRequest struct {
    model.Params

    // 根据tmall.product.update.schema.get生成的产品更新规则入参数据
    xmlData   string 

    // 产品编号
    productId   int64 

}

func NewTmallProductSchemaUpdateRequest() *TmallProductSchemaUpdateRequest{
    return &TmallProductSchemaUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TmallProductSchemaUpdateRequest) GetApiMethodName() string {
    return "tmall.product.schema.update"
}

func (r TmallProductSchemaUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallProductSchemaUpdateRequest) SetXmlData(xmlData string) error {
    r.xmlData = xmlData
    r.Set("xml_data", xmlData)
    return nil
}

func (r TmallProductSchemaUpdateRequest) GetXmlData() string {
    return r.xmlData
}

func (r *TmallProductSchemaUpdateRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r TmallProductSchemaUpdateRequest) GetProductId() int64 {
    return r.productId
}

