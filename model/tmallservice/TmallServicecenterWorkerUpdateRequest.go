package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改工人信息 APIRequest
tmall.servicecenter.worker.update

修改工人信息。该接口为多个业务公用，部分字段可忽略。对于电器预约安装业务，同一个服务商，通过工人姓名+手机号+biz_type 保证唯一性。工人已存在才可以修改。
错误码如下
100000, 系统错误
100001, 工人信息校验失败
100002, 用户校验失败
100003, 操作失败
10004,工人信息为空
10005,服务商id为空或者服务商名称为空
10006, 工人不存在
10007, 工人已存在
10008, 缺少工人姓名
10009, 缺少工人电话
10010, 网点不存在
11000, category_id 无效
11001, biz_type 无效
20001,已查询到最后一页
*/
type TmallServicecenterWorkerUpdateRequest struct {
    model.Params

    // 工人信息
    worker   *WorkerDto 

}

func NewTmallServicecenterWorkerUpdateRequest() *TmallServicecenterWorkerUpdateRequest{
    return &TmallServicecenterWorkerUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterWorkerUpdateRequest) GetApiMethodName() string {
    return "tmall.servicecenter.worker.update"
}

func (r TmallServicecenterWorkerUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterWorkerUpdateRequest) SetWorker(worker *WorkerDto) error {
    r.worker = worker
    r.Set("worker", worker)
    return nil
}

func (r TmallServicecenterWorkerUpdateRequest) GetWorker() *WorkerDto {
    return r.worker
}

