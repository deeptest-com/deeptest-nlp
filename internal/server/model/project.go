package model

type Project struct {
	BaseModel

	Name string `yaml:"name"`
	Desc string `yaml:"desc"`
}

func (Project) TableName() string {
	return "biz_project"
}
