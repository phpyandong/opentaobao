package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新旧属性的映射 APIRequest
alibaba.icbu.category.id.mapping

商品发布接口升级，需要传入新的类目。这个接口 根据旧的类目id，获取新的类目id
*/
type AlibabaIcbuCategoryIdMappingRequest struct {
    model.Params

    // 发布类目id
    catId   int64 

    // 属性值id
    attributeValueId   int64 

    // 属性id
    attributeId   int64 

    // 转化类型, 1 = 转化类目id 2= 转化属性id 3= 转化属性值id
    convertType   int64 

}

func NewAlibabaIcbuCategoryIdMappingRequest() *AlibabaIcbuCategoryIdMappingRequest{
    return &AlibabaIcbuCategoryIdMappingRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIcbuCategoryIdMappingRequest) GetApiMethodName() string {
    return "alibaba.icbu.category.id.mapping"
}

func (r AlibabaIcbuCategoryIdMappingRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIcbuCategoryIdMappingRequest) SetCatId(catId int64) error {
    r.catId = catId
    r.Set("cat_id", catId)
    return nil
}

func (r AlibabaIcbuCategoryIdMappingRequest) GetCatId() int64 {
    return r.catId
}

func (r *AlibabaIcbuCategoryIdMappingRequest) SetAttributeValueId(attributeValueId int64) error {
    r.attributeValueId = attributeValueId
    r.Set("attribute_value_id", attributeValueId)
    return nil
}

func (r AlibabaIcbuCategoryIdMappingRequest) GetAttributeValueId() int64 {
    return r.attributeValueId
}

func (r *AlibabaIcbuCategoryIdMappingRequest) SetAttributeId(attributeId int64) error {
    r.attributeId = attributeId
    r.Set("attribute_id", attributeId)
    return nil
}

func (r AlibabaIcbuCategoryIdMappingRequest) GetAttributeId() int64 {
    return r.attributeId
}

func (r *AlibabaIcbuCategoryIdMappingRequest) SetConvertType(convertType int64) error {
    r.convertType = convertType
    r.Set("convert_type", convertType)
    return nil
}

func (r AlibabaIcbuCategoryIdMappingRequest) GetConvertType() int64 {
    return r.convertType
}

