package domain

type NluDomain struct {
	Version       string        `yaml:"version"`
	Intents       []string      `yaml:"intents,flow"`
	Entities      []string      `yaml:"entities,flow"`
	Slots         Slots         `yaml:"slots"`
	Responses     Responses     `yaml:"responses"`
	SessionConfig SessionConfig `yaml:"session_config"`
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

type Slots struct {
	Print    Loglevel `yaml:"print"`
	Loglevel Loglevel `yaml:"loglevel"`
	Message  Loglevel `yaml:"message"`
	Xpath    Loglevel `yaml:"xpath"`
}

type Loglevel struct {
	Type                  string `yaml:"type"`
	InfluenceConversation bool   `yaml:"influence_conversation"`
}
