package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询订单基本信息 APIResponse
alibaba.ele.fengniao.order.query

查询订单基本信息
*/
type AlibabaEleFengniaoOrderQueryAPIResponse struct {
    model.CommonResponse
    AlibabaEleFengniaoOrderQueryResponse
}

type AlibabaEleFengniaoOrderQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_ele_fengniao_order_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // kvs
    
    Kvs   []Kvs `json:"kvs,omitempty" xml:"kvs>kvs,omitempty"`
    
    
    // 收件人经度
    
    ReceiverLongitude   string `json:"receiver_longitude,omitempty" xml:"receiver_longitude,omitempty"`

    
    // 寄件人纬度
    
    TransportLatitude   string `json:"transport_latitude,omitempty" xml:"transport_latitude,omitempty"`

    
    // 寄件人经度
    
    TransportLongitude   string `json:"transport_longitude,omitempty" xml:"transport_longitude,omitempty"`

    
    // 订单号
    
    OrderId   int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`

    
    // 收件人纬度
    
    ReceiverLatitude   string `json:"receiver_latitude,omitempty" xml:"receiver_latitude,omitempty"`

    
}
