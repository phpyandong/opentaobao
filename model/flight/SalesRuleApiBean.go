package flight

// SalesRuleApiBean 
type SalesRuleApiBean struct {

    // 适用舱位(可以多个，支持子舱，用“/”隔开)；
    
    Cabin   string `json:"cabin,omitempty" xml:"cabin,omitempty"`
    

    // 例外舱位等级,单选 0:头等,1:商务，2:经济，3:全部 适用和例外二选一填写
    
    ExceptCabinClass   int64 `json:"except_cabin_class,omitempty" xml:"except_cabin_class,omitempty"`
    

    // 最晚购买时间，航班起飞前n小时 (计划起飞时间-当前时间)>=n小时
    
    LatestBuyTime   int64 `json:"latest_buy_time,omitempty" xml:"latest_buy_time,omitempty"`
    

    // 例外航线(可以多个,机场码) PEKPVG/DMKYVR ods和exceptOds二选一填写。 适用和例外二选一填写。 最多允许510个字符。
    
    ExceptOds   string `json:"except_ods,omitempty" xml:"except_ods,omitempty"`
    

    // 旅行开始日期(yyyyMMdd)
    
    TravelDateStart   string `json:"travel_date_start,omitempty" xml:"travel_date_start,omitempty"`
    

    // 旅行结束日期(yyyyMMdd)，最大值当前+365且不能小于开始
    
    TravelDateEnd   string `json:"travel_date_end,omitempty" xml:"travel_date_end,omitempty"`
    

    // 最早购买时间，航班起飞前n小时 (计划起飞时间-当前时间)<=n小时 如无限制，请填写9000
    
    EarlistBuyTime   int64 `json:"earlist_buy_time,omitempty" xml:"earlist_buy_time,omitempty"`
    

    // 适用舱位等级,单选 0:头等,1:商务，2:经济，3:全部 适用和例外二选一填写
    
    CabinClass   int64 `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
    

    // 例外机场，如果为多个机场，按照PEK/NAY 格式返回 airports和exceptAirports二选一填写。 最多允许511个字符。
    
    ExceptAirports   string `json:"except_airports,omitempty" xml:"except_airports,omitempty"`
    

    // 销售结束日期(yyyyMMdd)，最大值当前+365且不能小于开始
    
    SaleDateEnd   string `json:"sale_date_end,omitempty" xml:"sale_date_end,omitempty"`
    

    // 发票类型 1:电子，2:纸质，3:不提供
    
    ReceiptType   int64 `json:"receipt_type,omitempty" xml:"receipt_type,omitempty"`
    

    // 库存数，当stockType=1或stockType=2时必填
    
    StockNum   int64 `json:"stock_num,omitempty" xml:"stock_num,omitempty"`
    

    // 例外航班(可以多个，用“/”隔开)； 适用和例外二选一填写。 最多允许512个字符。
    
    ExceptFlightNos   string `json:"except_flight_nos,omitempty" xml:"except_flight_nos,omitempty"`
    

    // 适用航线(可以多个,机场码) PEKPVG/DMKYVR ods和exceptOds二选一填写。 适用和例外二选一填写。 最多允许510个字符。
    
    Ods   string `json:"ods,omitempty" xml:"ods,omitempty"`
    

    // 适用航司二字码
    
    AirlineCodes   string `json:"airline_codes,omitempty" xml:"airline_codes,omitempty"`
    

    // 例外舱位(可以多个，支持子舱，用“/”隔开)；
    
    ExceptCabin   string `json:"except_cabin,omitempty" xml:"except_cabin,omitempty"`
    

    // 库存类型 1:共享总库存数，2:每日库存数（旅行日期下每天），3:不限制(不使用产品库库存)
    
    StockType   int64 `json:"stock_type,omitempty" xml:"stock_type,omitempty"`
    

    // 适用机场，如果为多个机场，按照PEK/NAY 格式返回 airports和exceptAirports二选一填写。 最多允许511个字符。
    
    Airports   string `json:"airports,omitempty" xml:"airports,omitempty"`
    

    // 发票领取规则 1:行程单，2:机场柜台 默认为1.当receiptType=1或receiptType=2时，receiptWay必须指定
    
    ReceiptWay   int64 `json:"receipt_way,omitempty" xml:"receipt_way,omitempty"`
    

    // 销售开始日期(yyyyMMdd)
    
    SaleDateStart   string `json:"sale_date_start,omitempty" xml:"sale_date_start,omitempty"`
    

    // 例外航司二字码，如果为多个航司，按照 AK/FD格式返回。 最多允许254个字符。
    
    ExceptAirlineCodes   string `json:"except_airline_codes,omitempty" xml:"except_airline_codes,omitempty"`
    

    // 共享航班是否可用 1:不能用于共享航班，2:可用
    
    CodeShareForbidden   bool `json:"code_share_forbidden,omitempty" xml:"code_share_forbidden,omitempty"`
    

    // 适用航班(可以多个，用“/”隔开)； 适用和例外二选一填写。 最多允许512个字符。
    
    FlightNos   string `json:"flight_nos,omitempty" xml:"flight_nos,omitempty"`
    

    // 适用出发航站楼（可以多个，用“/”隔开)。 最多允许10个字符。
    
    OutboundTerminal   string `json:"outbound_terminal,omitempty" xml:"outbound_terminal,omitempty"`
    

    // 库存规则，当stockType=1或stockType=2时必填。
    
    StockRule   string `json:"stock_rule,omitempty" xml:"stock_rule,omitempty"`
    

}
