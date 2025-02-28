package caipiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据卖家id与appkey获取商品信息 APIResponse
taobao.caipiao.goods.info.get

根据卖家id与appkey获取商品信息。
*/
type TaobaoCaipiaoGoodsInfoGetAPIResponse struct {
    model.CommonResponse
    TaobaoCaipiaoGoodsInfoGetResponse
}

type TaobaoCaipiaoGoodsInfoGetResponse struct {
    XMLName xml.Name `xml:"caipiao_goods_info_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回列表的大小
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`

    
    // 查询的结果列表
    
    Results   []LotteryWangcaiSellerGoodsInfo `json:"results,omitempty" xml:"results>lottery_wangcai_seller_goods_info,omitempty"`
    
    
}
