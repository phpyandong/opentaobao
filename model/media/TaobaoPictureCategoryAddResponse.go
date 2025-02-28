package media

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新增图片分类信息 APIResponse
taobao.picture.category.add

同一卖家最多添加500个图片分类，图片分类名称长度最大为20个字符
*/
type TaobaoPictureCategoryAddAPIResponse struct {
    model.CommonResponse
    TaobaoPictureCategoryAddResponse
}

type TaobaoPictureCategoryAddResponse struct {
    XMLName xml.Name `xml:"picture_category_add_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 图片分类信息
    
    PictureCategory   *PictureCategory `json:"picture_category,omitempty" xml:"picture_category,omitempty"`

    
}
