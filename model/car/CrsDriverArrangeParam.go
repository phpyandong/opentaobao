package car

// CrsDriverArrangeParam 
type CrsDriverArrangeParam struct {

    // 城市
    
    CityName   string `json:"city_name,omitempty" xml:"city_name,omitempty"`
    

    // 是否改派司机
    
    IsChangeDriver   bool `json:"is_change_driver,omitempty" xml:"is_change_driver,omitempty"`
    

    // 车型名称
    
    CarTypeName   string `json:"car_type_name,omitempty" xml:"car_type_name,omitempty"`
    

    // 司机名称
    
    DriverName   string `json:"driver_name,omitempty" xml:"driver_name,omitempty"`
    

    // 飞猪订单id
    
    OrderId   int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
    

    // 司机手机号
    
    DriverPhoneNum   string `json:"driver_phone_num,omitempty" xml:"driver_phone_num,omitempty"`
    

    // 司机身份证号
    
    DriverIdCard   string `json:"driver_id_card,omitempty" xml:"driver_id_card,omitempty"`
    

    // 电话区号
    
    PhoneAreaCode   string `json:"phone_area_code,omitempty" xml:"phone_area_code,omitempty"`
    

    // 车型号
    
    CarType   string `json:"car_type,omitempty" xml:"car_type,omitempty"`
    

    // 车牌号
    
    CarNumber   string `json:"car_number,omitempty" xml:"car_number,omitempty"`
    

    // 车辆品牌
    
    CarBrand   string `json:"car_brand,omitempty" xml:"car_brand,omitempty"`
    

}
