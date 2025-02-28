package traderate

// ItemStatisticVo 
type ItemStatisticVo struct {

    // 评论数量
    
    RateCnt   int64 `json:"rate_cnt,omitempty" xml:"rate_cnt,omitempty"`
    

    // tab筛选信息
    
    TabInfos   []TabInfo `json:"tab_infos,omitempty" xml:"tab_infos,omitempty"`
    

    // 总评分
    
    TotalScore   string `json:"total_score,omitempty" xml:"total_score,omitempty"`
    

    // 评分描述信息
    
    RankDesc   string `json:"rank_desc,omitempty" xml:"rank_desc,omitempty"`
    

    // 子评分项信息
    
    ScoreDetail   string `json:"score_detail,omitempty" xml:"score_detail,omitempty"`
    

}
