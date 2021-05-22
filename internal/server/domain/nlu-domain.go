package domain

type NluDomain struct {
	Version       string        `json:"version"`
	Intents       []string      `json:"intents"`
	Entities      []string      `json:"entities"`
	Slots         Slots         `json:"slots"`
	Responses     Responses     `json:"responses"`
	SessionConfig SessionConfig `json:"session_config"`
}

type Responses struct {
	UtterGreate []UtterGreate `json:"utter_greate"`
}

type UtterGreate struct {
	Text string `json:"text"`
}

type SessionConfig struct {
	SessionExpirationTime      int64 `json:"session_expiration_time"`
	CarryOverSlotsToNewSession bool  `json:"carry_over_slots_to_new_session"`
}

type Slots struct {
	Print    Loglevel `json:"print"`
	Loglevel Loglevel `json:"loglevel"`
	Message  Loglevel `json:"message"`
	Xpath    Loglevel `json:"xpath"`
}

type Loglevel struct {
	Type                  string `json:"type"`
	InfluenceConversation bool   `json:"influence_conversation"`
}
