package model

type NluDomain struct {
	BaseModel

	Version  string                 `yaml:"version"`
	Intents  []string               `yaml:"intents"`
	Entities []string               `yaml:"entities"`
	Slots    []map[string]SlotValue `yaml:"slots"`
	Actions  []string               `yaml:"actions"`

	Responses     Responses     `yaml:"responses"`
	SessionConfig SessionConfig `yaml:"session_config"`

	ProjectId uint `json:"projectId"`
}

type SlotValue struct {
	Type                  string `yaml:"type"`
	InfluenceConversation bool   `yaml:"influence_conversation"`
}

type Responses struct {
	ItemsMap map[string][]ResponseItem `yaml:"utter_greet"`
}
type ResponseItem struct {
	Text string `yaml:"text"`
}

type SessionConfig struct {
	SessionExpirationTime      int  `yaml:"session_expiration_time"`
	CarryOverSlotsToNewSession bool `yaml:"carry_over_slots_to_new_session"`
}

type Domain struct {
	BaseModel
	ProjectId uint `json:"projectId"`
}
type Action struct {
	BaseModel
	DomainId uint `json:"domainId"`
}

type Response struct {
	BaseModel
	Name     string `json:"name"`
	DomainId uint   `json:"domainId"`
}
type ResponseText struct {
	BaseModel
	Text       string `json:"name"`
	ResponseId uint   `json:"responseId"`
}

type SessionConf struct {
	SessionExpirationTime      int  `json:"sessionExpirationTime"`
	CarryOverSlotsToNewSession bool `json:"carryOverSlotsToNewSession"`
}
