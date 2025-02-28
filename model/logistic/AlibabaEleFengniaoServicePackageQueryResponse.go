package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
预采购服务包查询接口 APIResponse
alibaba.ele.fengniao.service.package.query

查询门店所在经纬度可用服务包的接口
*/
type AlibabaEleFengniaoServicePackageQueryAPIResponse struct {
    model.CommonResponse
    AlibabaEleFengniaoServicePackageQueryResponse
}

type AlibabaEleFengniaoServicePackageQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_ele_fengniao_service_package_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // servicePackages
    
    ServicePackages   []AlibabaEleFengniaoServicePackageQueryResult `json:"service_packages,omitempty" xml:"service_packages>alibaba_ele_fengniao_service_package_query_result,omitempty"`
    
    
}
