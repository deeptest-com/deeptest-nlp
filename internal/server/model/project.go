package model

type Project struct {
	BaseModel

	Name string `json:"name"`
	Desc string `json:"desc"`
}

func (Project) TableName() string {
	return "biz_project"
}
