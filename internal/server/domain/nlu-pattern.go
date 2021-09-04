package serverDomain

type NluTaskForPattern struct {
	Version string             `yaml:"version" default:"2.0"`
	Name    string             `yaml:"name"`
	Intents []NluIntentPattern `yaml:"intents,flow"`
}

type NluIntentPattern struct {
	Id   uint   `yaml:"id"`
	Name string `yaml:"name"`

	Sents []NluSentPattern `yaml:"examples"`
}

type NluSentPattern struct {
	Id      uint   `yaml:"id"`
	Example string `yaml:"example"`
}
