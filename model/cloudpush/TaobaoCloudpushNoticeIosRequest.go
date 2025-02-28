package cloudpush

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推送通知给ios设备 APIRequest
taobao.cloudpush.notice.ios

推送通知给ios设备
*/
type TaobaoCloudpushNoticeIosRequest struct {
    model.Params

    // 通知摘要
    summary   string 

    // 推送目标: device:推送给设备; account:推送给指定帐号,all: 推送给全部
    target   string 

    // 根据Target来设定，如Target=device, 则对应的值为 设备id1,设备id2. 多个值使用逗号分隔
    targetValue   string 

    // iOS的通知是通过APNS中心来发送的，需要填写对应的环境信息.  DEV:表示开发环境, PRODUCT: 表示生产环境.
    env   string 

    // 提供给IOS通知的扩展属性，如角标或者声音等,注意：参数值为json
    ext   string 

}

func NewTaobaoCloudpushNoticeIosRequest() *TaobaoCloudpushNoticeIosRequest{
    return &TaobaoCloudpushNoticeIosRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoCloudpushNoticeIosRequest) GetApiMethodName() string {
    return "taobao.cloudpush.notice.ios"
}

func (r TaobaoCloudpushNoticeIosRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoCloudpushNoticeIosRequest) SetSummary(summary string) error {
    r.summary = summary
    r.Set("summary", summary)
    return nil
}

func (r TaobaoCloudpushNoticeIosRequest) GetSummary() string {
    return r.summary
}

func (r *TaobaoCloudpushNoticeIosRequest) SetTarget(target string) error {
    r.target = target
    r.Set("target", target)
    return nil
}

func (r TaobaoCloudpushNoticeIosRequest) GetTarget() string {
    return r.target
}

func (r *TaobaoCloudpushNoticeIosRequest) SetTargetValue(targetValue string) error {
    r.targetValue = targetValue
    r.Set("target_value", targetValue)
    return nil
}

func (r TaobaoCloudpushNoticeIosRequest) GetTargetValue() string {
    return r.targetValue
}

func (r *TaobaoCloudpushNoticeIosRequest) SetEnv(env string) error {
    r.env = env
    r.Set("env", env)
    return nil
}

func (r TaobaoCloudpushNoticeIosRequest) GetEnv() string {
    return r.env
}

func (r *TaobaoCloudpushNoticeIosRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

func (r TaobaoCloudpushNoticeIosRequest) GetExt() string {
    return r.ext
}

