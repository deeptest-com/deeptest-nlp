package model

type NluIntent struct {
	BaseModel

	Name   string    `json:"name"`
	Sents  []NluSent `json:"sents" gorm:"-"`
	TaskId uint      `json:"taskId"`
}

func (NluIntent) TableName() string {
	return "nlu_intent"
}
