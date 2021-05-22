package model

type Project struct {
	BaseModel

	Name      string `json:"name"`
	Desc      string `json:"desc"`
	Path      string `json:"path"`
	IsDefault bool   `json:"isDefault"`
}

func (Project) TableName() string {
	return "biz_project"
}
