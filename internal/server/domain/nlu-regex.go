package domain

type NluRegex struct {
	Version  string     `json:"version"`
	RegexDef []RegexDef `json:"nlu"`
}

type RegexDef struct {
	Regex    string `json:"regex"`
	Examples string `json:"examples"`
}
