package model

type NluRule struct {
	BaseModel

	Text      string `json:"text"`
	IntentId  uint   `json:"intentId"`
	ProjectId uint   `json:"projectId"`
}

func (NluRule) TableName() string {
	return "nlu_rule"
}
