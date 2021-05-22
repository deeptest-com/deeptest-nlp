package domain

type NluRegex struct {
	Version  string        `yaml:"version"`
	RegexDef []NluRegexDef `yaml:"nlu,flow"`
}

type NluRegexDef struct {
	Regex    string `yaml:"regex"`
	Examples string `yaml:"examples"`
}
