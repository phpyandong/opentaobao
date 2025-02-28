package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
儿童音频列表 APIRequest
taobao.ailab.aicloud.top.freelisten.childrenalbum

儿童音频列表
*/
type TaobaoAilabAicloudTopFreelistenChildrenalbumRequest struct {
    model.Params

    // 账户体系隔离
    schema   string 

    // 用户ID，此处传入第三方账户体系的用户id
    userId   string 

    // 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
    utdId   string 

    // 扩展信息，用于存放APP类型等
    ext   string 

    // 音频类型，目前只支持以下几种类型：英语儿歌 英语故事 双语故事 国学启蒙 古典名著 成语故事 寓言故事 神话故事 诗词朗读 诗词婉唱 谚语故事 胎教音乐 经典儿歌 摇篮曲 睡前故事 绘本故事 儿童故事 儿童百科 经典故事 公主故事 名人故事 胎教故事
    param1   string 

    // 页数
    param2   int64 

    // 每页条目数
    param3   int64 

}

func NewTaobaoAilabAicloudTopFreelistenChildrenalbumRequest() *TaobaoAilabAicloudTopFreelistenChildrenalbumRequest{
    return &TaobaoAilabAicloudTopFreelistenChildrenalbumRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopFreelistenChildrenalbumRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.freelisten.childrenalbum"
}

func (r TaobaoAilabAicloudTopFreelistenChildrenalbumRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopFreelistenChildrenalbumRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

func (r TaobaoAilabAicloudTopFreelistenChildrenalbumRequest) GetSchema() string {
    return r.schema
}

func (r *TaobaoAilabAicloudTopFreelistenChildrenalbumRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r TaobaoAilabAicloudTopFreelistenChildrenalbumRequest) GetUserId() string {
    return r.userId
}

func (r *TaobaoAilabAicloudTopFreelistenChildrenalbumRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

func (r TaobaoAilabAicloudTopFreelistenChildrenalbumRequest) GetUtdId() string {
    return r.utdId
}

func (r *TaobaoAilabAicloudTopFreelistenChildrenalbumRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

func (r TaobaoAilabAicloudTopFreelistenChildrenalbumRequest) GetExt() string {
    return r.ext
}

func (r *TaobaoAilabAicloudTopFreelistenChildrenalbumRequest) SetParam1(param1 string) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r TaobaoAilabAicloudTopFreelistenChildrenalbumRequest) GetParam1() string {
    return r.param1
}

func (r *TaobaoAilabAicloudTopFreelistenChildrenalbumRequest) SetParam2(param2 int64) error {
    r.param2 = param2
    r.Set("param2", param2)
    return nil
}

func (r TaobaoAilabAicloudTopFreelistenChildrenalbumRequest) GetParam2() int64 {
    return r.param2
}

func (r *TaobaoAilabAicloudTopFreelistenChildrenalbumRequest) SetParam3(param3 int64) error {
    r.param3 = param3
    r.Set("param3", param3)
    return nil
}

func (r TaobaoAilabAicloudTopFreelistenChildrenalbumRequest) GetParam3() int64 {
    return r.param3
}

