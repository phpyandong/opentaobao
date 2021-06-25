package wdk

// EducationExpInfo 
type EducationExpInfo struct {

    // 学历
    Education   string `json:"education,omitempty"`

    // 开始日期
    GmtEnd   string `json:"gmt_end,omitempty"`

    // 结束日期
    GmtStart   string `json:"gmt_start,omitempty"`

    // 专业
    Major   string `json:"major,omitempty"`

    // 学校
    School   string `json:"school,omitempty"`

    // 学制年数
    SchoolingYears   string `json:"schooling_years,omitempty"`

}
