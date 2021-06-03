package domain

type NluTask struct {
	Version string      `yaml:"version"`
	Intents []NluIntent `yaml:"nlu,flow"`
}
