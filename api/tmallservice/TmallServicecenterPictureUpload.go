package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
上传图片 
tmall.servicecenter.picture.upload

给服务商ERP系统使用，用于上传图片保存在天猫，一般用于工单信息回传时候保存服务商的服务证明信息相关的图片。
*/
func TmallServicecenterPictureUpload(clt *core.SDKClient, req *tmallservice.TmallServicecenterPictureUploadRequest, session string) (*tmallservice.TmallServicecenterPictureUploadAPIResponse, error) {
    var resp tmallservice.TmallServicecenterPictureUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
