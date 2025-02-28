package xhotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线下信用住取消结账专用接口 APIRequest
taobao.xhotel.order.offline.settle.cancel

线下信用住取消结账专用接口
*/
type TaobaoXhotelOrderOfflineSettleCancelRequest struct {
    model.Params

    // 阿里旅行订单号，淘宝订单号或外部订单号二选一必填
    tid   int64 

    // 取消结账的原因
    reason   string 

    // 外部订单号，和tid二选一必填（建议都写入）
    outId   string 

    // 暂时无意义，无需传入
    notifyUrl   string 

    // 请求流水号
    outUuid   string 

}

func NewTaobaoXhotelOrderOfflineSettleCancelRequest() *TaobaoXhotelOrderOfflineSettleCancelRequest{
    return &TaobaoXhotelOrderOfflineSettleCancelRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelOrderOfflineSettleCancelRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.offline.settle.cancel"
}

func (r TaobaoXhotelOrderOfflineSettleCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelOrderOfflineSettleCancelRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoXhotelOrderOfflineSettleCancelRequest) GetTid() int64 {
    return r.tid
}

func (r *TaobaoXhotelOrderOfflineSettleCancelRequest) SetReason(reason string) error {
    r.reason = reason
    r.Set("reason", reason)
    return nil
}

func (r TaobaoXhotelOrderOfflineSettleCancelRequest) GetReason() string {
    return r.reason
}

func (r *TaobaoXhotelOrderOfflineSettleCancelRequest) SetOutId(outId string) error {
    r.outId = outId
    r.Set("out_id", outId)
    return nil
}

func (r TaobaoXhotelOrderOfflineSettleCancelRequest) GetOutId() string {
    return r.outId
}

func (r *TaobaoXhotelOrderOfflineSettleCancelRequest) SetNotifyUrl(notifyUrl string) error {
    r.notifyUrl = notifyUrl
    r.Set("notify_url", notifyUrl)
    return nil
}

func (r TaobaoXhotelOrderOfflineSettleCancelRequest) GetNotifyUrl() string {
    return r.notifyUrl
}

func (r *TaobaoXhotelOrderOfflineSettleCancelRequest) SetOutUuid(outUuid string) error {
    r.outUuid = outUuid
    r.Set("out_uuid", outUuid)
    return nil
}

func (r TaobaoXhotelOrderOfflineSettleCancelRequest) GetOutUuid() string {
    return r.outUuid
}

