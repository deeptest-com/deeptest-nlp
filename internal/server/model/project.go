package model

type Project struct {
	BaseModel

	Name      string `json:"name"`
	Desc      string `json:"desc"`
	IsDefault string `json:"isDefault"`
}

func (Project) TableName() string {
	return "biz_project"
}
