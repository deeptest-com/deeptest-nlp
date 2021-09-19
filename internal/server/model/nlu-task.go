package model

type NluTask struct {
	BaseModel

	Name    string      `json:"name"`
	Intents []NluIntent `json:"intents" gorm:"-"`
	Ordr    int         `json:"ordr"`

	ProjectId   uint   `json:"projectId"`
	ProjectName string `json:"projectName"`
}

func (NluTask) TableName() string {
	return "nlu_task"
}
