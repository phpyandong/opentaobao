package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
专辑详情 APIRequest
alibaba.ailabs.aligenie.albums.get

给予厂商查询专辑下的音频详情
*/
type AlibabaAilabsAligenieAlbumsGetRequest struct {
    model.Params

    // 账户体系隔离
    schema   string 

    // 用户ID，此处传入第三方账户体系的用户id
    userId   string 

    // 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
    utdId   string 

    // 扩展信息，用于存放APP类型等
    ext   string 

    // 专辑 id
    param1   int64 

    // 当前页（从1开始, 目前由于底层引擎限制，实际最多能查5000个，超过5000返回为空，请确保页码*分页大小不超过5000）
    param2   int64 

    // 每页数量（不超过50）
    param3   int64 

}

func NewAlibabaAilabsAligenieAlbumsGetRequest() *AlibabaAilabsAligenieAlbumsGetRequest{
    return &AlibabaAilabsAligenieAlbumsGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabsAligenieAlbumsGetRequest) GetApiMethodName() string {
    return "alibaba.ailabs.aligenie.albums.get"
}

func (r AlibabaAilabsAligenieAlbumsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabsAligenieAlbumsGetRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

func (r AlibabaAilabsAligenieAlbumsGetRequest) GetSchema() string {
    return r.schema
}

func (r *AlibabaAilabsAligenieAlbumsGetRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaAilabsAligenieAlbumsGetRequest) GetUserId() string {
    return r.userId
}

func (r *AlibabaAilabsAligenieAlbumsGetRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

func (r AlibabaAilabsAligenieAlbumsGetRequest) GetUtdId() string {
    return r.utdId
}

func (r *AlibabaAilabsAligenieAlbumsGetRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

func (r AlibabaAilabsAligenieAlbumsGetRequest) GetExt() string {
    return r.ext
}

func (r *AlibabaAilabsAligenieAlbumsGetRequest) SetParam1(param1 int64) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r AlibabaAilabsAligenieAlbumsGetRequest) GetParam1() int64 {
    return r.param1
}

func (r *AlibabaAilabsAligenieAlbumsGetRequest) SetParam2(param2 int64) error {
    r.param2 = param2
    r.Set("param2", param2)
    return nil
}

func (r AlibabaAilabsAligenieAlbumsGetRequest) GetParam2() int64 {
    return r.param2
}

func (r *AlibabaAilabsAligenieAlbumsGetRequest) SetParam3(param3 int64) error {
    r.param3 = param3
    r.Set("param3", param3)
    return nil
}

func (r AlibabaAilabsAligenieAlbumsGetRequest) GetParam3() int64 {
    return r.param3
}

