package model

type NluConfig struct {
	BaseModel

	Lang    string `yaml:"lang,omitempty"`
	Content string `yaml:"content"`

	ProjectId uint `json:"projectId"`
}
