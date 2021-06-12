package model

type NluRule struct {
	BaseModel

	Text     string `json:"text"`
	IntentId uint   `json:"intentId"`
}

func (NluRule) TableName() string {
	return "nlu_rule"
}
