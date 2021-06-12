package model

type NluIntent struct {
	BaseModel

	Name   string    `json:"name"`
	Sents  []NluSent `json:"sents" gorm:"-"`
	Rules  []NluRule `json:"rules" gorm:"-"`
	TaskId uint      `json:"taskId"`

	Ordr int `json:"ordr"`
}

func (NluIntent) TableName() string {
	return "nlu_intent"
}
