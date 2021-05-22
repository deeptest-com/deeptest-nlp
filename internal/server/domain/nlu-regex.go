package domain

type NluRegex struct {
	Version  string        `json:"version"`
	RegexDef []NluRegexDef `json:"nlu"`
}

type NluRegexDef struct {
	Regex    string `json:"regex"`
	Examples string `json:"examples"`
}
