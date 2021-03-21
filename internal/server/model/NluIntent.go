package model

type NluIntent struct {
	Version string `yaml:"version,default:0"`
	Nlu     []Nlu  `yaml:"nlu"`
}
type Nlu struct {
	Intent   string `yaml:"intent,omitempty"`
	Examples string `yaml:"examples"`
	Lookup   string `yaml:"lookup,omitempty"`
}

type Intent struct {
	BaseModel
	ProjectId uint `json:"projectId"`
}
type Lookup struct {
	BaseModel
	ProjectId uint `json:"projectId"`
}
type Sentence struct {
	BaseModel

	Content string `json:"content"`

	IntentId uint `json:"intentId"`
	LookupId uint `json:"lookupId"`
}
type Slot struct {
	BaseModel

	Name   string `json:"name"`
	Entity string `json:"entity"`
	Value  string `json:"value"`

	Type                  string `json:"value"`
	InfluenceConversation bool   `json:"influence_conversation"`

	IntentId uint `json:"projectId"`
}
