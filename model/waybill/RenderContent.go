package waybill

// RenderContent 
type RenderContent struct {

    // 打印数据
    
    PrintData   string `json:"print_data,omitempty" xml:"print_data,omitempty"`
    

    // 模板url
    
    TemplateUrl   string `json:"template_url,omitempty" xml:"template_url,omitempty"`
    

    // 是否获取加密数据
    
    Encrypted   bool `json:"encrypted,omitempty" xml:"encrypted,omitempty"`
    

    // 加密数据使用秘钥版本
    
    Ver   string `json:"ver,omitempty" xml:"ver,omitempty"`
    

    // 数据签名
    
    Signature   string `json:"signature,omitempty" xml:"signature,omitempty"`
    

    // 附加数据(用于修改数据)
    
    AddData   string `json:"add_data,omitempty" xml:"add_data,omitempty"`
    

}
