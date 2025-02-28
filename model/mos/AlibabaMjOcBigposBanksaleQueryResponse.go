package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
大pos银行卡查账接口 APIResponse
alibaba.mj.oc.bigpos.banksale.query

大pos银行卡查账接口，给收银员查询银行卡销售记录，便于调账
*/
type AlibabaMjOcBigposBanksaleQueryAPIResponse struct {
    model.CommonResponse
    AlibabaMjOcBigposBanksaleQueryResponse
}

type AlibabaMjOcBigposBanksaleQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_mj_oc_bigpos_banksale_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 明细数量
    
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`

    
    // 数据主体
    
    Datas   []AlibabaMjOcBigposBanksaleQueryData `json:"datas,omitempty" xml:"datas>alibaba_mj_oc_bigpos_banksale_query_data,omitempty"`
    
    
}
