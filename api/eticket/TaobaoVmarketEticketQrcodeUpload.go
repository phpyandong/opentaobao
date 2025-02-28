package eticket

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/eticket"
)

/* 
码商二维码图片上传 
taobao.vmarket.eticket.qrcode.upload

电子凭证的码商可以通过这个接口，上传他们发送的二维码图片
*/
func TaobaoVmarketEticketQrcodeUpload(clt *core.SDKClient, req *eticket.TaobaoVmarketEticketQrcodeUploadRequest, session string) (*eticket.TaobaoVmarketEticketQrcodeUploadAPIResponse, error) {
    var resp eticket.TaobaoVmarketEticketQrcodeUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
