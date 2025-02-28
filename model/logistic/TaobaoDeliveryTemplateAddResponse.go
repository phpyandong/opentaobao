package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新增运费模板 APIResponse
taobao.delivery.template.add

增加运费模板的外部接口
*/
type TaobaoDeliveryTemplateAddAPIResponse struct {
    model.CommonResponse
    TaobaoDeliveryTemplateAddResponse
}

type TaobaoDeliveryTemplateAddResponse struct {
    XMLName xml.Name `xml:"delivery_template_add_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 模板对象
    
    DeliveryTemplate   *DeliveryTemplate `json:"delivery_template,omitempty" xml:"delivery_template,omitempty"`

    
}
