package domain

type NluTask struct {
	Version string      `yaml:"version" default:"2.0"`
	Intents []NluIntent `yaml:"nlu,flow"`
}
