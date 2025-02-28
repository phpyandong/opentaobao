package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
图片银行分组操作接口 APIRequest
alibaba.icbu.photobank.group.operate

修改用户图片银行的分组信息，包括 新增分组，删除分组，分组重命名
*/
type AlibabaIcbuPhotobankGroupOperateRequest struct {
    model.Params

    // 图片分组操作请求对象
    photoGroupOperationRequest   *PhotoGroupOperationRequest 

}

func NewAlibabaIcbuPhotobankGroupOperateRequest() *AlibabaIcbuPhotobankGroupOperateRequest{
    return &AlibabaIcbuPhotobankGroupOperateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIcbuPhotobankGroupOperateRequest) GetApiMethodName() string {
    return "alibaba.icbu.photobank.group.operate"
}

func (r AlibabaIcbuPhotobankGroupOperateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIcbuPhotobankGroupOperateRequest) SetPhotoGroupOperationRequest(photoGroupOperationRequest *PhotoGroupOperationRequest) error {
    r.photoGroupOperationRequest = photoGroupOperationRequest
    r.Set("photo_group_operation_request", photoGroupOperationRequest)
    return nil
}

func (r AlibabaIcbuPhotobankGroupOperateRequest) GetPhotoGroupOperationRequest() *PhotoGroupOperationRequest {
    return r.photoGroupOperationRequest
}

