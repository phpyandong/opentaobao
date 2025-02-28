package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改物流公司和运单号 APIRequest
alibaba.ascp.logistics.consign.resend

支持卖家发货后修改运单号;支持在线下单和自己联系两种发货方式;使用条件：
1、必须是已发货订单，自己联系发货的必须24小时内才可修改；在线下单的，必须下单后物流公司未揽收成功前才可修改；
2、自己联系只能切换为自己联系的公司，在线下单也只能切换为在线下单的物流公司
*/
type AlibabaAscpLogisticsConsignResendRequest struct {
    model.Params

    // 订单id
    tid   string 

    // 拆单子订单列表，对应的数据是：子订单号列表。可以不传，但是如果传了则必须符合传递的规则。子订单必须是操作的物流订单的子订单的真子集
    subTids   string 

    // 包裹包含的运单号及快递公司编号,多包裹时，需要包含所有包裹的运单号等信息
    consignPkgs   []TopConsignPkgRequest 

}

func NewAlibabaAscpLogisticsConsignResendRequest() *AlibabaAscpLogisticsConsignResendRequest{
    return &AlibabaAscpLogisticsConsignResendRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpLogisticsConsignResendRequest) GetApiMethodName() string {
    return "alibaba.ascp.logistics.consign.resend"
}

func (r AlibabaAscpLogisticsConsignResendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpLogisticsConsignResendRequest) SetTid(tid string) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r AlibabaAscpLogisticsConsignResendRequest) GetTid() string {
    return r.tid
}

func (r *AlibabaAscpLogisticsConsignResendRequest) SetSubTids(subTids string) error {
    r.subTids = subTids
    r.Set("sub_tids", subTids)
    return nil
}

func (r AlibabaAscpLogisticsConsignResendRequest) GetSubTids() string {
    return r.subTids
}

func (r *AlibabaAscpLogisticsConsignResendRequest) SetConsignPkgs(consignPkgs []TopConsignPkgRequest) error {
    r.consignPkgs = consignPkgs
    r.Set("consign_pkgs", consignPkgs)
    return nil
}

func (r AlibabaAscpLogisticsConsignResendRequest) GetConsignPkgs() []TopConsignPkgRequest {
    return r.consignPkgs
}

