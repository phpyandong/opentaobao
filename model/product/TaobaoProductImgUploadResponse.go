package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
上传单张产品非主图，如果需要传多张，可调多次 APIResponse
taobao.product.img.upload

1.传入产品ID <br/>2.传入图片内容 <br/>注意：图片最大为500K,只支持JPG,GIF格式,如果需要传多张，可调多次
*/
type TaobaoProductImgUploadAPIResponse struct {
    model.CommonResponse
    TaobaoProductImgUploadResponse
}

type TaobaoProductImgUploadResponse struct {
    XMLName xml.Name `xml:"product_img_upload_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回产品图片结构中的：url,id,created,modified
    
    ProductImg   *ProductImg `json:"product_img,omitempty" xml:"product_img,omitempty"`

    
}
