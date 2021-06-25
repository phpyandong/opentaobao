package scbp

// TopP4pQuickForbiddenWordDto 
type TopP4pQuickForbiddenWordDto struct {

    // 屏蔽词
    ForbiddenWord   []String `json:"forbidden_word,omitempty"`

    // 1代表新增屏蔽词，2代表删除屏蔽词
    Action   int64 `json:"action,omitempty"`

    // 定向推广计划ID
    CampaignId   int64 `json:"campaign_id,omitempty"`

}
