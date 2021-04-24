package model

type NluIntent struct {
	BaseModel

	Name  string    `json:"name"`
	Sents []NluSent `json:"sents" gorm:"-"`
}

func (NluIntent) TableName() string {
	return "nlu_intent"
}
