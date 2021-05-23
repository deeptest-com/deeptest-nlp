package domain

type NluIntent struct {
	Version string          `yaml:"version"`
	Items   []NluIntentItem `yaml:"nlu,flow"`
}

type NluIntentItem struct {
	Intent   string `yaml:"intent"`
	Examples string `yaml:"examples"`
}
