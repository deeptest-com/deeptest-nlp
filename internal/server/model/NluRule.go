package model

type NluRule struct {
	Version string  `yaml:"version"`
	Rules   []Rules `yaml:"rules"`
}

type Rules struct {
	BaseModel

	Rule  string `json:"rule"`
	Steps []Step `json:"steps"`

	ProjectId uint `json:"projectId"`
}
