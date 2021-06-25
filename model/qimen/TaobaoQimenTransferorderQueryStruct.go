package qimen

// TaobaoQimenTransferorderQueryStruct 
type TaobaoQimenTransferorderQueryStruct struct {

    // 调拨单号
    TransferOrderCode   string `json:"transferOrderCode,omitempty"`

    // 111
    OwnerCode   string `json:"ownerCode,omitempty"`

    // ERP单号
    ErpOrderCode   string `json:"erpOrderCode,omitempty"`

    // 响应结果:success|failure,success,string(10),必填,
    Flag   string `json:"flag,omitempty"`

    // 响应码,0,string(50),,
    Code   string `json:"code,omitempty"`

    // 响应信息,invalid appkey,string(100),,
    Message   string `json:"message,omitempty"`

    // 调拨单细节
    TransferOrderDetail   *TransferOrderDetail `json:"transferOrderDetail,omitempty"`

}
