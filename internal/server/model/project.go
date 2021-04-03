package model

type Project struct {
	Name string    `yaml:"name"`
	Desc []NluSent `yaml:"desc"`
}

func (Project) TableName() string {
	return "biz_project"
}
