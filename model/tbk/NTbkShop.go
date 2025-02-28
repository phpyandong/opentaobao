package tbk

// NTbkShop 
type NTbkShop struct {

    // 卖家ID
    
    UserId   int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
    

    // 店铺名称
    
    ShopTitle   string `json:"shop_title,omitempty" xml:"shop_title,omitempty"`
    

    // 店铺类型，B：天猫，C：淘宝
    
    ShopType   string `json:"shop_type,omitempty" xml:"shop_type,omitempty"`
    

    // 卖家昵称
    
    SellerNick   string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
    

    // 店标图片
    
    PictUrl   string `json:"pict_url,omitempty" xml:"pict_url,omitempty"`
    

    // 店铺地址
    
    ShopUrl   string `json:"shop_url,omitempty" xml:"shop_url,omitempty"`
    

}
