package serverDomain

type NluLookup struct {
	Version string          `yaml:"version"`
	Items   []NluLookupItem `yaml:"nlu,flow"`
}

type NluLookupItem struct {
	Lookup   string `yaml:"lookup"`
	Examples string `yaml:"examples"`
}
