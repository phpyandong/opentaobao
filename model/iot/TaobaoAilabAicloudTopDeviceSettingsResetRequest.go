package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
重置设备个性化设置 APIRequest
taobao.ailab.aicloud.top.device.settings.reset

重置设备个性化设置
*/
type TaobaoAilabAicloudTopDeviceSettingsResetRequest struct {
    model.Params

    // 账户体系隔离
    schema   string 

    // 用户ID，此处传入第三方账户体系的用户id
    userId   string 

    // 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
    utdId   string 

    // 扩展信息，用于存放APP类型等
    ext   string 

    // 设备id
    param1   string 

}

func NewTaobaoAilabAicloudTopDeviceSettingsResetRequest() *TaobaoAilabAicloudTopDeviceSettingsResetRequest{
    return &TaobaoAilabAicloudTopDeviceSettingsResetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopDeviceSettingsResetRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.settings.reset"
}

func (r TaobaoAilabAicloudTopDeviceSettingsResetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopDeviceSettingsResetRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceSettingsResetRequest) GetSchema() string {
    return r.schema
}

func (r *TaobaoAilabAicloudTopDeviceSettingsResetRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceSettingsResetRequest) GetUserId() string {
    return r.userId
}

func (r *TaobaoAilabAicloudTopDeviceSettingsResetRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceSettingsResetRequest) GetUtdId() string {
    return r.utdId
}

func (r *TaobaoAilabAicloudTopDeviceSettingsResetRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceSettingsResetRequest) GetExt() string {
    return r.ext
}

func (r *TaobaoAilabAicloudTopDeviceSettingsResetRequest) SetParam1(param1 string) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceSettingsResetRequest) GetParam1() string {
    return r.param1
}

