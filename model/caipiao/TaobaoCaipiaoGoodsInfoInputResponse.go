package caipiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
录入参加送彩票商品信息 APIResponse
taobao.caipiao.goods.info.input

录入参加送彩票商品信息，如果录入成功，返回true，如果录入失败，返回false，后端会根据商品id与赠送类型（goodsid_presenttype_uk）来决定是新增数据还是修改数据。
*/
type TaobaoCaipiaoGoodsInfoInputAPIResponse struct {
    model.CommonResponse
    TaobaoCaipiaoGoodsInfoInputResponse
}

type TaobaoCaipiaoGoodsInfoInputResponse struct {
    XMLName xml.Name `xml:"caipiao_goods_info_input_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 录入操作是否成功
    
    InputResult   bool `json:"input_result,omitempty" xml:"input_result,omitempty"`

    
}
