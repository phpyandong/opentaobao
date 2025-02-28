package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
无线开放视频内容安全检查 APIRequest
taobao.wireless.video.check

无线开放内容检查，提供涉黄暴力政治音视频的异步检查。更详情介绍见 <a href="https://help.aliyun.com/document_detail/70436.html" target="blank">阿里云内容安全</a>

此API会进行三个场景的审核，检测不通过的视频将被隐藏，用户无法访问被隐藏的视频。

目前，该接口仅支持顽兔空间的视频扫描。
*/
type TaobaoWirelessVideoCheckRequest struct {
    model.Params

    // 视频的URL，必须为淘系安全域名地址。视频格式支持flv、mp4。
    url   string 

    // 视频截帧间隔，取值范围为[1, 60]，单位为秒。 截帧最多张数为200张，请根据视频时长，合理设置截帧间隔。
    interval   int64 

}

func NewTaobaoWirelessVideoCheckRequest() *TaobaoWirelessVideoCheckRequest{
    return &TaobaoWirelessVideoCheckRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWirelessVideoCheckRequest) GetApiMethodName() string {
    return "taobao.wireless.video.check"
}

func (r TaobaoWirelessVideoCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWirelessVideoCheckRequest) SetUrl(url string) error {
    r.url = url
    r.Set("url", url)
    return nil
}

func (r TaobaoWirelessVideoCheckRequest) GetUrl() string {
    return r.url
}

func (r *TaobaoWirelessVideoCheckRequest) SetInterval(interval int64) error {
    r.interval = interval
    r.Set("interval", interval)
    return nil
}

func (r TaobaoWirelessVideoCheckRequest) GetInterval() int64 {
    return r.interval
}

