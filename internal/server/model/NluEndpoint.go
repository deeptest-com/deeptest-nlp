package model

type NluEndpoint struct {
	BaseModel

	Url       string `yaml:"url"`
	ProjectId uint   `json:"projectId"`
}
