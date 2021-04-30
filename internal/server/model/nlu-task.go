package model

type NluTask struct {
	BaseModel

	Name    string      `json:"name"`
	Intents []NluIntent `json:"intents" gorm:"-"`

	ProjectId   uint `json:"projectId"`
	ProjectName uint `json:"projectName" gorm:"-"s`
}

func (NluTask) TableName() string {
	return "nlu_task"
}
