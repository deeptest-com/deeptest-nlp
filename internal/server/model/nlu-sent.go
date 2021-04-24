package model

type NluSent struct {
	BaseModel

	Content string    `json:"content"`
	Slots   []NluSlot `json:"slots" gorm:"-"`

	IntentId uint `json:"intentId"`
}

func (NluSent) TableName() string {
	return "nlu_sent"
}
