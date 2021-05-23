package domain

type NluDomain struct {
	Version       string                       `yaml:"version"`
	SessionConfig SessionConfig                `yaml:"session_config"`
	Responses     Responses                    `yaml:"responses"`
	Intents       []string                     `yaml:"intents,flow"`
	Entities      []string                     `yaml:"entities,flow"`
	Slots         map[string]map[string]string `yaml:"slots"`
}

type Responses struct {
	UtterGreate []UtterGreate `yaml:"utter_greate,flow"`
}

type UtterGreate struct {
	Text string `yaml:"text"`
}

type SessionConfig struct {
	SessionExpirationTime      int64 `yaml:"session_expiration_time"`
	CarryOverSlotsToNewSession bool  `yaml:"carry_over_slots_to_new_session"`
}

type Loglevel struct {
	Type                  string `yaml:"type"`
	InfluenceConversation bool   `yaml:"influence_conversation"`
}
