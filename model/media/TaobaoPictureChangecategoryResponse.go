package media

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改图片的分类 APIResponse
taobao.picture.changecategory

把批量的图片移动到某个分类下
*/
type TaobaoPictureChangecategoryAPIResponse struct {
    model.CommonResponse
    TaobaoPictureChangecategoryResponse
}

type TaobaoPictureChangecategoryResponse struct {
    XMLName xml.Name `xml:"picture_changecategory_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 移动图片是否成功：部分移动成功为true，全部移动失败为false。
    
    Done   bool `json:"done,omitempty" xml:"done,omitempty"`

    
}
