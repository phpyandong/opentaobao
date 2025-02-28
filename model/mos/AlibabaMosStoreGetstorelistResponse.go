package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据屏编号获取专柜集 APIResponse
alibaba.mos.store.getstorelist

根据屏编号获取专柜集
*/
type AlibabaMosStoreGetstorelistAPIResponse struct {
    model.CommonResponse
    AlibabaMosStoreGetstorelistResponse
}

type AlibabaMosStoreGetstorelistResponse struct {
    XMLName xml.Name `xml:"alibaba_mos_store_getstorelist_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlibabaMosStoreGetstorelistResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
