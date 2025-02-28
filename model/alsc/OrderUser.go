package alsc

// OrderUser 
type OrderUser struct {

    // 邮箱
    
    Email   string `json:"email,omitempty" xml:"email,omitempty"`
    

    // 名称
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 昵称
    
    NickName   string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
    

    // 手机号
    
    Phone   string `json:"phone,omitempty" xml:"phone,omitempty"`
    

    // 用户类型：  WECHAT-微信  ALIPAY-支付宝OTHER-其他
    
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
    

    // 用户userId（支付宝ID/微信ID）
    
    UserId   string `json:"user_id,omitempty" xml:"user_id,omitempty"`
    

}
