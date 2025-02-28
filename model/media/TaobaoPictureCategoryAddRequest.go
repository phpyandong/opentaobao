package media

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增图片分类信息 APIRequest
taobao.picture.category.add

同一卖家最多添加500个图片分类，图片分类名称长度最大为20个字符
*/
type TaobaoPictureCategoryAddRequest struct {
    model.Params

    // 图片分类名称，最大长度20字符，中文字符算2个字符，不能为空
    pictureCategoryName   string 

    // 图片分类的父分类,一级分类的parent_id为0,二级分类的则为其父分类的picture_category_id
    parentId   int64 

}

func NewTaobaoPictureCategoryAddRequest() *TaobaoPictureCategoryAddRequest{
    return &TaobaoPictureCategoryAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPictureCategoryAddRequest) GetApiMethodName() string {
    return "taobao.picture.category.add"
}

func (r TaobaoPictureCategoryAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPictureCategoryAddRequest) SetPictureCategoryName(pictureCategoryName string) error {
    r.pictureCategoryName = pictureCategoryName
    r.Set("picture_category_name", pictureCategoryName)
    return nil
}

func (r TaobaoPictureCategoryAddRequest) GetPictureCategoryName() string {
    return r.pictureCategoryName
}

func (r *TaobaoPictureCategoryAddRequest) SetParentId(parentId int64) error {
    r.parentId = parentId
    r.Set("parent_id", parentId)
    return nil
}

func (r TaobaoPictureCategoryAddRequest) GetParentId() int64 {
    return r.parentId
}

