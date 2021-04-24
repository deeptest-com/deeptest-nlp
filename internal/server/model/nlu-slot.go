package model

type NluSlot struct {
	BaseModel

	Name   string `json:"name"`
	Entity string `json:"entity"`
	Value  string `json:"value"`

	SentId uint `json:"sentId"`
}

func (NluSlot) TableName() string {
	return "nlu_slot"
}
