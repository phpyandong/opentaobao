package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
派样统计数据查询 APIResponse
alibaba.wdk.bm.paiyang.stat.data.query

派样统计数据查询
*/
type AlibabaWdkBmPaiyangStatDataQueryAPIResponse struct {
    model.CommonResponse
    AlibabaWdkBmPaiyangStatDataQueryResponse
}

type AlibabaWdkBmPaiyangStatDataQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_bm_paiyang_stat_data_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 出参对象
    
    Result   *BmPageResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
