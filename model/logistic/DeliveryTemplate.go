package logistic

// DeliveryTemplate 
type DeliveryTemplate struct {

    // 模块ID
    
    TemplateId   int64 `json:"template_id,omitempty" xml:"template_id,omitempty"`
    

    // 模板创建时间
    
    Created   string `json:"created,omitempty" xml:"created,omitempty"`
    

    // 模板名称，长度不能超过25个字节
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 可选值：0,1,2,3<br>,说明如下<br>
1)买家承担运费的模版<br>
0：买家承担服务费<br>
1: 卖家承担服务费<br>
2)卖家承担运费的模版<br>
2:卖家承担运费的模版（集市），卖家承担服务费<br>
3:卖家承担运费的模版（天猫），卖家承担服务费<br>
    
    Assumer   int64 `json:"assumer,omitempty" xml:"assumer,omitempty"`
    

    // 可选值：0<br/>说明：<br/>0:表示按宝贝件数计算运费<br/><br/><br/><br/>1:表示按宝贝重量计算运费<br/><br/><br/><br/>3:表示按宝贝体积计算运费
    
    Valuation   int64 `json:"valuation,omitempty" xml:"valuation,omitempty"`
    

    // 运费模板中运费详细信息对象，包含默认运费和指定地区运费
    
    FeeList   []TopFee `json:"fee_list,omitempty" xml:"fee_list,omitempty"`
    

    // 物流服务模板支持增值服务列表，多个用分号隔开。cod:货到付款 mops：刷卡付款
    
    Supports   string `json:"supports,omitempty" xml:"supports,omitempty"`
    

    // 模板修改时间
    
    Modified   string `json:"modified,omitempty" xml:"modified,omitempty"`
    

    // 运费模板上设置的发货地址
    
    Address   string `json:"address,omitempty" xml:"address,omitempty"`
    

    // 该模板上设置的卖家发货地址区域ID，如：address为浙江省杭州市西湖去文三路XX号那么这个consign_area_id的值就是西湖区的ID
    
    ConsignAreaId   int64 `json:"consign_area_id,omitempty" xml:"consign_area_id,omitempty"`
    

}
