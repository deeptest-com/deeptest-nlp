package domain

type NluIntent struct {
	Version   string      `json:"version"`
	IntentDef []IntentDef `json:"nlu"`
}

type IntentDef struct {
	Intent   string `json:"intent"`
	Examples string `json:"examples"`
}
