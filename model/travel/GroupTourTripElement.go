package travel

// GroupTourTripElement 
type GroupTourTripElement struct {

    // 必填，第几天的行程信息
    
    Day   int64 `json:"day,omitempty" xml:"day,omitempty"`
    

    // 必填，包含元素类型，1:住宿信息，2:景点，3:餐饮信息，4:购物信息，5:自由活动
    
    Type   int64 `json:"type,omitempty" xml:"type,omitempty"`
    

    // 具体的行程信息，根据type字段，将对象序列化成json串，以字符串的形式赋值给json_str，传到后端,每一个json_str都只能是对应单个对象，不能对应数组  当type=1时： json_str = {     "type":1, //包含元素类型，1:住宿信息，2:景点，3:餐饮信息，4:购物信息     "hotelType”:1,//住宿方式，1:酒店/客栈 2:住在交通工具上 3:住宿自理 4:露营     "hotelStarType":1, //1:酒店星级 2:酒店档次     "hotelStar":"三星级", //如果hotelStarType =1:酒店星级，hotelStar取值范围（一星级，二星级，三星级，四星级，五星级）；如果hotelStarType =1:酒店星级，hotelStar取值范围（舒适,高档,豪华,经济连锁，二星及以下）     "hotelCityName":"北京市", //酒店所在城市名称     "hotelName":"如家快捷北京北太平庄店”,//酒店名称     "roomType":"大床房”, //房型     "tripContentDetails":{ //该字段选填，imageList为数组类型，住宿图片，desc为住宿说明                           "imageList":[                                        "https://img.daily.taobaocdn.net/imgextra/i2/2024098454/TB2.BJ9XEw7LKJjyzdKXXaShXXa_!!2024098454.jpg",                                        "https://img.daily.taobaocdn.net/imgextra/i1/2024098454/TB2Ui4yXEw7LKJjyzdKXXaShXXa_!!2024098454.jpg"                                       ],                           "desc”:”住宿说明”                          } } 当type=2时： json_str = {     "type”:2, //包含元素类型，1:住宿信息，2:景点，3:餐饮信息，4:购物信息     "activityHour":10, //活动时间-小时     "activityMinute":30,//活动时间-分钟,这里是10个小时30分钟     "scenicName":"八达岭长城”,//景点名称     "scenicCity":"北京市” //景点所在城市,”classicScenic”:true, "tripContentDetails":{ //该字段选填，imageList为数组类型，景点图片，desc为景点详细说明                           "imageList":[                                        "https://img.daily.taobaocdn.net/imgextra/i2/2024098454/TB2.BJ9XEw7LKJjyzdKXXaShXXa_!!2024098454.jpg",                                        "https://img.daily.taobaocdn.net/imgextra/i1/2024098454/TB2Ui4yXEw7LKJjyzdKXXaShXXa_!!2024098454.jpg"                                       ],                           "desc”:”景点详情”                          } } 当type=3时： json_str ={     "type":3, //包含元素类型，1:住宿信息，2:景点，3:餐饮信息，4:购物信息     "foodInclude":false,//true：包含餐饮，false:不包含餐饮信息     "specialIllustrate":" 餐饮说明”,//餐饮说明     "foodType":[ //1:早餐，2:中餐，3:晚餐         1,         2,         3     ] } 当type=4时： json_str ={     "type":4, //包含元素类型，1:住宿信息，2:景点，3:餐饮信息，4:购物信息     "activityHour":1,//活动时间-小时     "activityMinute":5,//活动时间-分钟,这里是1个小时5分钟     "shoppingPlace":"家乐福”,//购物店名称     "shoppingProduct":"啥都有”//营业产品 }。当type=5时： json_str ={ "type":5, //包含元素类型，1:住宿信息，2:景点，3:餐饮信息，4:购物信息，5:自由活动 ,"activityHour":1,//活动时间-小时 "activityMinute":5,//活动时间-分钟,这里是1个小时5分钟 "scenicCity":"杭州”,//活动城市, "activityContent":"啥都有”//活动推荐 }
    
    JsonStr   string `json:"json_str,omitempty" xml:"json_str,omitempty"`
    

}
