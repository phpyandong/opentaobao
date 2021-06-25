package scbp

// TargetReportOperationDto 
type TargetReportOperationDto struct {

    // crowd/region
    Type   string `json:"type,omitempty"`

    // 开始时间(yyyy-MM-dd)
    DateBegin   string `json:"date_begin,omitempty"`

    // 结束时间(yyyy-MM-dd)
    DateEnd   string `json:"date_end,omitempty"`

    // 营销目标
    CampaignType   int64 `json:"campaign_type,omitempty"`

    // 计划ID
    CampaignId   int64 `json:"campaign_id,omitempty"`

}
