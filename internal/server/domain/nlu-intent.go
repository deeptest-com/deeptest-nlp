package domain

type NluIntent struct {
	Version   string      `yaml:"version"`
	IntentDef []IntentDef `yaml:"nlu,flow"`
}

type IntentDef struct {
	Intent   string `yaml:"intent"`
	Examples string `yaml:"examples"`
}
