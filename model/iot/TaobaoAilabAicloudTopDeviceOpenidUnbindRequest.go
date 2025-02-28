package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
openTaoBaoId解绑设备 APIRequest
taobao.ailab.aicloud.top.device.openid.unbind

openTaoBaoId解绑设备
*/
type TaobaoAilabAicloudTopDeviceOpenidUnbindRequest struct {
    model.Params

    // 淘宝openId
    openId   string 

    // 设备uuid
    uuid   string 

    // 账户体系隔离
    schema   string 

    // 用户ID，此处传入第三方账户体系的用户id
    userId   string 

    // 扩展信息，用于存放APP类型等
    ext   string 

    // 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
    utdId   string 

}

func NewTaobaoAilabAicloudTopDeviceOpenidUnbindRequest() *TaobaoAilabAicloudTopDeviceOpenidUnbindRequest{
    return &TaobaoAilabAicloudTopDeviceOpenidUnbindRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopDeviceOpenidUnbindRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.openid.unbind"
}

func (r TaobaoAilabAicloudTopDeviceOpenidUnbindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopDeviceOpenidUnbindRequest) SetOpenId(openId string) error {
    r.openId = openId
    r.Set("open_id", openId)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceOpenidUnbindRequest) GetOpenId() string {
    return r.openId
}

func (r *TaobaoAilabAicloudTopDeviceOpenidUnbindRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceOpenidUnbindRequest) GetUuid() string {
    return r.uuid
}

func (r *TaobaoAilabAicloudTopDeviceOpenidUnbindRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceOpenidUnbindRequest) GetSchema() string {
    return r.schema
}

func (r *TaobaoAilabAicloudTopDeviceOpenidUnbindRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceOpenidUnbindRequest) GetUserId() string {
    return r.userId
}

func (r *TaobaoAilabAicloudTopDeviceOpenidUnbindRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceOpenidUnbindRequest) GetExt() string {
    return r.ext
}

func (r *TaobaoAilabAicloudTopDeviceOpenidUnbindRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceOpenidUnbindRequest) GetUtdId() string {
    return r.utdId
}

