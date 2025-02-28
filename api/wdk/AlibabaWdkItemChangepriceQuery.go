package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
按照价格变更时间段，查询会变更价格的单据的商品 
alibaba.wdk.item.changeprice.query

/**
     * 按照价格变更时间段，查询会变更价格的单据的商品
     * 传入QueryPriceChangeTypeEnum.BASE_PRICE,返回变价时间在startTime-endTime之间的所有单据
     * 传入QueryPriceChangeTypeEnum.SKU_PROMOTION_START,
     * 返回活动开始时间在 startTime<=活动开始时间<endTime 之间的所有单品促销单据
     * 传入QueryPriceChangeTypeEnum.SKU_PROMOTION_END,
     * 返回活动结束时间在 startTime<=活动结束时间<endTime 之间的所有单品促销单据
     */
*/
func AlibabaWdkItemChangepriceQuery(clt *core.SDKClient, req *wdk.AlibabaWdkItemChangepriceQueryRequest, session string) (*wdk.AlibabaWdkItemChangepriceQueryAPIResponse, error) {
    var resp wdk.AlibabaWdkItemChangepriceQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
