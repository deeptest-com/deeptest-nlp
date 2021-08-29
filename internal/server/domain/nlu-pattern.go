package domain

type NluPatternTask struct {
	Version string       `yaml:"version" default:"2.0"`
	Name    string       `yaml:"name"`
	Intents []NluPattern `yaml:"intents,flow"`
}

type NluPattern struct {
	Id   uint   `yaml:"id"`
	Name string `yaml:"name"`

	Examples string `yaml:"examples"`
}
