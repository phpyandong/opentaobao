package icbu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
（新）获取商品信息 APIResponse
alibaba.icbu.product.schema.render

获取ICBU商品发布的字段填写规则和单个商品对应填写数据，适用于单个商品编辑场景，不包括草稿。
*/
type AlibabaIcbuProductSchemaRenderAPIResponse struct {
    model.CommonResponse
    Response *AlibabaIcbuProductSchemaRenderResponse `json:"alibaba_icbu_product_schema_render_response,omitempty"`
}

type AlibabaIcbuProductSchemaRenderResponse struct {

    // 商品发布规则和对应填写数据
    Data   string `json:"data,omitempty"`

    // 错误信息，数组形式的字符串，用;分割，支持中英繁，按照传入的语种参数决定
    Message   string `json:"message,omitempty"`

    // 返回的错误码，数组形式的字符串，用;分割
    MsgCode   string `json:"msg_code,omitempty"`

    // 请求是否成功
    BizSuccess   bool `json:"biz_success,omitempty"`

    // 错误追踪码，请务必打印在日志中，后续排查问题请提交此错误追踪码
    TraceId   string `json:"trace_id,omitempty"`

}
