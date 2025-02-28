package waybill

// UserInfoDto 
type UserInfoDto struct {

    // 发货地址需要通过<a href="http://open.taobao.com/doc2/detail.htm?spm=a219a.7629140.0.0.3OFCPk&treeId=17&articleId=104860&docType=1">search接口</a>
    
    Address   *AddressDto `json:"address,omitempty" xml:"address,omitempty"`
    

    // 手机号码（手机号和固定电话不能同时为空），长度小于20
    
    Mobile   string `json:"mobile,omitempty" xml:"mobile,omitempty"`
    

    // 姓名，长度小于40
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 固定电话（手机号和固定电话不能同时为空），长度小于20
    
    Phone   string `json:"phone,omitempty" xml:"phone,omitempty"`
    

    // 开放地址ID
    
    Oaid   string `json:"oaid,omitempty" xml:"oaid,omitempty"`
    

    // 菜鸟地址ID，针对电商平台加密订单场景使用，淘系订单使用oaid，非淘使用caid。
    
    Caid   string `json:"caid,omitempty" xml:"caid,omitempty"`
    

}
