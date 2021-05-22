package domain

type NluLookup struct {
	Version   string         `yaml:"version"`
	LookupDef []NluLookupDef `yaml:"nlu,flow"`
}

type NluLookupDef struct {
	Lookup   string `yaml:"lookup"`
	Examples string `yaml:"examples"`
}
