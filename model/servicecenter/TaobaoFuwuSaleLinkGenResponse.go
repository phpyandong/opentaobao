package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务平台营销链接生成接口 APIResponse
taobao.fuwu.sale.link.gen

服务商通过使用该接口来产生营销链接，通过把这种链接发送给商家来做自定义人群的服务营销<br><br/>注：session是param_str这个参数串创建者生成的session，这个创建者与入参中的nick是不一致的。例如：A开发者创建了一个param_str的字符串，要为B商家生成一个营销链接，session必须是A开发者创建的session。
*/
type TaobaoFuwuSaleLinkGenAPIResponse struct {
    model.CommonResponse
    TaobaoFuwuSaleLinkGenResponse
}

type TaobaoFuwuSaleLinkGenResponse struct {
    XMLName xml.Name `xml:"fuwu_sale_link_gen_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 通过营销链接接口生成的营销链接短地址
    
    Url   string `json:"url,omitempty" xml:"url,omitempty"`

    
}
