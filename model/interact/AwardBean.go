package interact

// AwardBean 
type AwardBean struct {

    // 奖品Id
    
    AwardId   int64 `json:"award_id,omitempty" xml:"award_id,omitempty"`
    

    // 奖品名称
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 奖品类型
    
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
    

}
