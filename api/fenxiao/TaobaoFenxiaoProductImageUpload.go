package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
产品图片上传 
taobao.fenxiao.product.image.upload

产品主图图片空间相对路径或绝对路径添加或更新，或者是图片上传。如果指定位置的图片已存在，则覆盖原有信息。如果位置为1,自动设为主图；如果位置为0，表示属性图片
*/
func TaobaoFenxiaoProductImageUpload(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductImageUploadRequest, session string) (*fenxiao.TaobaoFenxiaoProductImageUploadAPIResponse, error) {
    var resp fenxiao.TaobaoFenxiaoProductImageUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
