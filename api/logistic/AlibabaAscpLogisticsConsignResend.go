package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
修改物流公司和运单号 
alibaba.ascp.logistics.consign.resend

支持卖家发货后修改运单号;支持在线下单和自己联系两种发货方式;使用条件：
1、必须是已发货订单，自己联系发货的必须24小时内才可修改；在线下单的，必须下单后物流公司未揽收成功前才可修改；
2、自己联系只能切换为自己联系的公司，在线下单也只能切换为在线下单的物流公司
*/
func AlibabaAscpLogisticsConsignResend(clt *core.SDKClient, req *logistic.AlibabaAscpLogisticsConsignResendRequest, session string) (*logistic.AlibabaAscpLogisticsConsignResendAPIResponse, error) {
    var resp logistic.AlibabaAscpLogisticsConsignResendAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
