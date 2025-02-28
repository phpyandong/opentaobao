package tmc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发布单条消息 APIRequest
taobao.tmc.message.produce

发布单条消息
*/
type TaobaoTmcMessageProduceRequest struct {
    model.Params

    // 消息内容的JSON表述，必须按照topic的定义来填充
    content   string 

    // 消息类型
    topic   string 

    // 回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。
    mediaContent   []*model.File 

    // 回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。
    mediaContent2   []*model.File 

    // 回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。
    mediaContent3   []*model.File 

    // 回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。
    mediaContent4   []*model.File 

    // 回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。
    mediaContent5   []*model.File 

    // 目标分组，一般为default
    targetGroup   string 

}

func NewTaobaoTmcMessageProduceRequest() *TaobaoTmcMessageProduceRequest{
    return &TaobaoTmcMessageProduceRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTmcMessageProduceRequest) GetApiMethodName() string {
    return "taobao.tmc.message.produce"
}

func (r TaobaoTmcMessageProduceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTmcMessageProduceRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

func (r TaobaoTmcMessageProduceRequest) GetContent() string {
    return r.content
}

func (r *TaobaoTmcMessageProduceRequest) SetTopic(topic string) error {
    r.topic = topic
    r.Set("topic", topic)
    return nil
}

func (r TaobaoTmcMessageProduceRequest) GetTopic() string {
    return r.topic
}

func (r *TaobaoTmcMessageProduceRequest) SetMediaContent(mediaContent []*model.File) error {
    r.mediaContent = mediaContent
    r.Set("media_content", mediaContent)
    return nil
}

func (r TaobaoTmcMessageProduceRequest) GetMediaContent() []*model.File {
    return r.mediaContent
}

func (r *TaobaoTmcMessageProduceRequest) SetMediaContent2(mediaContent2 []*model.File) error {
    r.mediaContent2 = mediaContent2
    r.Set("media_content2", mediaContent2)
    return nil
}

func (r TaobaoTmcMessageProduceRequest) GetMediaContent2() []*model.File {
    return r.mediaContent2
}

func (r *TaobaoTmcMessageProduceRequest) SetMediaContent3(mediaContent3 []*model.File) error {
    r.mediaContent3 = mediaContent3
    r.Set("media_content3", mediaContent3)
    return nil
}

func (r TaobaoTmcMessageProduceRequest) GetMediaContent3() []*model.File {
    return r.mediaContent3
}

func (r *TaobaoTmcMessageProduceRequest) SetMediaContent4(mediaContent4 []*model.File) error {
    r.mediaContent4 = mediaContent4
    r.Set("media_content4", mediaContent4)
    return nil
}

func (r TaobaoTmcMessageProduceRequest) GetMediaContent4() []*model.File {
    return r.mediaContent4
}

func (r *TaobaoTmcMessageProduceRequest) SetMediaContent5(mediaContent5 []*model.File) error {
    r.mediaContent5 = mediaContent5
    r.Set("media_content5", mediaContent5)
    return nil
}

func (r TaobaoTmcMessageProduceRequest) GetMediaContent5() []*model.File {
    return r.mediaContent5
}

func (r *TaobaoTmcMessageProduceRequest) SetTargetGroup(targetGroup string) error {
    r.targetGroup = targetGroup
    r.Set("target_group", targetGroup)
    return nil
}

func (r TaobaoTmcMessageProduceRequest) GetTargetGroup() string {
    return r.targetGroup
}

