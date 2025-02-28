package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备授权码 APIRequest
taobao.ailab.aicloud.top.device.authcode.get

获取设备授权码
*/
type TaobaoAilabAicloudTopDeviceAuthcodeGetRequest struct {
    model.Params

    // 账户体系隔离，即硬件接入平台中取得的schema key。
    schema   string 

    // 用户ID，此处传入第三方账户体系的用户id，由开发者或厂商自行定义，每一schema key下唯一即可
    userId   string 

    // (废弃) 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
    utdId   string 

    // 扩展信息，用于存放APP类型等
    ext   string 

}

func NewTaobaoAilabAicloudTopDeviceAuthcodeGetRequest() *TaobaoAilabAicloudTopDeviceAuthcodeGetRequest{
    return &TaobaoAilabAicloudTopDeviceAuthcodeGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopDeviceAuthcodeGetRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.authcode.get"
}

func (r TaobaoAilabAicloudTopDeviceAuthcodeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopDeviceAuthcodeGetRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceAuthcodeGetRequest) GetSchema() string {
    return r.schema
}

func (r *TaobaoAilabAicloudTopDeviceAuthcodeGetRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceAuthcodeGetRequest) GetUserId() string {
    return r.userId
}

func (r *TaobaoAilabAicloudTopDeviceAuthcodeGetRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceAuthcodeGetRequest) GetUtdId() string {
    return r.utdId
}

func (r *TaobaoAilabAicloudTopDeviceAuthcodeGetRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceAuthcodeGetRequest) GetExt() string {
    return r.ext
}

