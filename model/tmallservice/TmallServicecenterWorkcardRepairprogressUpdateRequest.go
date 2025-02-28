package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新维修进度 APIRequest
tmall.servicecenter.workcard.repairprogress.update

提供给外部合作服务商的维修进度更改接口
*/
type TmallServicecenterWorkcardRepairprogressUpdateRequest struct {
    model.Params

    // 图片列表
    picUrlList   []string 

    // 请求节点的动作描述，唯一标识一个节点
    action   string 

    // 工单id
    workcardId   int64 

    // 真实接单服务商账号Nick
    realTpNick   string 

    // 服务目标物瑕疵信息
    targetGoodsDefects   string 

    // 衣服，鞋子
    receivedGoods   string 

}

func NewTmallServicecenterWorkcardRepairprogressUpdateRequest() *TmallServicecenterWorkcardRepairprogressUpdateRequest{
    return &TmallServicecenterWorkcardRepairprogressUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterWorkcardRepairprogressUpdateRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.repairprogress.update"
}

func (r TmallServicecenterWorkcardRepairprogressUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterWorkcardRepairprogressUpdateRequest) SetPicUrlList(picUrlList []string) error {
    r.picUrlList = picUrlList
    r.Set("pic_url_list", picUrlList)
    return nil
}

func (r TmallServicecenterWorkcardRepairprogressUpdateRequest) GetPicUrlList() []string {
    return r.picUrlList
}

func (r *TmallServicecenterWorkcardRepairprogressUpdateRequest) SetAction(action string) error {
    r.action = action
    r.Set("action", action)
    return nil
}

func (r TmallServicecenterWorkcardRepairprogressUpdateRequest) GetAction() string {
    return r.action
}

func (r *TmallServicecenterWorkcardRepairprogressUpdateRequest) SetWorkcardId(workcardId int64) error {
    r.workcardId = workcardId
    r.Set("workcard_id", workcardId)
    return nil
}

func (r TmallServicecenterWorkcardRepairprogressUpdateRequest) GetWorkcardId() int64 {
    return r.workcardId
}

func (r *TmallServicecenterWorkcardRepairprogressUpdateRequest) SetRealTpNick(realTpNick string) error {
    r.realTpNick = realTpNick
    r.Set("real_tp_nick", realTpNick)
    return nil
}

func (r TmallServicecenterWorkcardRepairprogressUpdateRequest) GetRealTpNick() string {
    return r.realTpNick
}

func (r *TmallServicecenterWorkcardRepairprogressUpdateRequest) SetTargetGoodsDefects(targetGoodsDefects string) error {
    r.targetGoodsDefects = targetGoodsDefects
    r.Set("target_goods_defects", targetGoodsDefects)
    return nil
}

func (r TmallServicecenterWorkcardRepairprogressUpdateRequest) GetTargetGoodsDefects() string {
    return r.targetGoodsDefects
}

func (r *TmallServicecenterWorkcardRepairprogressUpdateRequest) SetReceivedGoods(receivedGoods string) error {
    r.receivedGoods = receivedGoods
    r.Set("received_goods", receivedGoods)
    return nil
}

func (r TmallServicecenterWorkcardRepairprogressUpdateRequest) GetReceivedGoods() string {
    return r.receivedGoods
}

