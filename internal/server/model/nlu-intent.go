package model

type NluIntent struct {
	BaseModel

	Name   string    `json:"name"`
	Sents  []NluSent `json:"sents" gorm:"-"`
	TaskId uint      `json:"taskId"`

	Ordr int `json:"ordr"`
}

func (NluIntent) TableName() string {
	return "nlu_intent"
}
