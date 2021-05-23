package domain

type NluRegex struct {
	Version string         `yaml:"version"`
	Items   []NluRegexItem `yaml:"nlu,flow"`
}

type NluRegexItem struct {
	Regex    string `yaml:"regex"`
	Examples string `yaml:"examples"`
}
