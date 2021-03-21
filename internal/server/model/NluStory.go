package model

type NluStory struct {
	Version string  `yaml:"version"`
	Stories []Story `yaml:"stories"`
}

type Story struct {
	BaseModel

	Story string `yaml:"story"`
	Steps []Step `yaml:"steps"`

	ProjectId uint `json:"projectId"`
}

type Step struct {
	BaseModel

	Intent string `json:"intent"`
	Action string `json:"action"`

	StoryId uint `json:"storyId"`
}
