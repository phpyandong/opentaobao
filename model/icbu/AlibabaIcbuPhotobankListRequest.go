package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
国际站图片银行查询接口 APIRequest
alibaba.icbu.photobank.list

国际站图片银行查询接口
*/
type AlibabaIcbuPhotobankListRequest struct {
    model.Params

    // 当前翻页数
    currentPage   int64 

    // 图片组id
    groupId   string 

    // 存放位置 必要条件, 包括ALL_GROUP(所有目录), SUB_GROUP(指定图片组下),UNGROUP(未分组), TEMP(disable)四个值
    locationType   string 

    // 每页显示数
    pageSize   int64 

    // 额外的上下文信息. 例如:icvId
    extraContext   string 

}

func NewAlibabaIcbuPhotobankListRequest() *AlibabaIcbuPhotobankListRequest{
    return &AlibabaIcbuPhotobankListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIcbuPhotobankListRequest) GetApiMethodName() string {
    return "alibaba.icbu.photobank.list"
}

func (r AlibabaIcbuPhotobankListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIcbuPhotobankListRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

func (r AlibabaIcbuPhotobankListRequest) GetCurrentPage() int64 {
    return r.currentPage
}

func (r *AlibabaIcbuPhotobankListRequest) SetGroupId(groupId string) error {
    r.groupId = groupId
    r.Set("group_id", groupId)
    return nil
}

func (r AlibabaIcbuPhotobankListRequest) GetGroupId() string {
    return r.groupId
}

func (r *AlibabaIcbuPhotobankListRequest) SetLocationType(locationType string) error {
    r.locationType = locationType
    r.Set("location_type", locationType)
    return nil
}

func (r AlibabaIcbuPhotobankListRequest) GetLocationType() string {
    return r.locationType
}

func (r *AlibabaIcbuPhotobankListRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaIcbuPhotobankListRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AlibabaIcbuPhotobankListRequest) SetExtraContext(extraContext string) error {
    r.extraContext = extraContext
    r.Set("extra_context", extraContext)
    return nil
}

func (r AlibabaIcbuPhotobankListRequest) GetExtraContext() string {
    return r.extraContext
}

