package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取门店类目信息 APIResponse
taobao.place.storecategory.get

获取门店类目信息
*/
type TaobaoPlaceStorecategoryGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoPlaceStorecategoryGetResponse `json:"taobao_place_storecategory_get_response,omitempty"`
}

type TaobaoPlaceStorecategoryGetResponse struct {

    // 门店类目格式
    CategoryList   string `json:"category_list,omitempty"`

}
