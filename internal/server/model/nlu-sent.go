package model

type NluSent struct {
	BaseModel

	Html string `json:"html"`
	Text string `json:"text"`
	Ordr int    `json:"ordr"`

	IntentId uint      `json:"intentId"`
	Slots    []NluSlot `json:"slots" gorm:"foreignKey:SentRefer"`
}

func (NluSent) TableName() string {
	return "nlu_sent"
}
