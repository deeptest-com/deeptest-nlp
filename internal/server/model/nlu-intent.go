package model

type NluIntent struct {
	Version string    `yaml:"version,default:0"`
	Sents   []NluSent `yaml:"sents"`
}

type NluSent struct {
	BaseModel

	Content string    `json:"content"`
	Slots   []NluSlot `yaml:"slots"`

	IntentId uint `json:"intentId"`
}

type NluSlot struct {
	BaseModel

	Name   string `json:"name"`
	Entity string `json:"entity"`
	Value  string `json:"value"`

	SentId uint `json:"sentId"`
}

func (NluIntent) TableName() string {
	return "nlu_intent"
}
func (NluSent) TableName() string {
	return "nlu_sent"
}
func (NluSlot) TableName() string {
	return "nlu_slot"
}
